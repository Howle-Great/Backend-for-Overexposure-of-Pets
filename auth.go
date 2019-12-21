package main

import (
	"golang.org/x/crypto/bcrypt"
)

// type User struct {
// 	Nick     string
// 	Password string
// }

type AuthManager struct{}

func (am *AuthManager) SignIn(Items *ItemRepository, user *User) (string, error) {
	hashedPassword, err := Items.GetUsersPassword(user.Nickname)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		return "", err
	}
	return hashedPassword, nil
}
