package phone

import (
	"github.com/dicefm/extraterrestrial/countries"
)

type PhoneData struct {
	CountryData        countries.CountryData
	PhoneCode          string
	MobileBeginsWith   []string
	PhoneNumberLengths []int
	RemoveLeadingZeros bool
}

func NewPhoneData() *PhoneData {
	return &PhoneData{
		CountryData:        countries.CountryData{},
		MobileBeginsWith:   make([]string, 0, 0),
		PhoneNumberLengths: make([]int, 0, 0),
		RemoveLeadingZeros: false,
	}
}
