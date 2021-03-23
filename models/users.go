package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	users = map[string]User{
		"li": User{"li", "123456"},
	}
)

func FindUser(username string) *User {
	if user, ok := users[username]; ok {
		return &user
	}
	return nil
}

func CreateUser(username, password string) {
	users[username] = User{username, password}
}
