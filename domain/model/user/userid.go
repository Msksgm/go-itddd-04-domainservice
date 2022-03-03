package user

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
