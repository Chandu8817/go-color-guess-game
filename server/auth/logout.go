package auth

import "github.com/color-predection/server/storage"

func LogOut(email string) string {

	for i := 0; i < len(storage.Users); i++ {

		if email == storage.Users[i].Email {

			if !storage.Users[i].IsLogin {

				return "User not logged in"

			} else {
				storage.Users[i].IsLogin = false
			}

		}

	}

	return "logout success"

}
