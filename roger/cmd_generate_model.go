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
	dsn := c.Args()[0]

	var tablesCache []MySQLTable

	ret := []string{}
	for _, tableName := range showTables(dsn) {
		table := descTable(dsn, tableName)
		tablesCache = append(tablesCache, table)

		// struct
		ret = append(ret, ``)
		ret = append(ret, `// `+table.Name)
		lines := []string{}
		lines = append(lines, fmt.Sprintf(`type %s struct {`, table.StructName()))

		firstColumn := ``
		for _, column := range table.Columns {
			if firstColumn == `` {
				firstColumn = column.Name()
			}
			lines = append(lines, fmt.Sprintf(`%s %s`, column.Name(), column.GoType()))
		}
		lines = append(lines, `	beforeJSON     gin.H`)
		lines = append(lines, `	errors         error`)
		lines = append(lines, `}`)
		lines = append(lines, ``)
		lines = append(lines, ``)
		// define Table Name
		lines = append(lines, `func (`+table.StructName()+`) TableName() string { return "`+table.Name+`" }`)

		// String method
		lines = append(lines, `func (x *`+table.StructName()+`) String() string { return fmt.Sprintf("[%d]", x.`+firstColumn+`) }`)

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

	ret2 := []string{}
	for _, table := range tablesCache {
		controllerName := table.Name

		t := `// {{ link_to "%s" }}`
		t2 := fmt.Sprintf(t, controllerName)
		ret2 = append(ret2, t2)
	}

	tmp2 := strings.Join(ret2, CRLN)
	fmt.Println(tmp2)

	return nil
}
