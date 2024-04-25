package auth

import (
	"regexp"

	"github.com/color-predection/client/storage"
)

func isValidEmail(email string) bool {
	// Regular expression for basic email validation
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}

func isExist(email string) bool {
	for i := 0; i < len(storage.Users); i++ {
		if email == storage.Users[i].Email {

			return true
		}

	}
	return false
}

func GetUser(email string) storage.UserDetail {
	for i := 0; i < len(storage.Users); i++ {
		if email == storage.Users[i].Email {

			return storage.Users[i]
		}

	}
	return storage.UserDetail{}

}

func UserList() []storage.UserDetail {
	return storage.Users
}
