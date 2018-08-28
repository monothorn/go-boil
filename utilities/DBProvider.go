package utilities

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var once sync.Once

//GetInstance singleton pattern using sync
func GetInstance(dsn string) (*sql.DB, error) {
	once.Do(func() {
		fmt.Print("Connecting to .." + dsn)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err.Error())
		}
		Db = db
	})
	return Db, nil
}
