package zapp_test

import (
	"github.com/ikasamt/zapp/zapp"
	"testing"
)

func Test_ParseEmailTemplateFile(t *testing.T) {
	zapp.EmailTemplateDir = `../testdata/email_templates/`
	fileName := "test1.txt"
	d:= map[string]interface{}{`name`: `John`, `age`: 21}
	actualSubject, actualBody := zapp.ParseEmailTemplateFile(fileName, d)

	expectedSubject := `Johnさんからメッセージが届きました`
	if actualSubject != expectedSubject {
		t.Errorf("got: %v\nwant: %v", actualSubject, expectedSubject)
	}

	expectedBody := `John（21）さんからのメッセージが届きました。

--------------------------
こんにちは。
--------------------------`
	if actualBody != expectedBody {
		t.Errorf("got: %v\nwant: %v", actualBody, expectedBody)
	}


}

