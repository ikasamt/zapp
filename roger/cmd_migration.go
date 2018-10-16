package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli"
)

const RogerTableName = `roger_migrated`

type Roger struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func rogerMigratedTableFound(dsn string) bool {
	for _, tableName := range showTables(dsn) {
		if tableName == RogerTableName {
			return true
		}
	}
	return false
}

func isMigrated(dsn string, name string) bool {
	db := GetSqlDB(dsn)
	defer db.Close()

	q := `SELECT * FROM %s`
	q = fmt.Sprintf(q, RogerTableName)

	rows, err := db.Query(q)
	if err != nil {
		log.Println(`Err`)
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var r Roger
		rows.Scan(
			&r.ID,
			&r.Name,
			&r.CreatedAt,
			&r.UpdatedAt,
		)

		if r.Name == name {
			return true
		}
	}
	return false
}

func createRogerMigrated(dsn string) {
	db := GetSqlDB(dsn)
	defer db.Close()

	sql := `
	CREATE TABLE %s (
		id int(11) NOT NULL AUTO_INCREMENT,
		name varchar(255) NOT NULL,
		created_at datetime DEFAULT NULL,
		updated_at datetime DEFAULT NULL,
		PRIMARY KEY (id)
	) ENGINE=InnoDB AUTO_INCREMENT=96 DEFAULT CHARSET=utf8mb4;
	`
	sql = fmt.Sprintf(sql, RogerTableName)
	db.Exec(sql)
	log.Println(`RogerTableName created`)
}

func migration(c *cli.Context) error {
	dsn := c.Args()[0]
	log.Println(dsn)

	if !rogerMigratedTableFound(dsn) {
		createRogerMigrated(dsn)
	}

	db := GetSqlDB(dsn)
	defer db.Close()

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, _ := ioutil.ReadDir(currentDir)
	for _, f := range files {
		ext := filepath.Ext(f.Name())
		if ext == `.sql` {
			if !isMigrated(dsn, f.Name()) {
				sqlBytes, err := ioutil.ReadFile(f.Name())
				if err != nil {
					log.Println(err)
				}
				sqlStr := string(sqlBytes)
				db.Exec(sqlStr)
				log.Println(fmt.Sprintf(`Migrate: %s`, f.Name()))

				insertSQL := `INSERT INTO %s (name, created_at, updated_at) VALUES ( "%s" , NOW(), NOW()) `
				insertSQL = fmt.Sprintf(insertSQL, RogerTableName, f.Name())
				db.Exec(insertSQL)
			}
		}
	}

	return nil
}
