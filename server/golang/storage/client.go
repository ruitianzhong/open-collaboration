package storage

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

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

func Store() error {
	client := makeClient()
	name := "test/objectPut.go"
	f := strings.NewReader("test")
	_, err := client.Object.Put(context.Background(), name, f, nil)
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
