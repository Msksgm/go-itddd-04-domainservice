package user

type User struct {
	userId   UserId
	userName UserName
}

func NewUser(userId UserId, userName UserName) (*User, error) {
	return &User{userId: userId, userName: userName}, nil
}
