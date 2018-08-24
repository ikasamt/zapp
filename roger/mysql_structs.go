package main

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

type MySQLProcess struct {
	ID      int
	User    string
	Host    string
	Db      string
	Command string
	Time    int
	State   string
	Info    string
}

type MySQLExplain struct {
	ID           int
	SelectType   string
	Table        string
	Type         string
	PossibleKeys string
	Key          string
	KeyLen       string
	Ref          string
	Rows         int
	Extra        string
}

type MySQLColumn struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

func (mc *MySQLColumn) Name() string {
	// id -> ID
	if mc.Field == `id` {
		return `ID`
	}
	// foo_id -> FooID
	if strings.HasSuffix(mc.Field, `_id`) {
		tmp := mc.Field
		tmp = strings.TrimRight(tmp, `_id`)
		tmp = tmp + `_ID`
		return strcase.ToCamel(tmp)
	}

	// else
	return strcase.ToCamel(mc.Field)
}

func (mc *MySQLColumn) IsPrimary() bool {
	if mc.Key == `PRI` {
		return true
	}
	return false
}

func (mc *MySQLColumn) IsNullable() bool {
	if mc.Null == `Yes` {
		return true
	}
	return false
}

func (mc *MySQLColumn) HasDefault() bool {
	if mc.Default != `NULL` {
		return true
	}
	return false
}

func (mc *MySQLColumn) HasAutoIncrement() bool {
	if mc.Default == `auto_increment` {
		return true
	}
	return false
}

func (mc *MySQLColumn) GoType() string {
	switch mc.Type {
	case `tinyint(11)`:
		return `bool`
	case `int(11)`:
		return `int`
	case `bigint(20)`:
		return `int64`
	case `datetime`:
		return `time.Time`
	default:
		return `string`
	}
}

type MySQLTable struct {
	Name    string
	Columns []MySQLColumn
}

func (mt *MySQLTable) StructName() string {
	tmp := inflection.Singular(mt.Name)
	return strcase.ToCamel(tmp)
}
