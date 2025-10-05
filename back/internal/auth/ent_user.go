package auth

import "github.com/google/uuid"

type User struct {
	UUID  string
	Name  string
	Email string
}

func NewUser(name string, email string) *User {
	return &User{
		UUID:  uuid.NewString(),
		Name:  name,
		Email: email,
	}
}
