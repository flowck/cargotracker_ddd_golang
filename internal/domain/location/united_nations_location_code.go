package location

import (
	"errors"

	"github.com/biter777/countries"
)

type UnitedNationsLocationCode struct {
	value string
}

func newUnitedNationsLocationCode(countryNameOrCode string) (UnitedNationsLocationCode, error) {
	if countryNameOrCode == "" {
		return UnitedNationsLocationCode{}, errors.New("country/code cannot be empty")
	}

	c := countries.ByName(countryNameOrCode)
	if c == countries.Unknown {
		return UnitedNationsLocationCode{}, errors.New("country cannot be invalid")
	}

	return UnitedNationsLocationCode{value: c.Alpha2()}, nil
}

func NewUnitedNationsLocationCodeFromCountry(country string) (UnitedNationsLocationCode, error) {
	return newUnitedNationsLocationCode(country)
}

func NewUnitedNationsLocationCodeFromCountryCode(code string) (UnitedNationsLocationCode, error) {
	return newUnitedNationsLocationCode(code)
}

func (c UnitedNationsLocationCode) IsZero() bool {
	return c.value == ""
}
