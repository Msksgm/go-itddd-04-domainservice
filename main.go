package main

import (
	"database/sql"
	"fmt"

	"github.com/Msksgm/itddd-go-04-domainservice/domain/model/user"
)

func main() {
}

func CreateUser(db *sql.DB, name string) (err error) {
	userName, err := user.NewUserName(name)
	if err != nil {
		return
	}
	newUser, err := user.NewUser(*userName)
	if err != nil {
		return
	}

	userService := user.NewUserService(db)
	isExists, err := userService.Exists(newUser)
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
	if err != nil {
		return err
	}
	if isExists {
		return fmt.Errorf("%s is already exists.", name)
	}

	_, err = db.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", newUser.UserId(), newUser.UserName())
	if err != nil {
		return
	}
	return nil
}
