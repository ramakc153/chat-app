package database

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string
	Username string
	Password string
}

type UserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func InsertUser(username, password string) error {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(
		"INSERT INTO users (id, username, password, created_at) VALUES (gen_random_uuid(), $1, $2, NOW())",
		username, string(hashed_password),
	)
	if err != nil {
		return fmt.Errorf("users duplicate")
	}
	return nil
}

func GetUser(username, password string) (*User, error) {
	var QueriedUser User
	if err := DB.QueryRow(
		"SELECT id, username, password FROM users WHERE username=$1", username,
	).Scan(&QueriedUser.Id, &QueriedUser.Username, &QueriedUser.Password); err != nil {
		return nil, fmt.Errorf("User not found")
	}
	err := bcrypt.CompareHashAndPassword([]byte(QueriedUser.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("wrong password")
	}

	return &QueriedUser, nil
}

func GetAllUsers(username string) ([]UserResponse, error) {
	rows, err := DB.Query("SELECT id, username FROM users WHERE username != $1", username)
	if err != nil {
		return nil, fmt.Errorf("error when querying all users: %s", err)
	}
	var users []UserResponse

	for rows.Next() {
		var user UserResponse

		err = rows.Scan(&user.Id, &user.Username)
		if err != nil {
			return nil, fmt.Errorf("error when scanning getallusers: %s", err)
		}
		users = append(users, user)
	}
	return users, nil
}
