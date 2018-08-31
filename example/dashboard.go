package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ikasamt/zapp/zapp"
)

func adminDashboardHandler(c *gin.Context) {
	db := GetMasterDBInstance()
	defer db.Close()

	counts := map[string]interface{}{}

	objs := []interface{}{
		User{},
		Organization{},
		Report{},
		Phrase{},
	}
	for _, obj := range objs {
		var count int
		db.Model(obj).Count(&count)

		key := zapp.GetType(obj)
		counts[key] = count
	}

	context := map[string]interface{}{}
	context[`counts`] = counts
	zapp.Render(c, `admin`, context)
}
