package transport

import (
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewBagge(t *testing.T) {
	id := "id"
	baggageId, err := NewBaggageId(id)
	if err != nil {
		log.Fatal(err)
	}

	got, err := NewBaggage(baggageId)
	if err != nil {
		log.Fatal(err)
	}
	want := &Baggage{baggageId: *baggageId}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Baggage{}, BaggageId{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
