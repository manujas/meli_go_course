package domain

// User struct
type User struct {
	Name string
	Nick string
	Pass string
	Mail string
}

// NewUser function
func NewUser(userName, userNick, userMail, userPass string) *User {
	return &User{userName, userNick, userPass, userMail}
}
