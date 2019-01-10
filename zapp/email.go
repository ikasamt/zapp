package zapp

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var (
	EmailTemplateDir = "email_templates"
	EmailTemplateLineSepChar = "\n"
)

func ParseEmailTemplateFile(fileName string, data map[string]interface{}) (subject, body string, err error) {
	// テンプレートから文字列を生成する
	var buf bytes.Buffer
	templateFilename := filepath.Join(EmailTemplateDir, fileName)
	bytes, err := ioutil.ReadFile(templateFilename)
	if err != nil {
		return
	}

	// 変数適用
	tmpl, err := template.New("tmpl").Parse(string(bytes))
	if err != nil {
		return
	}
	if err = tmpl.Execute(&buf, data); err != nil {
		return
	}

	// １行目を題名、２行目以降を本文とする
	s := buf.String()
	lines := strings.Split(s, EmailTemplateLineSepChar)
	subject = lines[0]                                       // first line
	body = strings.Join(lines[1:], EmailTemplateLineSepChar) // rest of all
	return
}

