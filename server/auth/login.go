package auth

import (
	"fmt"

	"github.com/color-predection/server/storage"
)

func UserLogin(email string, password string) (user storage.UserDetail, err error) {
	fmt.Println("xxxxxxxxx heres", email, password)
	if !isValidEmail(email) {
		fmt.Printf("Invaild mail %s", email)
		err = fmt.Errorf("invaild to login ")
		return storage.UserDetail{}, err

	}
	if isExist(email) {
		fmt.Println("xxxxxxxxx isExist check", email, password)

		for i := 0; i < len(storage.Users); i++ {

			if storage.Users[i].IsLogin && storage.Users[i].Email == email {
				fmt.Println("User already logged in:", email)
				err = fmt.Errorf("user already logged in")

				return storage.Users[i], err

			}

			if email == storage.Users[i].Email && password == storage.Users[i].Password {
				storage.Users[i].IsLogin = true
				fmt.Println("Logging in with email:", email)
				err = nil
				return storage.Users[i], err

			}
		}

	} else {
		println("user not found")
		err = fmt.Errorf("user not found")

		return storage.UserDetail{}, err
	}

	return storage.UserDetail{}, fmt.Errorf("failed login") // Return an empty User if login fails
}
