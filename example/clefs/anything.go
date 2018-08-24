package clefs

import "github.com/gin-gonic/gin"

// Anything is defined as base Struct for genny
type Anything struct { //generic.Type
	ID         int64 //generic.Type
	UserID     int64 //generic.Type
	errors     error //generic.Type
	beforeJSON gin.H //generic.Type
} //generic.Type
