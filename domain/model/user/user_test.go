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

func TestUserEquals(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		name := "userName"
		userName, err := NewUserName(name)
		if err != nil {
			t.Fatal(err)
		}

		user, err := NewUser(*userName)
		if err != nil {
			t.Fatal(err)
		}

		// duplicate by assign otherUser bacause uuid of UserId cannnot duplicate.
		otherUser := user

		if !user.Equals(otherUser) {
			t.Errorf("user %v must be equal to otherUser: %v", user, otherUser)
		}
	})
	t.Run("not equal", func(t *testing.T) {
		name1 := "userName1"
		userName1, err := NewUserName(name1)
		if err != nil {
			t.Fatal(err)
		}

		user1, err := NewUser(*userName1)
		if err != nil {
			t.Fatal(err)
		}

		name2 := "userName2"
		userName2, err := NewUserName(name2)
		if err != nil {
			t.Fatal(err)
		}

		user2, err := NewUser(*userName2)
		if err != nil {
			t.Fatal(err)
		}

		if user1.Equals(user2) {
			t.Errorf("userI %v must not be equal to user2: %v", user1, user2)
		}
	})
}
