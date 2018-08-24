package zapp

import (
	"fmt"
	"html/template"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Pager struct {
	page        int
	perPage     int
	totalCount  int
	queryString template.URL
	baseURL     string
}

func NewPager(c *gin.Context, totalCount interface{}) Pager {
	page, _ := strconv.Atoi(c.DefaultQuery(`page`, `0`))
	perPage, _ := strconv.Atoi(c.DefaultQuery(`per_page`, fmt.Sprintf(`%d`, DefaultPerPage)))

	pager := Pager{}
	pager.page = page
	pager.perPage = perPage
	pager.baseURL = c.Request.URL.Path

	if totalCount != nil {
		pager.totalCount = totalCount.(int)

	}

	// 現状のクエリパラメータから特殊な数個を除いた文字列を作成する（ページネーション用途）
	queries := c.Request.URL.Query()
	delete(queries, `page`)
	delete(queries, `per_page`)
	if queries != nil {
		keys := make([]string, 0, len(queries))
		for k, vs := range queries {
			var kvs []string
			for _, v := range vs {
				kvs = append(kvs, k+`=`+v)
			}
			keys = append(keys, strings.Join(kvs, `&`))
		}
		sort.Strings(keys)
		pager.queryString = template.URL(strings.Join(keys, `&`))
	}
	return pager
}

func (p *Pager) Pages() map[int]int {

	var totalCount = UnknownTotalCount
	if p.totalCount != 0 {
		totalCount = p.totalCount
	}
	labels := make(map[int]int)
	for i := 0; i < (totalCount/p.perPage)+1; i++ {
		labels[i] = i + 1 // labels[0] = 1, labels[1] = 2 ...
	}
	return labels
}
