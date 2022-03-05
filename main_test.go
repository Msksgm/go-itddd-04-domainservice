package main

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

const (
	id   = "id"
	name = "userName"
)

func TestCreateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE username = $1`)).
			WithArgs(name).
			WillReturnRows(mock.NewRows([]string{"id", "username"}))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO users").
			WithArgs(sqlmock.AnyArg(), name).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		if err = CreateUser(db, name); err != nil {
			t.Errorf("error was not expected CreateUser() %s:", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	t.Run("fail already exists", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE username = $1`)).
			WithArgs(name).
			WillReturnRows(mock.NewRows([]string{"id", "username"}).AddRow(id, name))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectCommit()

		want := fmt.Sprintf("%s is already exists.", name)
		err = CreateUser(db, name)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
