package storage

type UserDetail struct {
	Name     string
	Password string
	Email    string
	Age      int
	IsLogin  bool
}

var Users []UserDetail

// s = sign up
// l = login
// ls = users list
