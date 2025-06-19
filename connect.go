package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id        int
	username  string
	password  string
	createdAt time.Time
}

func initUsersTable(db *sql.DB) (sql.Result, error) {
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

func createUser(db *sql.DB, u user) (int64, error) {
	rows, err := db.Query(`SELECT id FROM users WHERE username = ?`, u.username)
	var exists bool
	for rows.Next() {
		exists = true
	}
	if err != nil {
		return 0, err
	}
	if exists {
		// TODO: Improve this code
		return 0, nil
	}

	createdAt := time.Now()
	query := "INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)"
	result, err := db.Exec(query, u.username, u.password, createdAt)
	if err != nil {
		return 0, err
	}
	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Println("User ID:", userID)
	return userID, nil
}

func listUsers(db *sql.DB) ([]user, error) {
	query := `SELECT id, username, password, created_at FROM users`
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var users []user
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func getUserById(db *sql.DB, userId int) (user, error) {
	var u user
	query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	err := db.QueryRow(query, userId).Scan(&u.id, &u.username, &u.password, &u.createdAt)
	if err != nil {
		return user{}, err
	}
	return u, nil
}

func getUserByUsername(db *sql.DB, username string) (user, error) {
	var u user
	query := `SELECT id, username, password, created_at FROM users WHERE username = ?`
	err := db.QueryRow(query, username).Scan(&u.id, &u.username, &u.password, &u.createdAt)
	if err != nil {
		return user{}, err
	}
	return u, nil
}

func deleteUser(db *sql.DB, userID int) (user, error) {
	found, err := getUserById(db, userID)
	if err != nil {
		return user{}, err
	}
	query := `DELETE FROM users WHERE id = ?`
	_, err = db.Exec(query, userID)
	if err != nil {
		return user{}, err
	}
	return found, nil
}

func printListUsers(db *sql.DB) []user {
	list, err := listUsers(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("List users:")
	for _, u := range list {
		fmt.Println(u)
	}
	return list
}

func main() {
	// TODO: Replace with environment variables
	db, err := sql.Open("mysql", "root:123456@(localhost:3306)/gomysql?parseTime=true")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// _, err = initUsersTable(db)
	// if err != nil {
	// 	panic(err)
	// }

	// List users
	printListUsers(db)
	fmt.Println()

	// Create user
	var userData user
	userData.username = "Mark Callaway"
	// userData.username = "John Doe"
	userData.password = "123456"
	createUser(db, userData)

	// List users (after create)
	list := printListUsers(db)
	fmt.Println()

	// Get one user
	u, err := getUserById(db, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Get user:")
	fmt.Println(u)
	fmt.Println()

	// Delete user
	deleteUser(db, list[len(list)-1].id)

	// List users (after delete)
	printListUsers(db)
}
