package zapp

import (
	"strings"

	"github.com/gin-gonic/gin"
)

var AdminPrefix = `/admin/`
var DefaultActionName = `list`
var DefaultPerPage = 15
var UnknownTotalCount = 1000
var TemplateDir = `templates`

func ExtractControllerActionName(c *gin.Context) (string, string) {
	pathURL := strings.Replace(c.Request.URL.Path, AdminPrefix, "", 1)
	paths := strings.Split(pathURL, `/`)
	controllerName := paths[0]

	actionName := paths[1]
	if paths[1] == `` {
		actionName = DefaultActionName
	}
	return controllerName, actionName
}

func RenderAdmin(c *gin.Context, context map[string]interface{}, templateName ...string) error {

	// context に controllerName, actionName を追加する
	controllerName, actionName := ExtractControllerActionName(c)
	context[`controllerName`] = controllerName
	if templateName != nil {
		// override actionName when it sets
		actionName = templateName[0]
	}
	context[`actionName`] = actionName

	// ページネーション
	context[`pager`] = NewPager(c, context[`total_count`])

	// テンプレートファイルがなければ初期ファイルをわたす
	RenderJade(c, `admin`, controllerName, actionName, context)
	return nil
}
