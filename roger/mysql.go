package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetSqlDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
	return db
}

func descTable(dsn string, tableName string) (table MySQLTable) {
	db := GetSqlDB(dsn)
	defer db.Close()

	table.Name = tableName

	q := fmt.Sprintf(`DESC %s`, tableName)
	rows, err := db.Query(q)
	if err != nil {
		log.Println(`Err`)
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var column MySQLColumn
		rows.Scan(
			&column.Field,
			&column.Type,
			&column.Null,
			&column.Key,
			&column.Default,
			&column.Extra,
		)
		table.Columns = append(table.Columns, column)
	}
	return
}

func showTables(dsn string) (tableNames []string) {
	db := GetSqlDB(dsn)
	defer db.Close()

	q := fmt.Sprintf(`SHOW TABLES`)
	rows, err := db.Query(q)
	if err != nil {
		log.Println(`Err`)
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		rows.Scan(&name)
		tableNames = append(tableNames, name)
	}
	return
}
