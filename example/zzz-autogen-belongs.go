// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ikasamt/zapp/zapp"
)

func (x Organization) Users() (instances []User) {
	db := GetMasterDBInstance()
	defer db.Close()
	lower := strings.ToLower(reflect.TypeOf(x).Name())
	db.Debug().Where(lower+`_id = ?`, x.ID).Find(&instances)
	return
}

func (x User) OrganizationOptions() string {
	tmp := ""
	tmp += "<option></option>"
	for _, any := range selectAllOrganizations("id, name") {
		ID_ := reflect.ValueOf(x.OrganizationID)
		var ID int
		if ID_.Type().Kind() == reflect.Ptr {
			if !ID_.IsNil() {
				ID = int(ID_.Elem().Int())
			}
		} else {
			ID = int(ID_.Int())
		}

		label := zapp.CallMethod(any, `String`, fmt.Sprintf(`[%d]`, any.ID))

		if ID == any.ID {
			tmp += fmt.Sprintf("<option value='%d' selected>%s</option>", any.ID, label)
		} else {
			tmp += fmt.Sprintf("<option value='%d'>%s</option>", any.ID, label)
		}
	}
	return tmp
}
