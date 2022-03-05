package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func TestNewUserId(t *testing.T) {
	uuidV4 := uuid.New().String()
	userId, err := NewUserId(uuidV4)
	if err != nil {
		t.Fatal(err)
	}

	got := userId
	want := &UserId{id: uuidV4}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(UserId{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestUserIdEquals(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		uuidV4 := uuid.New().String()
		userId1, err := NewUserId(uuidV4)
		if err != nil {
			t.Fatal(err)
		}
		userId2, err := NewUserId(uuidV4)
		if err != nil {
			t.Fatal(err)
		}

		if !userId1.Equals(userId2) {
			t.Errorf("userId1: %v must be equal to userId2: %v", userId1, userId2)
		}
	})

	t.Run("not equal", func(t *testing.T) {
		uuidV4 := uuid.New().String()
		userId1, err := NewUserId(uuidV4)
		if err != nil {
			t.Fatal(err)
		}
		uuidV4 = uuid.New().String()
		userId2, err := NewUserId(uuidV4)
		if err != nil {
			t.Fatal(err)
		}

		if userId1.Equals(userId2) {
			t.Errorf("user1: %v must not be equal to user2: %v", userId1, userId2)
		}
	})
}
