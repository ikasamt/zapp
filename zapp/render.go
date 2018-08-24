package zapp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/Joker/jade"
	humanize "github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"github.com/golang/go/src/html/template"
)

func baseFuncMap() template.FuncMap {
	return template.FuncMap{
		"Comma": func(value interface{}) string {
			switch t := value.(type) {
			case int:
				v := value.(int)
				return humanize.Comma(int64(v))
			case int64:
				v := value.(int64)
				return humanize.Comma(v)
			case float64:
				v := value.(float64)
				return humanize.Comma(int64(v))
			default:
				return fmt.Sprintf(`unknown comma error(%s)`, t)
			}
		},
		"FormatNumber": func(value float64) string {
			return fmt.Sprintf("%.2f", value)
		},
		"FormatJST": func(value time.Time) string {
			return value.Format("2006/01/02 15:04:05")
		},
	}
}

// ConvertJadeToHTML は、jadeファイルを読みHTML文字列にする
func ConvertJadeToHTML(templateFilename string) (html string, err error) {
	jadeBytes, err := ioutil.ReadFile(templateFilename)
	if err != nil {
		log.Println(err)
		return "", err
	}

	html, err = jade.Parse(templateFilename, string(jadeBytes))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return html, nil
}

//
func RenderJade(c *gin.Context, layoutFilename string, templateFilename string, context map[string]interface{}) error {
	fn := TemplateDir + "/" + layoutFilename + ".jade"
	layoutHTML, err := ConvertJadeToHTML(fn)
	if err != nil {
		log.Println(err)
		return err
	}

	fn = TemplateDir + "/" + templateFilename + ".jade"
	contentHTML, err := ConvertJadeToHTML(fn)
	if err != nil {
		log.Println(err)
		return err
	}

	return ExecuteTemplate(c, layoutHTML, contentHTML, context)
}

func executeTemplateToHTML(templateFilename string, context map[string]interface{}) (template.HTML, error) {
	outPut := new(bytes.Buffer)
	includeHTML, _ := ConvertJadeToHTML(templateFilename)
	partialHTML, _ := template.New(templateFilename).Parse(includeHTML)
	partialHTML.Execute(outPut, context)
	return template.HTML(outPut.String()), nil
}

//
func ExecuteTemplate(c *gin.Context, layoutHTML string, contentHTML string, context map[string]interface{}) error {

	// session由来の変数を当てはめる
	context["flashMessage"] = GetFlashMessage(c)

	// テンプレート関数
	funcMap := baseFuncMap()
	funcMap[`include`] = func(includePath string) (template.HTML, error) {
		return executeTemplateToHTML("templates/"+includePath+".jade", context)
	}

	// 変数適用
	tmpl, err := template.New("layout").Funcs(funcMap).Parse(layoutHTML)
	tmpl.New("tmpl").Parse(contentHTML)
	if err != nil {
		return err
	}

	err = tmpl.Execute(c.Writer, context)
	if err != nil {
		return err
	}
	return nil
}
