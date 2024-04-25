package auth

import (
	"fmt"

	"github.com/color-predection/server/storage"
)

func UserLogin(email string, password string) (user storage.UserDetail, err error) {
	if !isValidEmail(email) {
		fmt.Printf("Invaild mail %s", email)
		return storage.UserDetail{}, err

	}
	if isExist(email) {

		for i := 0; i < len(storage.Users); i++ {

			if storage.Users[i].IsLogin && storage.Users[i].Email == email {
				fmt.Println("User already logged in:", email)
				return storage.Users[i], err

			}

			if email == storage.Users[i].Email && password == storage.Users[i].Password {
				storage.Users[i].IsLogin = true
				fmt.Println("Logging in with email:", email)
				return storage.Users[i], err

			}
		}

	}

	return storage.UserDetail{}, err // Return an empty User if login fails
}
