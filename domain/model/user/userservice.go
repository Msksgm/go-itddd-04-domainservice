package user

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) (*UserService, error) {
	return &UserService{db: db}, nil
}

func (userService *UserService) Exists(user *User) (isExists bool, err error) {
	tx, err := userService.db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	rows, err := tx.Query("SELECT * FROM users WHERE username = $1", user.UserName())
	if err != nil {
		return false, fmt.Errorf("userservice.Exists(): %v", err)
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}
	return false, nil
}
