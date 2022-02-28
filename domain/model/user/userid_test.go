package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func TestNewUserId(t *testing.T) {
	newUuid := uuid.New().String()
	userId, err := NewUserId(newUuid)
	if err != nil {
		t.Fatal(err)
	}

	got := userId
	want := &UserId{id: newUuid}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(UserId{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
