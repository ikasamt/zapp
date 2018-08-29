package zapp

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"html/template"

	"github.com/gin-gonic/gin"
)

type Pager struct {
	Name       string
	Page       int
	PerPage    int
	TotalCount int
	BaseURL    string
	Links      []PagerLink
}

type PagerLink struct {
	Number int
	Label  string
	Href   template.HTML
}

func NewPager(c *gin.Context, totalCount interface{}) Pager {
	page, _ := strconv.Atoi(c.DefaultQuery(`page`, `0`))
	perPage, _ := strconv.Atoi(c.DefaultQuery(`per_page`, fmt.Sprintf(`%d`, DefaultPerPage)))

	pager := Pager{}
	pager.Name = `aaa`
	pager.Page = page
	pager.PerPage = perPage
	pager.BaseURL = c.Request.URL.Path

	if totalCount != nil {
		pager.TotalCount = totalCount.(int)
	}

	var queryString string
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
		queryString = strings.Join(keys, `&`)

	}

	var _totalCount = UnknownTotalCount
	if pager.TotalCount != 0 {
		_totalCount = pager.TotalCount
	}

	// for i := 0; i < (_totalCount/pager.PerPage)+1; i++ {

	lastPageNumber := int(_totalCount/pager.PerPage) + 1

	// prev
	pagerLink := PagerLink{}
	pagerLink.Number = 0
	if pager.Page > 0 {
		pagerLink.Number = pager.Page - 1
	}
	pagerLink.Label = "<<"
	pagerLink.Href = template.HTML(fmt.Sprintf("%s/?%s&page=%d&per_page=%d", c.Request.URL.Path, queryString, pagerLink.Number, pager.PerPage))
	pager.Links = append(pager.Links, pagerLink)

	// numbers
	numbers := []int{}
	for _, i := range []int{-9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		number := pager.Page + i
		if number <= lastPageNumber {
			if number > 0 {
				numbers = append(numbers, number)
			}
		}
	}

	for i := range numbers {
		pagerLink := PagerLink{}
		pagerLink.Number = i
		pagerLink.Label = fmt.Sprintf("%d", i+1)
		pagerLink.Href = template.HTML(fmt.Sprintf("%s/?%s&page=%d&per_page=%d", c.Request.URL.Path, queryString, pagerLink.Number, pager.PerPage))
		pager.Links = append(pager.Links, pagerLink)
	}

	// next
	pagerLink = PagerLink{}
	pagerLink.Number = lastPageNumber - 1
	if pager.Page < lastPageNumber-1 {
		pagerLink.Number = pager.Page + 1
	}
	pagerLink.Label = ">>"
	pagerLink.Href = template.HTML(fmt.Sprintf("%s/?%s&page=%d&per_page=%d", c.Request.URL.Path, queryString, pagerLink.Number, pager.PerPage))
	pager.Links = append(pager.Links, pagerLink)

	return pager
}
