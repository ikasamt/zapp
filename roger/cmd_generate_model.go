package main

import (
	"fmt"
	"go/format"
	"log"
	"strings"

	"github.com/urfave/cli"
)

const CRLN = "\r\n"

func generateModel(c *cli.Context) error {
	dsn := ZappEnvironment[`mysql`].(string)

	ret := []string{}
	for _, tableName := range showTables(dsn) {
		table := descTable(dsn, tableName)
		// struct
		ret = append(ret, ``)
		ret = append(ret, `// `+table.Name)
		lines := []string{}
		lines = append(lines, fmt.Sprintf(`type %s struct {`, table.StructName()))
		for _, column := range table.Columns {
			lines = append(lines, fmt.Sprintf(`%s %s`, column.Name(), column.GoType()))
		}
		lines = append(lines, `	beforeJSON     gin.H`)
		lines = append(lines, `	errors         error`)
		lines = append(lines, `}`)
		lines = append(lines, ``)
		lines = append(lines, ``)
		// String method
		lines = append(lines, `func (x *`+table.StructName()+`) String() string { return fmt.Sprintf("[%d]", x.ID) }`)

		// Search method

		tmpStr := strings.Join(lines, CRLN)
		tmpByte, err := format.Source([]byte(tmpStr))
		if err != nil {
			log.Println(err)
		}
		goStructStr := string(tmpByte)
		ret = append(ret, goStructStr)
		ret = append(ret, ``)
		ret = append(ret, ``)
	}
	tmp := strings.Join(ret, CRLN)
	fmt.Println(tmp)
	return nil
}
