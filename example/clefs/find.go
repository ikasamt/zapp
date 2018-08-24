package clefs

import (
	"errors"

	"github.com/jinzhu/gorm"
)

//go:generate genny -pkg=main -in=$GOFILE -out=../zzz-autogen-$GOFILE gen "Anything=User"

// find One
func findAnything(db *gorm.DB, anyID int64) (any Anything, e error) {
	db.Debug().Where("id = ?", anyID).First(&any)
	if any.ID == 0 {
		return Anything{}, errors.New("Error: USER NOT FOUND")
	}
	return any, nil
}

// find All
func findAllAnythings(db *gorm.DB, selects string) (instances []Anything) {
	db.Debug().Select(selects).Find(&instances)
	return
}
