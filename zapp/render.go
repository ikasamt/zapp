package zapp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"time"

	"html/template"

	"github.com/Joker/jade"
	humanize "github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
)

var TemplateDir = `templates`

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
		"FormatJST": func(value time.Time, fmt string) string {
			return value.Format(fmt)
		},
		"FormatJSTTime": func(value time.Time) string {
			return value.Format("15:04")
		},
		"FormatJSTDay": func(value time.Time) string {
			return value.Format("2006/01/02")
		},
		"FormatJSTM": func(value time.Time) string {
			return value.Format("2006/01/02 15:04")
		},
		"FormatJSTL": func(value time.Time) string {
			return value.Format("2006/01/02 15:04:05")
		},
		"safehtml": func(text string) template.HTML {
			return template.HTML(text)
		},
		"nl2br": func(text string) template.HTML {
			t := strings.Replace(text, "\n", `<br/>`, -1)
			return template.HTML(t)
		},
		"ToCamel": func(text string) template.HTML {
			t := strcase.ToCamel(text)
			return template.HTML(t)
		},
		"ToSnake": func(text string) template.HTML {
			t := strcase.ToSnake(text)
			return template.HTML(t)
		},
	}
}

// ConvertJadeToHTML は、jadeファイルを読みHTML文字列にする
func ConvertJadeToHTML(templateFilename string) (html string, err error) {
	jadeBytes, err := ioutil.ReadFile(templateFilename)
	if err != nil {
		return "", err
	}

	html, err = jade.Parse(templateFilename, string(jadeBytes))
	if err != nil {
		return "", err
	}
	return html, nil
}

//
func RenderJade(c *gin.Context, dirName string, controllerName string, actionName string, context map[string]interface{}) error {
	fn := filepath.Join(TemplateDir, dirName, "layout.jade")
	layoutHTML, err := ConvertJadeToHTML(fn)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, dir := range [2]string{controllerName, `scaffold`} {
		fn = filepath.Join(TemplateDir, dirName, dir, actionName+".jade")
		contentHTML, err := ConvertJadeToHTML(fn)
		if err == nil {
			return ExecuteTemplate(c, dirName, controllerName, layoutHTML, contentHTML, context)
		}
		// log.Println(err)
	}

	// 両ファイルともエラーなら
	return err
}

func executeTemplateToHTML(templateFilename string, funcMap template.FuncMap, context map[string]interface{}) (template.HTML, error) {
	outPut := new(bytes.Buffer)
	includeHTML, err := ConvertJadeToHTML(templateFilename)
	if err != nil {
		log.Println(err)
		return template.HTML(``), err
	}
	partialHTML, err := template.New(templateFilename).Funcs(funcMap).Parse(includeHTML)
	if err != nil {
		log.Println(err)
		return template.HTML(``), err
	}
	err = partialHTML.Execute(outPut, context)
	if err != nil {
		log.Println(err)
		return template.HTML(``), err
	}
	return template.HTML(outPut.String()), nil
}

//
func ExecuteTemplate(c *gin.Context, dirName string, controllerName string, layoutHTML string, contentHTML string, context map[string]interface{}) error {

	// session由来の変数を当てはめる
	context["flashMessage"] = GetFlashMessage(c)

	// for GAE
	_, isDev := c.Get("__is_dev")
	if isDev {
		context[`__is_dev`] = true
	}

	me, ok := c.Get("me")
	if ok {
		context[`me`] = me
	}

	// テンプレート関数
	funcMap := baseFuncMap()
	funcMap[`link_to`] = func(name string) (template.HTML, error) {
		context[`__link_to_controllerName`] = name
		context[`__link_to_label`] = strcase.ToCamel(name)
		return executeTemplateToHTML(TemplateDir+"/_link_to.jade", funcMap, context)
	}
	funcMap[`include`] = func(includePath string) (template.HTML, error) {
		retval, err := executeTemplateToHTML(TemplateDir+"/"+dirName+"/"+controllerName+"/"+includePath+".jade", funcMap, context)
		if err == nil {
			return retval, nil
		}
		retval, err = executeTemplateToHTML(TemplateDir+"/"+includePath+".jade", funcMap, context)
		if err == nil {
			return retval, nil
		}
		log.Println(err)
		return retval, err
	}

	// 変数適用
	tmpl, err := template.New("layout").Funcs(funcMap).Parse(layoutHTML)
	tmpl.New("tmpl").Parse(contentHTML)
	if err != nil {
		log.Println(err)
		return err
	}

	err = tmpl.Execute(c.Writer, context)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func RenderDirect(c *gin.Context, templateFilename string, context interface{}) error {

	// 個別ファイル読み込み
	data, err := ioutil.ReadFile(TemplateDir + "/" + templateFilename + ".jade")
	if err != nil {
		log.Println(err)
		return err
	}
	contentmpl, err := jade.Parse("template", string(data))
	if err != nil {
		log.Println(err)
		return err
	}

	// 変数適用
	tmpl, err := template.New("contentmpl").Parse(contentmpl)
	if err != nil {
		log.Println(err)
		return err
	}

	err = tmpl.Execute(c.Writer, context)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
