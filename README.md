## zapp 

zapp is a library with gin-gonic to make admin scaffold.



## roger

roger is a command line tool. It has two features, generate-model and migration.

- generate-model : dumps go structs from database.

- migration : exec sql statement files in current directory.

### usage

```
example/sql $ roger mi app:@tcp(127.0.0.1:3306)/myapp 
2018/08/24 12:44:26 root:@tcp(127.0.0.1:3306)/zappandroger?parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4
2018/08/24 12:44:26 Migrate: 201808240725_init.sql
2018/08/24 12:44:26 Migrate: 201808240728_organization.sql
2018/08/24 12:44:26 Migrate: 201808240822_zessions.sql
2018/08/24 12:44:26 Migrate: 201808240824_alter_users.sql
```


```
$ roger g app:@tcp(127.0.0.1:3306)/myapp 

....

// roger_migrated
type RogerMigrated struct {
	ID         int
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	beforeJSON gin.H
	errors     error
}

....

```
