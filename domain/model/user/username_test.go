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
}
