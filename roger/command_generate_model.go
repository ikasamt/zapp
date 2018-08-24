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
	dsn := c.Args().First()
	log.Println(dsn)

	ret := []string{}
	for _, tableName := range showTables(dsn) {
		table := descTable(dsn, tableName)
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
