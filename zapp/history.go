package zapp

import (
	"time"

	"github.com/gin-gonic/gin"
)

// 履歴
type History struct {
	ID         int
	Model      string
	InstanceID int
	Data       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	beforeJSON gin.H
	errors     error
}
