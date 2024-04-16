package auth

import (
	"database/sql"
	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore
var (
	DB *sql.DB
)

func InitDB(db *sql.DB) {
	DB = db

}

func InitCookieStore(s *sessions.CookieStore) {
	store = s
}
