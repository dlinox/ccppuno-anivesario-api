package database

import (
	"database/sql"
	"fmt"
	"log"
	"go-test/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	cfg := config.LoadConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBName)
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}