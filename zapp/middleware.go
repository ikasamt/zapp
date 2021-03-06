package zapp

import (
	"github.com/gin-gonic/gin"
)

func ErrorMiddleware(layoutName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Writer.Status()
		switch code {
		case 401, 402, 403, 404:
			RenderDirect(c, `app/404`, gin.H{})
		case 500:
			RenderDirect(c, `app/500`, gin.H{})
		}
		c.Next()
	}
}

func IsDevMiddleware(isDev bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if isDev {
			c.Set(`__is_dev`, true)
		}
		c.Next()
	}
}
