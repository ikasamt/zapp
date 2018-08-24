package clefs

import (
	"reflect"
	"strings"
)

//go:generate genny -pkg=main -in=$GOFILE -out=../zzz-autogen-$GOFILE gen "Anything=Organization Something=User"

type Something struct { //generic.Type
	ID         int64 //generic.Type
	UserID     int64 //generic.Type
	AnythingID int64 //generic.Type
	errors     error //generic.Type
} //generic.Type

func (x Anything) Somethings() (instances []Something) {
	db := GetMasterDBInstance()
	defer db.Close()
	lower := strings.ToLower(reflect.TypeOf(x).Name())
	db.Debug().Where(lower+`_id = ?`, x.ID).Find(&instances)
	return
}
