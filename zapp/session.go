package zapp

import (
	"log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// セッションから取得する
func GetSession(c *gin.Context, key string, defaultValue interface{}) interface{} {
	session := sessions.Default(c)
	val := session.Get(key)
	if val == nil {
		return defaultValue
	}
	return val
}

// セッションを保存する
func SetSession(c *gin.Context, key string, value interface{}) {
	session := sessions.Default(c)
	session.Set(key, value)
	err:=session.Save()
	log.Println(err)
}

// Flashメッセージを取得する
func GetFlashMessage(c *gin.Context) string {
	session := sessions.Default(c)
	val := session.Get(`flash_message`)
	session.Set(`flash_message`, ``) // 読んだらすぐクリアーする
	session.Save()
	if val == nil {
		return ``
	}
	return val.(string)
}

// Flashメッセージを保存する
func SetFlashMessage(c *gin.Context, value string) {
	session := sessions.Default(c)
	session.Set(`flash_message`, value)
	session.Save()
}
