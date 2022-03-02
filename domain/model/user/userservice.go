package user

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (UserService *UserService) Exists(db *sql.DB, user *User) (isExists bool, err error) {
	tx, err := db.Begin()
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
