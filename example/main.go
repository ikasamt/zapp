package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/ikasamt/zapp/zapp"
	"github.com/jinzhu/gorm"
)

func GetMasterDBInstance() (db *gorm.DB) {
	dsn := ZappEnvironment[`mysql`].(string)
	return zapp.GetDBInstance(dsn)
}

var ZappEnv *string
var ZappEnvironment zapp.Environment

func main() {

	// YAML 設定ファイル読み込み
	ZappEnv := os.Getenv("ZAPP_ENV")
	if ZappEnv == `` {
		ZappEnv = "development"
	}
	ZappEnvironments := zapp.ReadEnvironments()
	ZappEnvironment = ZappEnvironments[ZappEnv]

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// gin初期化
	r := gin.Default()
	// セッションを利用する設定
	store := cookie.NewStore([]byte("this_is_secret_salt_message"))
	r.Use(sessions.Sessions("mysession", store))

	// 静的ファイル
	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	r.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	r.GET("/sample", sampleHandler)

	// User
	r.GET("/admin/user/", adminUserListHandler)
	r.GET("/admin/user/new", adminUserNewHandler)
	r.GET("/admin/user/edit/:id", adminUserEditHandler)
	r.GET("/admin/user/show/:id", adminUserShowHandler)
	r.POST("/admin/user/create", adminUserCreateHandler)

	// Organization
	r.GET("/admin/organization/", adminOrganizationListHandler)
	r.GET("/admin/organization/new", adminOrganizationNewHandler)
	r.GET("/admin/organization/edit/:id", adminOrganizationEditHandler)
	r.GET("/admin/organization/show/:id", adminOrganizationShowHandler)

	r.Run(fmt.Sprintf(":%d", ZappEnvironment[`port`].(int)))

}
