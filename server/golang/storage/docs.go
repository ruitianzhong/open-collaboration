package storage

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type AddDocsRequest struct {
	Markdown string `json:"markdown"`
	Group    string `json:"group"`
	Title    string `json:"title"`
	Id       string `json:"id"`
}
type DocsResponse struct {
	Id       string `json:"id"`
	Ok       bool   `json:"ok"`
	Markdown string `json:"markdown"`
	Title    string `json:"title"`
}

var (
	COSTitle      = "x-cos-meta-title"
	COSAuthor     = "x-cos-meta-author"
	COSCreateTime = "x-cos-meta-created-time"
)

type UpdateDocsRequest struct {
	Markdown string `json:"markdown"`
	Group    string `json:"group"`
	Title    string `json:"title"`
	Id       string `json:"id"`
}

func genDocsPath(id, group, user string) string {
	return "docs/" + group + "/" + id
}

func ParseRequest2Json(r *http.Request) (*AddDocsRequest, error) {
	if r.ContentLength <= 0 || r.ContentLength > 1024*1024*10 {
		return nil, errors.New("request too long")
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	var adr AddDocsRequest

	if err = json.Unmarshal(b, &adr); err != nil {
		return nil, err
	}

	return &adr, nil
}

func AddNewDocs(w http.ResponseWriter, r *http.Request) {

	adr, err := ParseRequest2Json(r)
	if err != nil {
		HandleError(err, w, http.StatusBadRequest)
		return
	}
	if adr.Title == "" {
		WriteJson(w, DocsResponse{})
		return
	}
	user := w.Header().Get(UserHeader)
	if err = check(user, adr.Group); err != nil {
		HandleError(err, w, http.StatusUnauthorized)
		return
	}

	next, err := allocateID()
	if err != nil {
		HandleError(err, w, http.StatusInternalServerError)
		return
	}
	header := &http.Header{}
	header.Set(COSTitle, adr.Title)
	header.Set(COSAuthor, user)
	header.Set(COSCreateTime, time.Now().Format("2006-01-02 15:01:05"))
	header.Set(COSMetaGroup, adr.Group)
	err = Store(strings.NewReader(adr.Markdown), genDocsPath(strconv.FormatInt(next, 10), adr.Group, user), header)

	var response DocsResponse
	if err == nil {
		response.Id = strconv.FormatInt(next, 10)
		response.Ok = true
	}
	WriteJson(w, response)
}

func UpdateDocs(w http.ResponseWriter, r *http.Request) {
	adr, err := ParseRequest2Json(r)
	if err != nil {
		HandleError(err, w, http.StatusBadRequest)
		return
	}
	if adr.Title == "" {
		WriteJson(w, DocsResponse{})
		return
	}
	user := w.Header().Get(UserHeader)
	if err = check(user, adr.Group); err != nil {
		HandleError(err, w, http.StatusUnauthorized)
		return
	}
	response := &DocsResponse{}
	if adr.Id == "" {
		WriteJson(w, response)
		return
	}

	client := makeClient()
	key := genDocsPath(adr.Id, adr.Group, user)

	resp, err := client.Object.Head(context.Background(), key, nil)

	if err != nil {
		log.Println(err)
		WriteJson(w, response)
		return
	}
	header := http.Header{}
	header.Set(COSTitle, adr.Title)
	header.Set(COSAuthor, user)
	header.Set(COSCreateTime, resp.Header.Get(COSCreateTime))
	header.Set(COSMetaGroup, adr.Group)
	if err = Store(strings.NewReader(adr.Markdown), key, &header); err == nil {
		response.Ok = true
	}
	WriteJson(w, response)
}

func DeleteDocs(w http.ResponseWriter, r *http.Request) {
	adr, err := ParseRequest2Json(r)
	if err != nil {
		HandleError(err, w, http.StatusBadRequest)
		return
	}
	user := w.Header().Get(UserHeader)
	if err = check(user, adr.Group); err != nil {
		HandleError(err, w, http.StatusUnauthorized)
		return
	}
	response := DocsResponse{}
	if err = Delete(genDocsPath(adr.Id, adr.Group, user)); err == nil {
		response.Ok = true
	}
	WriteJson(w, response)
}

type DocsMetaData struct {
	Title        string `json:"title"`
	Creator      string `json:"creator"`
	LastModified string `json:"lastModified"`
	CreatedTime  string `json:"createdTime"`
	Id           string `json:"id"`
}

type DocsMetaDataResponse struct {
	Docs []DocsMetaData `json:"docs"`
}

func gatherDocs(group string) ([]DocsMetaData, error) {
	client := makeClient()
	var f []DocsMetaData
	opt := &cos.BucketGetOptions{
		Prefix:    "docs/" + group + "/",
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
			meta := DocsMetaData{
				LastModified: content.LastModified,
				Creator:      h.Get(COSAuthor),
				CreatedTime:  h.Get(COSCreateTime),
				Title:        h.Get(COSTitle),
				Id:           strings.Split(content.Key, "/")[2],
			}
			f = append(f, meta)
		}
		isTruncated = v.IsTruncated
		marker = v.NextMarker
	}
	return f, nil

}

func ListDocs(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("group") {
		HandleError(errors.New("not enough argument"), w, http.StatusBadRequest)
		return
	}
	group := r.URL.Query().Get("group")
	user := w.Header().Get(UserHeader)
	if err := check(user, group); err != nil {
		HandleError(err, w, http.StatusUnauthorized)
		return
	}
	var docs DocsMetaDataResponse
	d, err := gatherDocs(group)
	if err != nil {
		log.Println(err)
	}
	docs.Docs = d
	WriteJson(w, docs)
}

func allocateID() (int64, error) {

	tx, err := DB.Begin()
	if err != nil {
		return 0, err
	}
	s1 := `SELECT count+1 from id_allocation where topic='docs' for update`
	var next int64
	if err = tx.QueryRow(s1).Scan(&next); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	s2 := `Update id_allocation set count=count+1 where topic='docs'`

	if _, err = tx.Exec(s2); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return next, nil
}

func GetDocById(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if !query.Has("group") || !query.Has("id") {
		HandleError(errors.New("not enough argument"), w, http.StatusBadRequest)
		return
	}

	user, group, id := w.Header().Get(UserHeader), query.Get("group"), query.Get("id")
	if err := check(user, group); err != nil {
		HandleError(err, w, http.StatusUnauthorized)
		return
	}
	resp, err := Load(genDocsPath(id, group, user))
	if err != nil {
		HandleError(err, w, http.StatusNotFound)
		return
	}
	var docsResponse DocsResponse
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		WriteJson(w, docsResponse)
		return
	}
	docsResponse.Markdown = string(b)
	docsResponse.Title = resp.Header.Get(COSTitle)
	docsResponse.Ok = true
	WriteJson(w, docsResponse)
}
