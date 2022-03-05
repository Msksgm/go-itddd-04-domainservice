package transport

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewBaggageId(t *testing.T) {
	id := "id"
	got, err := NewBaggageId(id)
	if err != nil {
		t.Fatal(err)
	}
	want := &BaggageId{id: id}

	if diff := cmp.Diff(want, got, cmp.AllowUnexported(BaggageId{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
