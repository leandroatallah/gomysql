package models

import (
	"database/sql"
	"time"

	"github.com/leandroatallah/gomysql/pkg/config"
)

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

var db *sql.DB

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetAllUsers() ([]User, error) {
	query := `SELECT id, username, password, created_at FROM users`
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Username, &u.Password, &u.CreatedAt)
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

func GetUserById(userId int) (User, error) {
	var u User
	query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	err := db.QueryRow(query, userId).Scan(&u.Id, &u.Username, &u.Password, &u.CreatedAt)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func GetUserByUsername(username string) (User, error) {
	var u User
	query := `SELECT id, username, password, created_at FROM users WHERE username = ?`
	err := db.QueryRow(query, username).Scan(&u.Id, &u.Username, &u.Password, &u.CreatedAt)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func CreateUser(u User) (int64, error) {
	rows, err := db.Query(`SELECT id FROM users WHERE username = ?`, u.Username)
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
	result, err := db.Exec(query, u.Username, u.Password, createdAt)
	if err != nil {
		return 0, err
	}
	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func DeleteUser(userID int) (User, error) {
	found, err := GetUserById(userID)
	if err != nil {
		return User{}, err
	}
	query := `DELETE FROM users WHERE id = ?`
	_, err = db.Exec(query, userID)
	if err != nil {
		return User{}, err
	}
	return found, nil
}
