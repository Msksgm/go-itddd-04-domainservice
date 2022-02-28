package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewUserName(t *testing.T) {
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
}
