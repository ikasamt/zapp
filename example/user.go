package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ikasamt/zapp/zapp"
	"github.com/jinzhu/gorm"
)

// Setter
func (x *User) Setter(c *gin.Context) {
	x.OrganizationID, _ = strconv.Atoi(zapp.GetParams(c, "organization_id"))
	x.Name = zapp.GetParams(c, "name")
	x.Email = zapp.GetParams(c, "email")
	if zapp.GetParams(c, "password") != `` {
		salt := ZappEnvironment[`password_salt`].(string)
		x.HashedPassword = zapp.HashPassword(salt, zapp.GetParams(c, "password"))
		x.IsInitialPassword = !(zapp.GetParams(c, "is_initial_password") == ``)
	}
}

// Search
func (x *User) Search(q *gorm.DB) *gorm.DB {
	if x.Email != `` {
		q = q.Where("email LIKE ? ", "%"+x.Email+"%")
	}
	if x.Name != `` {
		q = q.Where("name LIKE ? ", "%"+x.Name+"%")
	}
	if x.OrganizationID != 0 {
		q = q.Where("organization_id = ?", x.OrganizationID)
	}
	return q
}

// String
func (x *Organization) String() string {
	return fmt.Sprintf(`[%20d] aa%saa`, x.ID, x.Name)
}
