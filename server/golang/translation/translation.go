package translation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const ApiPath = "https://fanyi-api.baidu.com/api/trans/vip/translate"

var (
	KEY   string
	APPID string
)

// Response According to https://api.fanyi.baidu.com/product/113#
type Response struct {
	From   string   `json:"from"`
	To     string   `json:"to"`
	Result []Result `json:"trans_result"`
}

type Result struct {
	Src       string `json:"src"`
	Dst       string `json:"dst"`
	ErrorCode string `json:"error_code"`
}

func InitTranslation(appid, key string) {
	APPID = appid
	KEY = key
}

type language string

const (
	ZH  = language("zh")
	YUE = language("yue")
	EN  = language("en")
)

func Translate(from, to language, query string) (string, error) {
	dataUrl := url.Values{}

	salt := generateSecureRandom()
	dataUrl.Set("q", query)
	dataUrl.Set("from", string(from))
	dataUrl.Set("to", string(to))
	dataUrl.Set("appid", APPID)
	dataUrl.Set("sign", sign(query, salt, APPID, KEY))
	dataUrl.Set("salt", salt)

	request, err := http.NewRequest("POST", ApiPath, strings.NewReader(dataUrl.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)

	response := Response{}
	err = json.Unmarshal(respBytes, &response)
	fmt.Println(string(respBytes))
	if err != nil {
		return "", err
	}

	fmt.Printf("%v\n", response)

	return "", nil
}
