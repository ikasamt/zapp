package zapp

import (
	"math/rand"
	"time"
)

func RandomString(length int) string {
	var sources = []string{"abcdefghijkpqrstuvwxyz", "2345679", "ABCDEFGHIJKLMNPQRSTUVWXYZ"}
	var retval []byte
	rand.Seed(time.Now().UnixNano())
	retval = make([]byte, 3*length, 3*length)
	cnt := 0
	for _, source := range sources {
		for i := 0; i < length; i++ {
			retval[cnt] = source[rand.Intn(len(source))]
			cnt += 1
		}
	}
	return string(retval)
}
