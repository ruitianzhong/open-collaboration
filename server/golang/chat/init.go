package chat

import (
	"log"
	"strconv"
)

var (
	appid int
	key   string
	admin string
)

func Init(id, appKey, adminId string) {
	key = appKey
	var err error
	appid, err = strconv.Atoi(id)
	admin = adminId
	if err != nil {
		log.Fatal(err)
	}

}
