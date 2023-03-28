package voyage

import "errors"

//
// A valueObject to identify a voyage.
//

type Number struct {
	value string
}

func NewNumber(value string) (Number, error) {
	if value == "" {
		return Number{}, errors.New("number cannot be empty")
	}

	return Number{value: value}, nil
}

func (n Number) String() string {
	return n.value
}

func (n Number) IsEmpty() bool {
	return n.value == ""
}
