package main

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/ikasamt/zapp/zapp"
)

func dbFindHandler(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/teame_feedback?parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4"
	db := zapp.GetDBInstance(dsn)
	defer db.Close()

	var user User
	db.Debug().First(&user)

	var comment Comment
	db.Debug().First(&comment)

	context := map[string]interface{}{`user`: user, `comment`: comment}
	zapp.RenderJade(c, `layout`, `index`, context)
}

func main() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// gin初期化
	r := gin.Default()
	// セッションを利用する設定
	store := cookie.NewStore([]byte("this_is_secret_salt_message"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/dbfind", dbFindHandler)
	r.Run(":3001")

}
