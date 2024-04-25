package auth

import (
	"fmt"

	"github.com/color-predection/client/storage"
)

func UserSignUp(name string, password string, email string, age int) storage.UserDetail {
	if !isValidEmail(email) {
		fmt.Printf("Invaild mail %s", email)
		return storage.UserDetail{}

	}
	if isExist(email) {
		fmt.Print("User already exist")
		return storage.UserDetail{}

	}

	fmt.Println("Signing up user:", name)
	user := storage.UserDetail{Name: name, Password: password, Email: email, Age: age, IsLogin: false}
	storage.Users = append(storage.Users, user)
	return user
}
