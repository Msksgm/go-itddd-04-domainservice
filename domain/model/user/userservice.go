package user

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type UserServicer interface {
	Exists(user *User) (bool, error)
}

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) UserServicer {
	return &UserService{db: db}
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

	rows, err := tx.Query("SELECT * FROM users WHERE name = $1", user.UserName())
	if err != nil {
		return
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}
	return false, nil
}
