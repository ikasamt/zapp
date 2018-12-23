package zapp

import (
	"bytes"
	"log"
	"path/filepath"
	"strings"
	"html/template"
)

var (
	EmailTemplateDir = "email_templates"
	EmailTemplateLineSepChar = "\n"
)

func ParseEmailTemplateFile(fileName string, data map[string]interface{}) (subject, body string) {
	// テンプレートから文字列を生成する
	var buf bytes.Buffer
	fn := filepath.Join(EmailTemplateDir, fileName)
	log.Println("reading: ", fn)
	t := template.Must(template.ParseFiles(fn))
	if err := t.ExecuteTemplate(&buf, fn, data); err != nil {
		log.Fatal(err)
	}

	// １行目を題名、２行目以降を本文とする
	s := buf.String()
	lines := strings.Split(s, EmailTemplateLineSepChar)
	subject = lines[0]                                       // first line
	body = strings.Join(lines[1:], EmailTemplateLineSepChar) // rest of all
	return
}

