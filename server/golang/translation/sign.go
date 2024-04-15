package translation

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	unsaferand "math/rand"
	"strconv"
)

func sign(query, salt, appid, key string) string {
	s := appid + query + salt + key
	data := []byte(s)
	b := md5.Sum(data)
	return fmt.Sprintf("%x", b)
}

func generateSecureRandom() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return strconv.Itoa(unsaferand.Int())
	}
	return fmt.Sprintf("%x", b)
}
