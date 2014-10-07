package phone

import (
	"fmt"
	"regexp"
	"strings"
)

func GetDefaultISO3166() *PhoneData {
	return ISO3166_Data[0]
}

func GetISO3166(country string) *PhoneData {
	var res *PhoneData
	switch len(country) {
	case 0:
		res = ISO3166_Data[0]
		break
	case 2:
		for _, v := range ISO3166_Data {
			if country == v.Alpha2 {
				res = v
				break
			}
		}
	case 3:
		for _, v := range ISO3166_Data {
			if country == v.Alpha3 {
				res = v
				break
			}
		}
	default:
		lower := strings.ToLower(country)
		for _, v := range ISO3166_Data {
			if lower == strings.ToLower(v.CountryName) {
				res = v
				break
			}
		}
	}
	return res
}

func GetISO3166ByPhone(phone string) (*PhoneData, error) {
	lenOfPhoneNum := len(phone)

	for _, v := range ISO3166_Data {
		if reg, err := regexp.Compile(fmt.Sprintf("^%s", v.CountryCode)); err == nil && reg.MatchString(phone) {

			lenOfCountryCode := len(v.CountryCode)

			for _, x := range v.PhoneNumberLengths {
				if lenOfPhoneNum == (lenOfCountryCode + x) {
					// it match.. but may have more than one result.
					// e.g. USA and Canada. need to check mobileBeginWith

					for _, y := range v.MobileBeginsWith {
						if reg, err := regexp.Compile(fmt.Sprintf("^%s%s", v.CountryCode, y)); err == nil {
							if reg.MatchString(phone) {
								return v, nil
							}
						} else {
							return nil, err
						}
					}
				}
			}

		} else if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func IsValidPhoneISO3166(p string, data *PhoneData) bool {
	if reg, err := regexp.Compile(fmt.Sprintf("^%s", data.CountryCode)); err == nil {
		phone := reg.ReplaceAllString(p, "")
		lenOfPhoneNum := len(phone)
		for _, v := range data.PhoneNumberLengths {
			if lenOfPhoneNum == v {
				for _, x := range data.MobileBeginsWith {
					if reg, err := regexp.Compile(fmt.Sprintf("^%s", x)); err == nil {
						if reg.MatchString(phone) {
							return true
						}
					} else {
						return false
					}
				}
			}
		}
	}
	return false
}
