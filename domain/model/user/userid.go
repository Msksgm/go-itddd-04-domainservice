package user

import "reflect"

type UserId struct {
	id string
}

func NewUserId(uuid string) (*UserId, error) {
	userId := new(UserId)
	userId.id = uuid
	return userId, nil
}

func (userId *UserId) Id() string {
	return userId.id
}

func (userName *UserId) Equals(other *UserId) bool {
	return reflect.DeepEqual(userName.id, other.id)
}
