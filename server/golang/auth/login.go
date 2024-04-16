package auth

import (
	"github.com/gorilla/schema"
	"net/http"
)

type LoginForm struct {
	Password string `schema:"password,required"`
	UserId   string `schema:"userid,required"`
}
type LoginResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	decoder := schema.NewDecoder()
	err := r.ParseForm()
	if err != nil {
		HandleError(err, w, http.StatusBadRequest)
		return
	}
	var lf LoginForm
	err = decoder.Decode(&lf, r.PostForm)
	if err != nil {
		HandleError(err, w, http.StatusBadRequest)
		return
	}
	db := DB
	s := "SELECT password from user where id=?"
	lr := LoginResponse{}
	var passwd string
	if err = db.QueryRow(s, lf.UserId).Scan(&passwd); err != nil || passwd != lf.Password {
		lr.Code = "100"
		WriteJson(w, lr)
		return
	}
	session, _ := store.Get(r, "dm-session")
	session.Values["authenticated"] = true
	session.Values["username"] = lf.UserId
	session.Values["auth_level"] = 0
	lr.Code = "200"
	if err = session.Save(r, w); err != nil {
		HandleError(err, w, http.StatusInternalServerError)
		return
	}
	WriteJson(w, lr)
}
