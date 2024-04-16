package chat

import (
	"database/sql"
	"github.com/gorilla/sessions"
	"log"
	"strconv"
)

var (
	appid int
	key   string
	admin string
	DB    *sql.DB
)
var store *sessions.CookieStore

func Init(id, appKey, adminId string) {
	key = appKey
	var err error
	appid, err = strconv.Atoi(id)
	admin = adminId
	if err != nil {
		log.Fatal(err)
	}

}

func InitDB(db *sql.DB) {
	DB = db
}

func InitCookieStore(s *sessions.CookieStore) {
	store = s
}
