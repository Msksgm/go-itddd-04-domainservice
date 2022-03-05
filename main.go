package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Msksgm/itddd-go-04-domainservice/domain/model/user"
	_ "github.com/lib/pq"
)

func main() {
	uri := fmt.Sprintf("postgres://%s/%s?sslmode=disable&user=%s&password=%s&port=%s&timezone=Asia/Tokyo",
		os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"))
	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("successfully connected to database")

	err = CreateUser(db, "test-user")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("test-user is successfully added in users table")
	}
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
		return fmt.Errorf("main.CreateUser(): %s is already exists.", name)
	}

	_, err = db.Exec("INSERT INTO users (id, username) VALUES ($1, $2)", newUser.UserId(), newUser.UserName())
	if err != nil {
		return fmt.Errorf("main.CreateUser(): %v", err)
	}
	return nil
}
