package auth

import (
	"log"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "dm-session")
	session.Values["authenticated"] = false
	w.WriteHeader(http.StatusOK)
}

type UserInfo struct {
	UserId  string `json:"userId"`
	GroupId string `json:"groupId"`
}

func FetchUserInfo(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "dm-session")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		log.Println("reject")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	s1, ok1 := session.Values["username"].(string)
	s2, ok2 := session.Values["group_id"].(string)
	if !ok1 || !ok2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u := UserInfo{
		UserId:  s1,
		GroupId: s2,
	}
	WriteJson(w, u)

}
