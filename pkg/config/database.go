package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() {
	// TODO: Replace with environment variables
	var err error
	db, err = sql.Open("mysql", "root:123456@(localhost:3306)/gomysql?parseTime=true")
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

// TODO: Is it needed?
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
