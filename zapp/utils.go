package zapp

import (
	"fmt"
	"math/rand"
	"net/url"
	"regexp"
	"strings"
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
			cnt++
		}
	}
	return string(retval)
}

func RandomDigitString(length int) string {
	var source = "012345679"
	rand.Seed(time.Now().UnixNano())
	retval := make([]byte, length, length)
	cnt := 0
	for i := 0; i < length; i++ {
		pos := rand.Intn(len(source))
		retval[cnt] = source[pos]
		cnt++
	}
	return string(retval)
}


func Hashtag2Link(text string, href string) string {
	rep := regexp.MustCompile(`(#[^\s]*)`)
	matches := rep.FindAllStringSubmatch(text, -1)
	for _, m := range matches {
		s := m[0]
		e := url.QueryEscape(s)
		text = strings.Replace(text, m[0], fmt.Sprintf("<a href=%s%s>%s</a>", href, e, s), -1)
	}
	return text
}
