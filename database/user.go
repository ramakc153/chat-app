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

func GetUser(username, password string) error {
	var QueriedUser User
	if err := DB.QueryRow(
		"SELECT id, username, password FROM users WHERE username=$1", username,
	).Scan(&QueriedUser.Id, &QueriedUser.Username, &QueriedUser.Password); err != nil {
		return fmt.Errorf("User not found")
	}
	err := bcrypt.CompareHashAndPassword([]byte(QueriedUser.Password), []byte(password))
	if err != nil {
		return fmt.Errorf("wrong password")
	}

	return nil
}
