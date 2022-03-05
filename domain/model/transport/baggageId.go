package transport

type BaggageId struct {
	id string
}

func NewBaggageId(id string) (*BaggageId, error) {
	return &BaggageId{id: id}, nil
}
