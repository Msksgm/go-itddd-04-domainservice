package user

import "fmt"

type UserName struct {
	name string
}

func NewUserName(name string) (*UserName, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required.")
	}
	return &UserName{name: name}, nil
}
