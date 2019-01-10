package zapp

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetParams は GET, POST を意識することなくパラメータを取得できる
func GetParams(c *gin.Context, key string) string {

	// URL内のパラメータを優先
	if c.Param(key) != `` {
		return c.Param(key)
	}

	// Post
	if c.Request.Method == `POST` {
		return c.PostForm(key)
	}
	// Get
	return c.Query(key)
}

func GetID(c *gin.Context) (int, error) {
	IDStr := GetParams(c, `id`)
	ID, _ := strconv.Atoi(IDStr)
	if ID == 0 {
		return 0, errors.New("Error: NOT FOUND")
	}
	return ID, nil
}

func ParseCheckbox(c *gin.Context, key string) bool {
	if GetParams(c, key) == `on` {
		return true
	}
	return false
}

func GetInt(c *gin.Context, key string) (int, error) {
	iStr := GetParams(c, key)
	i, _ := strconv.Atoi(iStr)
	if i == 0 {
		msg := fmt.Sprintf("Error: NOT FOUND %s", key)
		return 0, errors.New(msg)
	}
	return i, nil
}
