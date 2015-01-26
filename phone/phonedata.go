package phone

import (
	"github.com/dicefm/extraterrestrial/countries"
)

type PhoneData struct {
	CountryData        countries.CountryData
	PhoneCode          string
	MobileBeginsWith   []string
	PhoneNumberLengths []int
}

func NewPhoneData() *PhoneData {
	return &PhoneData{
		CountryData:        countries.NewCountryData(),
		MobileBeginsWith:   make([]string, 0, 0),
		PhoneNumberLengths: make([]int, 0, 0),
	}
}
