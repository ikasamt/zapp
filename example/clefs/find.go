package clefs

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ikasamt/zapp/zapp"
)

//go:generate genny -pkg=main -in=$GOFILE -out=../zzz-autogen-$GOFILE gen "Anything=User,Organization,Report,Phrase"

// select All
func selectAllAnythings(selects string) (instances []*Anything) {
	db := GetMasterDBInstance()
	defer db.Close()
	db.Debug().Select(selects).Find(&instances)
	return
}

// fetch One
func fetchAnything(anyID int) (any Anything) {
	db := GetMasterDBInstance()
	defer db.Close()
	db.Debug().Where("id = ?", anyID).First(&any)
	any.beforeJSON = zapp.CallMethod(any, `AsJSON`, gin.H{}).(gin.H)
	return any
}

// get One
func getAnything(c *gin.Context) (instance Anything, e error) {
	// 対象IDを取得
	ID, err := zapp.GetID(c)
	if err != nil {
		log.Println(err)
		return instance, err
	}
	// DBから取得
	instance = fetchAnything(ID)
	return instance, nil
}
