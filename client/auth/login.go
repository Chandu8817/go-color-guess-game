package auth

import (
	"fmt"

	"github.com/color-predection/client/storage"
)

func UserLogin(email string, password string) storage.UserDetail {
	if !isValidEmail(email) {
		fmt.Printf("Invaild mail %s", email)
		return storage.UserDetail{}

	}
	if isExist(email) {

		for i := 0; i < len(storage.Users); i++ {

			if storage.Users[i].IsLogin {
				fmt.Println("User already logged in:", email)
				return storage.Users[i]

			}

			if email == storage.Users[i].Email && password == storage.Users[i].Password {
				storage.Users[i].IsLogin = true
				fmt.Println("Logging in with email:", email)
				return storage.Users[i]

			}
		}

	}

	return storage.UserDetail{} // Return an empty User if login fails
}
