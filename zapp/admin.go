package zapp

import (
	"strings"

	"github.com/gin-gonic/gin"
)

var DefaultActionName = `list`

func ExtractControllerActionName(path string, prefix string) (controllerName string, actionName string) {
	paths := strings.Split(path, `/`)

	// first string is blacnk
	if paths[0] == `` {
		paths = paths[1:]
	}

	// first strings equals prefix skip it
	// ex paths=admin/user/new and prefix= admin
	if paths[0] == prefix {
		paths = paths[1:]
	}

	controllerName = paths[0]
	if len(paths) == 1 {
		actionName = DefaultActionName
	} else {
		actionName = paths[1]
		if paths[1] == `` {
			actionName = DefaultActionName
		}
	}
	return controllerName, actionName
}

//
func Render(c *gin.Context, dir string, context map[string]interface{}, templateName ...string) error {

	// context に controllerName, actionName を追加する
	controllerName, actionName := ExtractControllerActionName(c.Request.URL.Path, dir)
	context[`controllerName`] = controllerName
	if templateName != nil {
		// override actionName when it sets
		actionName = templateName[0]
	}
	context[`actionName`] = actionName

	// ページネーション
	context[`pager`] = NewPager(c, context[`total_count`])

	// テンプレートファイルがなければ初期ファイルをわたす
	RenderJade(c, dir, controllerName, actionName, context)
	return nil
}

//
func RenderAdmmin(c *gin.Context, context map[string]interface{}, templateName ...string) error {
	return Render(c, `admin`, context, templateName...)
}
