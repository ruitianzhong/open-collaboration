package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"server/auth"
	"server/chat"
	"server/config"
	"server/storage"
	"server/translation"
)

var c *config.Config

func main() {
	c = config.ReadConfig()
	initCookieStore(c.Auth.SessionKey)
	initDB()
	translation.InitTranslation(c.Translation.Appid, c.Translation.Key)
	storage.InitStorage(c.Storage.SecretId, c.Storage.SecretKey, c.Storage.BucketPath)
	chat.Init(c.Chat.AppId, c.Chat.Key, c.Chat.AdminId)
	//chat.AddUserAccount()

	r := mux.NewRouter()
	r.HandleFunc("/auth/login", auth.Login).Methods(http.MethodPost)
	r.HandleFunc("/api/translate", translation.Serve).Methods(http.MethodPost)
	r.HandleFunc("/chat/refresh", chat.RefreshToken).Methods(http.MethodPost)
	r.HandleFunc("/files/upload", storage.UploadFile).Methods(http.MethodPost)
	r.HandleFunc("/files/delete", storage.DeleteFile).Methods(http.MethodPost)
	r.HandleFunc("/files/download", storage.DownloadFile).Methods(http.MethodGet)
	r.HandleFunc("/files/list", storage.ListFiles).Methods(http.MethodGet)
	r.HandleFunc("/docs/list", storage.ListDocs).Methods(http.MethodGet)
	r.HandleFunc("/docs/update", storage.UpdateDocs).Methods(http.MethodPost)
	r.HandleFunc("/docs/new", storage.AddNewDocs).Methods(http.MethodPost)
	r.HandleFunc("/docs/delete", storage.DeleteDocs).Methods(http.MethodPost)
	r.HandleFunc("/docs/byid", storage.GetDocById).Methods(http.MethodGet)
	r.HandleFunc("/auth/logout", auth.Logout).Methods(http.MethodPost)
	r.HandleFunc("/auth/fetch-user-info", auth.FetchUserInfo).Methods(http.MethodGet)
	r.Use(auth.AuthMiddleware)
	err := http.ListenAndServe(":"+c.Port, r)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func initCookieStore(sessionKey string) {
	store := sessions.NewCookieStore([]byte(sessionKey))
	auth.InitCookieStore(store)
}

func ConnectionDriverAndPath(address, port, dbName, username, password string) (driverName string, connectionPath string) {
	sqlConnectionPath := username + ":" + password + "@(" + address + ":" + port + ")/" + dbName + "?parseTime=true&interpolateParams=true"
	driverName = "mysql"
	return driverName, sqlConnectionPath

}

func initDB() {
	if c.Db == nil {
		return
	}
	driverName, path := ConnectionDriverAndPath(c.Db.Address, c.Db.Port,
		c.Db.DbName, c.Db.UserName,
		c.Db.Password)

	DB, err := sql.Open(driverName, path)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	auth.InitDB(DB)
	chat.InitDB(DB)
	storage.InitDB(DB)
}
