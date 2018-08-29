package zapp

import (
	"crypto/md5"
	"encoding/hex"
)

func calcMd5(bytes []byte) string {
	hasher := md5.New()
	hasher.Write(bytes)
	retval := hex.EncodeToString(hasher.Sum(nil))
	return retval
}

func HashPassword(salt string, password string) string {
	return calcMd5([]byte(salt + password))
}
