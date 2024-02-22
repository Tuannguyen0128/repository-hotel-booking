package util

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"repository-hotel-booking/internal/app/model"
)

func InitConnection(config model.DBConfig) *sql.DB {
	cfg := mysql.Config{
		User:                 config.Username,
		Passwd:               config.Password,
		Net:                  "tcp",
		Addr:                 config.Host,
		DBName:               config.Schema,
		AllowNativePasswords: true,
	}
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	return db
}
