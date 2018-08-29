package zapp

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetParams は GET, POST を意識することなくパラメータを取得できる
func GetParams(c *gin.Context, key string) string {

	// URL内のパラメータを優先
	if c.Param("id") != `` {
		return c.Param("id")
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
