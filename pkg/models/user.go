package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/leandroatallah/gomysql/pkg/config"
	"github.com/leandroatallah/gomysql/pkg/utils"
)

var db *sql.DB

func init() {
	config.Connect()
	db = config.GetDB()
}

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) isValid() error {
	if u.Id <= 0 {
		return errors.New("Invalid ID")
	}
	if len(u.Username) < 8 {
		return errors.New("Username length must be greater than or equal to 8")
	}
	if len(u.Password) < 6 {
		return errors.New("Password length must be greater than or equal to 6")
	}
	return nil
}

type UserResponse struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func UserToUserResponse(user User) UserResponse {
	var userRes UserResponse
	userRes.Username = user.Username
	userRes.CreatedAt = user.CreatedAt
	return userRes
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

func CreateUser(u User) (User, error) {
	if err := u.isValid(); err != nil {
		return User{}, err
	}
	rows, err := db.Query(`SELECT id FROM users WHERE username = ?`, u.Username)
	var exists bool
	for rows.Next() {
		exists = true
	}
	if err != nil {
		return User{}, err
	}
	if exists {
		return User{}, nil
	}

	createdAt := time.Now()
	passwordHash, err := utils.HashPassword(u.Password)
	if err != nil {
		return User{}, err
	}
	query := "INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)"
	result, err := db.Exec(query, u.Username, passwordHash, createdAt)
	if err != nil {
		return User{}, err
	}
	userID, err := result.LastInsertId()
	if err != nil {
		return User{}, err
	}
	u.Id = int(userID)
	u.CreatedAt = createdAt
	return u, nil
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

func UpdateUser(u User) (User, error) {
	if err := u.isValid(); err != nil {
		return User{}, err
	}
	query := `UPDATE users SET uername = ?, password = ? WHERE id = ?`
	_, err := db.Exec(query, u.Username, u.Password, u.Id)
	if err != nil {
		return User{}, err
	}
	return u, nil
}
