package main

import (
	"Parser/internal/parser"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
	"time"
)

var pool *sql.DB // Database connection pool.

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Can not load environment vars")
	}
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
	driver := os.Getenv("DRIVER")
	dsName := os.Getenv("DATA_SOURCE_NAME")

	pool, err = sql.Open(driver, dsName)
	if err != nil {
		panic(err)
	}
	pool.SetConnMaxLifetime(time.Minute * 3)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	parser.Parse()
}
