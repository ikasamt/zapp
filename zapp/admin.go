package zapp

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var AdminPrefix = `/adimn/`
var DefaultActionName = `list`
var DefaultPerPage = 15
var UnknownTotalCount = 1000
var TemplateDir = `templates`

func extractControllerActionName(c *gin.Context) (string, string) {
	pathURL := strings.Replace(c.Request.URL.Path, AdminPrefix, "", 1)
	paths := strings.Split(pathURL, `/`)
	controllerName := paths[0]
	actionName := paths[1]
	if actionName == "" {
		actionName = DefaultActionName
	}
	return controllerName, actionName
}

func findTemplateFilename(controllerName string, actionName string) (templateFilename string) {
	if _, err := os.Stat(TemplateDir + "/admin/" + controllerName + "/" + actionName + ".jade"); os.IsNotExist(err) {
		templateFilename = "admin/scaffold/" + actionName
	} else {
		templateFilename = "admin/" + controllerName + "/" + actionName
	}
	return templateFilename
}

func renderAdmin(c *gin.Context, context map[string]interface{}, templateName ...string) error {

	// context に controllerName, actionName を追加する
	controllerName, actionName := extractControllerActionName(c)
	context[`controllerName`] = controllerName
	if templateName != nil {
		// override actionName when it sets
		actionName = templateName[0]
	}
	context[`actionName`] = actionName

	// ページネーション
	context[`pager`] = NewPager(c, context[`total_count`])

	// テンプレートファイルがなければ初期ファイルをわたす
	templateFilename := findTemplateFilename(controllerName, actionName)
	RenderJade(c, `admin/layout.jade`, templateFilename, context)
	return nil
}
