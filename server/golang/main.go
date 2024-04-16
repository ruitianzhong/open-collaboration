package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"log"
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
	//_, err := translation.Translate(translation.ZH, translation.EN, "你好，世界")
	//if err != nil {
	//	fmt.Println(err)
	//}
	fmt.Println(c.Translation.Appid, c.Translation.Key)
	//storage.Store()
	//storage.Load()
	//storage.Delete()
	chat.Init(c.Chat.AppId, c.Chat.Key, c.Chat.AdminId)
	chat.AddUserAccount()
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
}
