package main

import (
	"os"

	"github.com/ikasamt/zapp/zapp"
	"github.com/urfave/cli"
)

var ZappEnv *string
var ZappEnvironment zapp.Environment

func main() {
	// YAML 設定ファイル読み込み
	ZappEnv := os.Getenv("ZAPP_ENV")
	if ZappEnv == `` {
		ZappEnv = "development"
	}
	ZappEnvironments := zapp.ReadEnvironments()
	ZappEnvironment = ZappEnvironments[ZappEnv]

	app := cli.NewApp()
	app.Name = "Roger"
	app.Usage = "This app for zapp"
	app.Version = "0.1.2"
	app.Commands = []cli.Command{
		{
			Name:    "generateModel",
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
