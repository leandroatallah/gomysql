package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() {
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", DBUser, DBPass, DBHost, DBPort, DBName))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}

func InitUsersTable(db *sql.DB) (sql.Result, error) {
	query := `
	CREATE TABLE users (
		id INT AUTO_INCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME,
		PRIMARY KEY (id)
	)
	`
	return db.Exec(query)
}
