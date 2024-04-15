package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"log"
	"server/auth"
	"server/config"
)

var c *config.Config

func main() {
	c = config.ReadConfig()
	initCookieStore(c.Auth.SessionKey)
	initDB()
	fmt.Println(c.Db.Port, c.Auth.SessionKey)

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
