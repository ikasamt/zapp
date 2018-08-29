package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ikasamt/zapp/zapp"
)

func sampleHandler(c *gin.Context) {
	db := GetMasterDBInstance()
	defer db.Close()

	var user User
	db.Debug().First(&user)

	var organization Organization
	db.Debug().First(&organization)

	context := map[string]interface{}{`user`: user, `organization`: organization}
	zapp.RenderJade(c, `default`, `top`, `list`, context)
}
