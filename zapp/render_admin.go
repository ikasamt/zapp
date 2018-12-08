package zapp

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var DefaultActionName = `index`

func ExtractControllerActionName(path string, prefix string) (controllerName string, actionName string) {

	tmp := strings.TrimPrefix(path, "/"+prefix)
	if tmp == `/` {
		controllerName = `top`
		actionName = `index`
		return
	}

	tmp2 := strings.Split(tmp, `.`)
	tmp3 := strings.Trim(tmp2[0], `/`)
	paths := strings.Split(tmp3, `/`)

	switch len(paths) {
	case 1:
		// paths == [foo], path='/foo'
		controllerName = paths[0]
		actionName = `index`
	case 2, 3:
		// paths == [foo bar], path='/foo/bar'
		controllerName = paths[0]
		actionName = paths[1]
	}

	return controllerName, actionName
}

//
func Render(c *gin.Context, dir string, context map[string]interface{}, templateName ...string) error {

	var controllerName, actionName string

	// context に controllerName, actionName を追加する
	controllerName, actionName = ExtractControllerActionName(c.Request.URL.Path, dir)
	if templateName != nil {
		// override actionName when it sets
		tmp := strings.Split(templateName[0], `/`)
		if len(tmp) == 1 {
			actionName = tmp[0]
		} else {
			controllerName = tmp[0]
			actionName = tmp[1]
		}
	}
	context[`controllerName`] = controllerName
	context[`actionName`] = actionName

	context[`CalendarSupport`] = CalendarSupport{Now: time.Now()}

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
