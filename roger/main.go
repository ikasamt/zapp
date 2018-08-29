package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Roger"
	app.Usage = "This app for zapp"
	app.Version = "0.1"
	app.Commands = []cli.Command{
		{
			Name:    "generate_model",
			Aliases: []string{"gm"},
			Usage:   "generate db-model structs from database desc. This command needs `dsn` string  like app:@tcp(127.0.0.1:3306)/myapp ",
			Action:  generateModel,
		},
		{
			Name:    "migration",
			Aliases: []string{"mi"},
			Usage:   "migrate each sql queries to database. This command needs `dsn` string  like app:@tcp(127.0.0.1:3306)/myapp ",
			Action:  migration,
		},
	}
	app.Run(os.Args)
}
