package storage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	UserHeader = "X_USER_ID"
)

type UploadFileRequest struct {
	GroupId string `schema:"groupId,required"`
}

var (
	secretId    string
	secretKey   string
	bucketPath  *url.URL
	servicePath *url.URL
)

func InitStorage(id, key, bucket string) {
	secretKey = key
	secretId = id
	var err error
	bucketPath, err = url.Parse(bucket)
	if err != nil {
		log.Fatal(err)
	}
	servicePath, err = url.Parse("https://cos.ap-nanjing.myqcloud.com")
	if err != nil {
		log.Fatal(err)
	}

}

func Store(r io.Reader, filePath string) error {
	client := makeClient()
	name := "tes\\objectPut.txt"

	_, err := client.Object.Put(context.Background(), name, r, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Load() error {
	client := makeClient()
	name := "test/objectPut.go"
	resp, err := client.Object.Get(context.Background(), name, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Printf("%s\n", string(bs))
	return nil
}

func Delete() error {
	client := makeClient()
	name := "test/objectPut.go"
	_, err := client.Object.Delete(context.Background(), name)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func makeClient() *cos.Client {
	baseUrl := &cos.BaseURL{BucketURL: bucketPath, ServiceURL: servicePath}
	client := cos.NewClient(baseUrl, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretId,
			SecretKey: secretKey},
	})
	return client
}

func getUserAndGroup(r *http.Request) (user, group int, e error) {
	if !r.URL.Query().Has("groupId") {
		return -1, -1, fmt.Errorf("not enough argument")
	}

	groupId, userid := r.URL.Query().Get("groupId"), r.Header.Get(UserHeader)
	var err error
	if user, err = strconv.Atoi(userid); err != nil {
		return -1, -1, err
	}
	if group, err = strconv.Atoi(groupId); err != nil {
		return -1, -1, err
	}

	return user, group, nil

}

func checkUserInGroup(user, group int, tx *sql.Tx) error {
	s := `select id from user where id=? and user_group=?`
	var id string
	if err := tx.QueryRow(s, user, group).Scan(&id); err != nil {
		return err
	}
	return nil
}

type FileResponse struct {
	Ok     bool   `json:"ok"`
	Reason string `json:"reason"`
}

func genFilePath(group int, filename string) string {
	g := strconv.Itoa(group)
	return g + "/" + filename
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		HandleError(err, w, http.StatusBadRequest)
		return
	}

	user, group, err := getUserAndGroup(r)
	if err != nil {
		HandleError(err, w, http.StatusBadRequest)
		return
	}

	response := &FileResponse{}
	uploadFile, header, err := r.FormFile("file")
	if err != nil {
		HandleError(err, w, http.StatusBadRequest)
		return
	}
	if strings.ContainsRune(header.Filename, '/') || strings.ContainsRune(header.Filename, '\\') {
		response.Reason = "Illegal file name"
		WriteJson(w, response)
		return
	}

	defer func(uploadFile multipart.File) {
		err := uploadFile.Close()
		if err != nil {
			log.Println(err)
		}
	}(uploadFile)
	tx, err := DB.Begin()

	if err != nil {
		_ = tx.Rollback()
		HandleError(err, w, http.StatusInternalServerError)
		return
	}
	if err := checkUserInGroup(user, group, tx); err != nil {
		_ = tx.Rollback()
		HandleError(err, w, http.StatusUnauthorized)
		return
	}

	s := `Insert into files (uploader,user_group,filename,uploadTime) values (?,?,?,?)`

	_, err = tx.Exec(s, user, group, header.Filename, time.Now().Unix())

	if err != nil {
		_ = tx.Rollback()
		response.Reason = "File existed"
		WriteJson(w, response)
		return
	}

	if err := Store(uploadFile, genFilePath(group, header.Filename)); err != nil {
		HandleError(err, w, http.StatusInternalServerError)
		_ = tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		response.Reason = "Failed to commit"
	} else {
		response.Ok = true
	}
	WriteJson(w, response)
}
