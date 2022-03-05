package user

import (
	"reflect"

	"github.com/google/uuid"
)

type User struct {
	userId   UserId
	userName UserName
}

func NewUser(userName UserName) (*User, error) {
	userId, err := NewUserId(uuid.New().String())
	if err != nil {
		return nil, err
	}
	return &User{userId: *userId, userName: userName}, nil
}

func (user *User) UserName() string {
	return user.userName.Name()
}

func (user *User) UserId() string {
	return user.userId.Id()
}

func (user *User) Equals(other *User) bool {
	return reflect.DeepEqual(user.userId, other.userId)
}
