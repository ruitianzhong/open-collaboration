package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)
import "github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"

const URL = "https://console.tim.qq.com/v4/im_open_login_svc/account_import?"

type AddAccountRequest struct {
	UserID string `json:"UserID"`
	Nick   string `json:"Nick"`
}

func genPairs(key, value string) string {
	return key + "=" + value
}

func AddUserAccount() error {
	s := make([]string, 5)

	s[0] = genPairs("sdkappid", strconv.Itoa(appid))
	s[1] = genPairs("identifier", admin)

	sig, err := tencentyun.GenUserSig(appid, key, admin, 7*24*60*60*30)

	if err != nil {
		log.Println(err)
		return err
	}
	s[2] = genPairs("usersig", sig)
	s[3] = genPairs("random", strconv.Itoa(rand.Int()))

	s[4] = genPairs("contenttype", "json")
	param := ""
	for i := 0; i < len(s); i++ {
		param += s[i]
		if i != len(s)-1 {
			param += "&"
		}
	}

	req := AddAccountRequest{UserID: "42", Nick: "RT"}
	client := http.Client{}
	b, err := json.Marshal(req)
	fmt.Println(string(b))
	if err != nil {
		log.Println(err)
		return err
	}

	response, err := client.Post(URL+param, "application/json", bytes.NewReader(b))
	if err != nil {
		log.Println(err)

		return err
	}

	b, err = ioutil.ReadAll(response.Body)
	if err != nil {

		return err
	}

	fmt.Println(string(b))
	fmt.Println(param)
	return nil

}

type RefreshRequest struct {
	Sig string `schema:"sig,required"`
}

type RefreshResponse struct {
	Ok  bool   `json:"ok"`
	Sig string `json:"sig"`
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	id := w.Header().Get("X_USER_ID")
	if id == "" {
		return
	}
	w.Header().Del("X_USER_ID")

	err := r.ParseForm()

	if err != nil {
		HandleError(err, w, http.StatusBadRequest)
		return
	}

	d := schema.NewDecoder()
	refreshRequest := RefreshRequest{}
	err = d.Decode(&refreshRequest, r.PostForm)
	if err != nil {
		HandleError(err, w, http.StatusBadRequest)
		return
	}

	err = tencentyun.VerifyUserSig(uint64(appid), key, id, refreshRequest.Sig, time.Now())
	var resp RefreshResponse
	if err != nil {
		sig, err := tencentyun.GenUserSig(appid, key, id, 7*24*60*60)
		if err != nil {
			HandleError(err, w, http.StatusBadRequest)
			return
		}
		resp.Sig = sig
	} else {
		resp.Ok = true
	}
	WriteJson(w, resp)
}
