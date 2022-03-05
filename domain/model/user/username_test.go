package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewUserName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		name := "userName"
		userName, err := NewUserName(name)
		if err != nil {
			t.Fatal(err)
		}

		got := userName
		want := &UserName{name: name}
		if diff := cmp.Diff(want, got, cmp.AllowUnexported(UserName{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail name is empty", func(t *testing.T) {
		name := ""
		_, err := NewUserName(name)
		want := "name is required."
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail name is less than 3 characters", func(t *testing.T) {
		name := "na"
		_, err := NewUserName(name)
		want := "name must not be less than 3 characters."
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestName(t *testing.T) {
	name := "userName"
	userName, err := NewUserName(name)
	if err != nil {
		t.Fatal(err)
	}

	got := userName.Name()
	want := name
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestEquals(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		name := "userName"
		userName1, err := NewUserName(name)
		if err != nil {
			t.Fatal(err)
		}
		userName2, err := NewUserName(name)
		if err != nil {
			t.Fatal(err)
		}

		if !userName1.Equals(userName2) {
			t.Errorf("userName1: %v must be equal to userName2: %v", userName1, userName2)
		}
	})
	t.Run("not equal", func(t *testing.T) {
		name := "userName1"
		userName1, err := NewUserName(name)
		if err != nil {
			t.Fatal(err)
		}
		name = "userName2"
		userName2, err := NewUserName(name)
		if err != nil {
			t.Fatal(err)
		}

		if userName1.Equals(userName2) {
			t.Errorf("userName1: %v must not be equal to userName2: %v", userName1, userName2)
		}
	})
}
