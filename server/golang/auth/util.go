package auth

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJson(w http.ResponseWriter, v any) (bool, []byte) {
	marshal, err := json.Marshal(v)
	if err != nil {
		HandleError(err, w, http.StatusInternalServerError)
		return true, nil
	}
	_, err = w.Write(marshal)
	if err != nil {
		HandleError(err, w, http.StatusInternalServerError)
		return true, nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return false, marshal
}
func HandleError(e error, w http.ResponseWriter, statusCode int) {
	log.Println(e.Error())
	w.WriteHeader(statusCode)
}
