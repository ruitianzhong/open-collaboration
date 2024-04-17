package storage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gorilla/schema"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

const (
	UserHeader      = "X_USER_ID"
	COSMetaUploader = "x-cos-meta-uploader"
	COSMetaGroup    = "x-cos-meta-group"
	COSFilename     = "x-cos-meta-filename"
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

type FileMetaResponse struct {
	Files []FileMetaData `json:"files"`
	Ok    bool           `json:"ok"`
}

type FileMetaData struct {
	LastModified string `json:"lastModified"`
	Size         int64  `json:"size"`
	Uploader     string `json:"uploader"`
	Filename     string `json:"filename"`
}

func List(prefix string) ([]FileMetaData, error) {
	client := makeClient()
	var f []FileMetaData
	opt := &cos.BucketGetOptions{
		Prefix:    prefix + "/",
		Delimiter: "/",
		MaxKeys:   1000,
	}
	var marker string
	isTruncated := true
	for isTruncated {
		opt.Marker = marker
		v, _, err := client.Bucket.Get(context.Background(), opt)
		if err != nil {
			return nil, err
		}
		for _, content := range v.Contents {
			r, err := client.Object.Head(context.Background(), content.Key, nil)
			if err != nil {
				return nil, err
			}
			h := r.Header
			meta := FileMetaData{Filename: h.Get(COSFilename),
				Size:         content.Size,
				LastModified: content.LastModified,
				Uploader:     h.Get(COSMetaUploader),
			}
			f = append(f, meta)
		}
		isTruncated = v.IsTruncated
		marker = v.NextMarker
	}
	return f, nil

}

func Store(r io.Reader, filePath string, header *http.Header) error {

	if r == nil {
		r = strings.NewReader("")
	}
	client := makeClient()
	option := &cos.ObjectPutOptions{ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{XCosMetaXXX: header}}
	_, err := client.Object.Put(context.Background(), filePath, r, option)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Load(filePath string) (*cos.Response, error) {
	client := makeClient()
	resp, err := client.Object.Get(context.Background(), filePath, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

func Delete(filePath string) error {
	client := makeClient()
	_, err := client.Object.Delete(context.Background(), filePath)
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

func getGroup(r *http.Request) (group string, e error) {
	if !r.URL.Query().Has("groupId") {
		return "", fmt.Errorf("not enough argument")
	}

	groupId := r.URL.Query().Get("groupId")
	return groupId, nil

}

func checkUserInGroup(user, group string, tx *sql.Tx) error {
	s := `select user_id from user where user_id=? and group_id=?`
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

func genFilePath(group string, filename string) string {

	return group + "/" + filename
}
func check(user, group string) error {
	tx, err := DB.Begin()
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	if err = checkUserInGroup(user, group, tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return err
	}
	return nil
}

type FileRequest struct {
	Filename string `schema:"filename,required"`
	Group    string `schema:"group,required"`
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		HandleError(err, w, http.StatusBadRequest)
		return
	}
	user := w.Header().Get(UserHeader)

	var fr FileRequest

	d := schema.NewDecoder()
	err = d.Decode(&fr, r.PostForm)
	if err != nil {
		HandleError(err, w, http.StatusBadRequest)
		return
	}

	err = check(user, fr.Group)
	if err != nil {
		HandleError(err, w, http.StatusBadRequest)
		return
	}

	var response FileResponse
	err = Delete(genFilePath(fr.Group, fr.Filename))
	if err != nil {
		response.Reason = "Some errors occurred"
	} else {
		response.Ok = true
	}
	WriteJson(w, response)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	user := w.Header().Get(UserHeader)
	group, err := getGroup(r)
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

	if err = check(user, group); err != nil {
		HandleError(err, w, http.StatusBadRequest)
		return
	}
	h := http.Header{}
	h.Set(COSMetaUploader, user)
	h.Set(COSMetaGroup, group)
	h.Set(COSFilename, header.Filename)
	err = Store(uploadFile, genFilePath(group, header.Filename), &h)
	if err != nil {
		response.Reason = "Some problems occurred when uploading the files"
		WriteJson(w, response)
		return
	}
	response.Ok = true

	WriteJson(w, response)
}

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	user := w.Header().Get(UserHeader)
	q := r.URL.Query()
	if !q.Has("group") || !q.Has("filename") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	group, filename := q.Get("group"), q.Get("filename")
	err := check(user, group)
	if err != nil {
		HandleError(err, w, http.StatusBadRequest)
		return
	}
	var fr FileResponse
	var response *cos.Response
	if response, err = Load(genFilePath(group, filename)); err != nil {
		fr.Reason = "File does not exist"
		WriteJson(w, fr)
		return
	}
	w.Header().Set("Content-Type", response.Header.Get("Content-Type"))
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		HandleError(err, w, http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(b); err != nil {
		HandleError(err, w, http.StatusInternalServerError)
		return
	}

}

func ListFiles(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if !query.Has("group") {
		HandleError(fmt.Errorf("not enough argument"), w, http.StatusBadRequest)
		return
	}
	group := query.Get("group")
	user := w.Header().Get(UserHeader)
	err := check(user, group)
	if err != nil {
		HandleError(err, w, http.StatusBadRequest)
		return
	}

	d, err := List(group)
	var meta FileMetaResponse
	if err == nil {
		meta.Files = d
		meta.Ok = true
	}

	WriteJson(w, meta)

}
