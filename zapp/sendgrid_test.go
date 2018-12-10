package zapp_test

import (
	"log"
	"testing"

	"github.com/ikasamt/zapp/zapp"
)

func Test_SendgridMulti(t *testing.T) {
	sendgridAPIKey := `----------------`
	from := `--------`
	tos := []string{`--------`, `--------`}
	err := zapp.SendEmailMultiTo(sendgridAPIKey, from, tos, `testsubject`, `testbody`)
	err2 := zapp.SendEmail(sendgridAPIKey, from, `ikasamt+to@gmail.com`, `test`, `test`)
	log.Println(err)
	log.Println(err2)
}
