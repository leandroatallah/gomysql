package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	sourceURL    = "file://migrations"
	dbIdentifier = "mysql"
)

func Connect() *sql.DB {
	dbConfig := NewDBConfig()

	var (
		db  *sql.DB
		err error
	)
	db, err = sql.Open("mysql",
		fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true&multiStatements=true",
			dbConfig.user, dbConfig.pass, dbConfig.host, dbConfig.port, dbConfig.name,
		),
	)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		sourceURL,
		dbIdentifier,
		driver,
	)
	if err != nil {
		panic(err)
	}
	m.Up()

	return db
}
