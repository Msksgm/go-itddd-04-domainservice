package user

import "fmt"

type UserName struct {
	name string
}

func NewUserName(name string) (*UserName, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required.")
	}
	if len(name) < 3 {
		return nil, fmt.Errorf("name must not be less than 3 characters.")
	}
	return &UserName{name: name}, nil
}

func (userName *UserName) Name() string {
	return userName.name
}
