package translation

import (
	"encoding/json"
	"github.com/gorilla/schema"
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
	ZH   = language("zh")
	YUE  = language("yue")
	EN   = language("en")
	AUTO = language("auto")
)

func Translate(from, to language, query string) (*Response, error) {
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
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)

	response := Response{}
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type TranslateRequest struct {
	Source string `schema:"source,required"`
	Target string `schema:"target,required"`
}

type TranslateResponse struct {
	Dst string `json:"dst"`
}

func Serve(w http.ResponseWriter, r *http.Request) {
	decoder := schema.NewDecoder()
	err := r.ParseForm()
	if err != nil {
		HandleError(err, w, http.StatusBadRequest)
		return
	}
	var tr TranslateRequest
	err = decoder.Decode(&tr, r.PostForm)
	if err != nil {
		HandleError(err, w, http.StatusBadRequest)
		return
	}
	target := language(tr.Target)
	switch target {
	case EN:
	case ZH:
	case YUE:
		break
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := Translate(AUTO, target, tr.Source)
	if err != nil {
		HandleError(err, w, http.StatusBadRequest)
		return
	}
	translateResp := TranslateResponse{}
	s := ""
	for _, v := range resp.Result {
		s += v.Dst
	}
	translateResp.Dst = s
	WriteJson(w, translateResp)
}
