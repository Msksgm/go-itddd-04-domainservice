package user

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

func TestExists(t *testing.T) {
	uuidV4 := uuid.New().String()
	name := "userName"
	user := &User{userId: UserId{id: uuidV4}, userName: UserName{name: name}}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("true", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE name = $1`)).
			WithArgs("userName").
			WillReturnRows(mock.NewRows([]string{"id", "name"}).AddRow(uuidV4, "userName"))
		mock.ExpectCommit()

		isExists, err := Exists(db, user)
		if err != nil {
			t.Errorf("error was not expected: %s", err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
		if !isExists {
			t.Errorf("isExists: must be true, not false")
		}
	})
	t.Run("false", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE name = $1`)).
			WithArgs("userName").
			WillReturnRows(mock.NewRows([]string{}))
		mock.ExpectCommit()

		isExists, err := Exists(db, user)
		if err != nil {
			t.Errorf("error was not expected: %s", err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
		if isExists {
			t.Errorf("isExists: must be false, not false")
		}
	})
}
