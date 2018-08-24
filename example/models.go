package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

// organizations
type Organization struct {
	ID         int
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	beforeJSON gin.H
	errors     error
}

// roger_migrated
type RogerMigrated struct {
	ID         int
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	beforeJSON gin.H
	errors     error
}

// users
type User struct {
	ID                int
	OrganizationID    int
	Name              string
	Email             string
	HashedPassword    string
	PhotoKeyName      string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	IsInitialPassword string
	Department        string
	beforeJSON        gin.H
	errors            error
}

// zessions
type Zession struct {
	ID         int
	SessionID  string
	Data       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	beforeJSON gin.H
	errors     error
}
