package auth

import (
	"fmt"

	"github.com/color-predection/server/storage"
)

func UserSignUp(user storage.UserDetail) storage.UserDetail {
	if !isValidEmail(user.Email) {
		fmt.Printf("Invaild mail %s", user.Email)
		return storage.UserDetail{}

	}
	if isExist(user.Email) {
		fmt.Print("User already exist")
		return storage.UserDetail{}

	}

	fmt.Println("Signing up user:", user.Name)
	storage.Users = append(storage.Users, user)
	return user
}
