package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func TestNewUser(t *testing.T) {
	uuidV4 := uuid.New().String()
	userId, err := NewUserId(uuidV4)
	if err != nil {
		t.Fatal(err)
	}

	name := "userName"
	userName, err := NewUserName(name)
	if err != nil {
		t.Fatal(err)
	}

	user, err := NewUser(*userId, *userName)
	if err != nil {
		t.Fatal(err)
	}

	wantUser := &User{userId: *userId, userName: *userName}
	if diff := cmp.Diff(wantUser, user, cmp.AllowUnexported(User{}, UserName{}, UserId{})); diff != "" {
		t.Errorf("mismatch (-wantUser, +got):\n%s", diff)
	}
}
