package clefs

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
	"github.com/ikasamt/zapp/zapp"
	"google.golang.org/appengine/search"

	"google.golang.org/appengine"
)

// List
func AdminAnythingListHandler(c *gin.Context) {
	instance := Anything{}
	instance.Set(c)
	totalCount, instances := searchAnythings(c, instance)
	context := map[string]interface{}{"instances": instances, "instance": instance, "total_count": totalCount}
	zapp.Render(c, `admin`, context)
}

// New
func AdminAnythingNewHandler(c *gin.Context) {
	instance := Anything{}
	context := map[string]interface{}{"instance": instance}
	zapp.Render(c, `admin`, context)
}

// Edit
func AdminAnythingEditHandler(c *gin.Context) {
	instance, _ := getAnything(c)
	context := map[string]interface{}{"instance": instance}
	zapp.Render(c, `admin`, context)
}

// Show
func AdminAnythingShowHandler(c *gin.Context) {
	instance, err := getAnything(c)
	if err != nil || instance.ID == 0 {
		log.Println(err)
	}
	context := map[string]interface{}{"instance": instance}
	zapp.Render(c, `admin`, context)
}

// Create
func AdminAnythingCreateHandler(c *gin.Context) {
	// DB接続を取得
	db, _ := NewGormDB(c)
	defer db.Close()

	instance := &Anything{}
	instance.Set(c)
	instance.Validate()
	if instance.Errors != nil {
		context := map[string]interface{}{"instance": instance}
		zapp.Render(c, `admin`, context, `new`)
		return
	}

	// 値の更新 ------------------------------------------>
	SaveAnything(db, instance)
	// <-------------------------------------------------

	// 完了ページへリダイレクト
	message := fmt.Sprintf("%v 追加しました", instance)
	zapp.SetFlashMessage(c, message)
	adminPrefix := `admin`
	controllerName, _ := zapp.ExtractControllerActionName(c.Request.URL.Path, adminPrefix)
	backURL := fmt.Sprintf("/%s/%s", adminPrefix, controllerName)
	c.Redirect(http.StatusFound, backURL)
}

// Update
func AdminAnythingUpdateHandler(c *gin.Context) {
	// DB接続を取得
	db, _ := NewGormDB(c)
	defer db.Close()

	a, err := getAnything(c)
	instance := &a
	if err != nil {
		zapp.RenderDirect(c, `admin/500`, map[string]interface{}{"message": err})
		return
	}

	instance.Set(c)
	instance.Validate()
	if instance.Errors != nil {
		context := map[string]interface{}{"instance": instance}
		zapp.Render(c, `admin`, context, `edit`)
		return
	}

	// 値の更新 ------------------------------------------>
	SaveAnything(db, instance)
	// <-------------------------------------------------

	// 完了ページへリダイレクト
	adminPrefix := `admin`
	controllerName, _ := zapp.ExtractControllerActionName(c.Request.URL.Path, adminPrefix)
	backURL := fmt.Sprintf("/%s/%s/show/%d", adminPrefix, controllerName, instance.ID)
	c.Redirect(http.StatusFound, backURL)
}

func AnythingFulltextListHandler(c *gin.Context) {
	ctx := appengine.NewContext(c.Request)
	db, _ := NewGormDB(c)
	defer db.Close()

	instance := Anything{}
	variables := map[string]interface{}{"instance": instance}

	// 入力文字列を空白区切りで分割する　山田　太郎　→　['山田' '太郎']
	q := c.Query(`q`)
	q = strings.Trim(q, ` `) // 前後の空白を除く
	// 半角・全角で区切る
	searchWords := []string{}
	for _, word := range strings.Split(q, ` `) { // 半角
		for _, w := range strings.Split(word, `　`) { // 全角
			if w != `` {
				searchWords = append(searchWords, w)
			}
		}
	}
	anyIDs := zapp.SearchByGAEFulltext(ctx, `user2`, searchWords)

	var instances []Anything
	db.Debug().Where(`id in (?)`, anyIDs).Find(&instances)

	variables[`q`] = q
	variables[`search_words`] = searchWords
	variables[`instances`] = instances
	c.Set(`variables`, variables)
}

func AnythingPutFulltextAllHandler(c *gin.Context) {
	ctx := appengine.NewContext(c.Request)
	db, _ := NewGormDB(c)
	defer db.Close()

	var instances []Anything
	db.Debug().Find(&instances)

	lower := strings.ToLower(reflect.TypeOf(Anything{}).Name())
	err := AnythingPutFulltexts(ctx, lower, instances)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
	c.String(200, "OK")
	c.Set("rendered", true)
}

func AnythingPutFulltexts(ctx context.Context, idx string, anys []Anything) error {
	searchAPIIndex, err := search.Open(idx)
	if err != nil {
		log.Println("failed to open index foo : %#v", err)
		return err
	}

	keys := []string{}
	values := []interface{}{}
	for _, any := range anys {
		r := reflect.ValueOf(any)
		method := r.MethodByName("Ngrams")
		if method.IsValid() {
			response := method.Call(nil)[0]
			ngrams := response.Interface().([]string)
			f := &zapp.Fulltext{
				Ngram:     strings.Join(ngrams, ` `),
				CreatedAt: any.CreatedAt,
				UpdatedAt: any.UpdatedAt,
			}
			keys = append(keys, fmt.Sprintf("%d", any.ID))
			values = append(values, f)
		}
	}

	if _, err := searchAPIIndex.PutMulti(ctx, keys, values); err != nil {
		log.Println("%v", err)
		return err
	}
	return nil
}

func AppendAnythingResources(group *gin.RouterGroup) {
	structName := zapp.GetType(Anything{})
	controllerName := strcase.ToSnake(structName)
	msg := fmt.Sprintf("Append %s as %s", structName, controllerName)
	log.Println(msg)
	group.GET(fmt.Sprintf("/%s/", controllerName), AdminAnythingListHandler)
	group.GET(fmt.Sprintf("/%s/search", controllerName), AnythingFulltextListHandler)
	group.GET(fmt.Sprintf("/%s/new", controllerName), AdminAnythingNewHandler)
	group.GET(fmt.Sprintf("/%s/edit/:id", controllerName), AdminAnythingEditHandler)
	group.GET(fmt.Sprintf("/%s/show/:id", controllerName), AdminAnythingShowHandler)
	group.POST(fmt.Sprintf("/%s/create", controllerName), AdminAnythingCreateHandler)
	group.POST(fmt.Sprintf("/%s/update", controllerName), AdminAnythingUpdateHandler)
}
