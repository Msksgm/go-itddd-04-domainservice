package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewUser(t *testing.T) {
	name := "userName"
	userName, err := NewUserName(name)
	if err != nil {
		t.Fatal(err)
	}

	got, err := NewUser(*userName)
	if err != nil {
		t.Fatal(err)
	}

	want := &User{userName: *userName}
	opts := cmp.Options{
		cmp.AllowUnexported(User{}, UserName{}, UserId{}),
		cmpopts.IgnoreFields(*got, "userId"),
	}
	if diff := cmp.Diff(want, got, opts); diff != "" {
		t.Errorf("mismatch (-wantUser, +got):\n%s", diff)
	}
}

func TestUserName(t *testing.T) {
	name := "userName"
	userName, err := NewUserName(name)
	if err != nil {
		t.Fatal(err)
	}

	user, err := NewUser(*userName)
	if err != nil {
		t.Fatal(err)
	}

	want := name
	if got := user.UserName(); got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
