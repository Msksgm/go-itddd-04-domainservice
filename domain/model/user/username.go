package user

type UserName struct {
	name string
}

func NewUserName(name string) (*UserName, error) {
	return &UserName{name: name}, nil
}
