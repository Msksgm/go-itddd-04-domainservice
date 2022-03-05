package transport

type Baggage struct {
	baggageId BaggageId
}

func NewBaggage(id *BaggageId) (*Baggage, error) {
	return &Baggage{baggageId: *id}, nil
}
