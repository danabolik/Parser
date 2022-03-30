package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var pool *sql.DB // Database connection pool.

func init() {
	prepareDb()
}

func main() {
	defer func(pool *sql.DB) {
		err := pool.Close()
		if err != nil {
			panic(err)
		}
	}(pool)
}

func prepareDb() {
	var err error
	pool, err = sql.Open("mysql", "groot:parser@tcp(mysql:3306)/parser")
	if err != nil {
		panic(err)
	}
	pool.SetConnMaxLifetime(time.Minute * 3)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)
}
