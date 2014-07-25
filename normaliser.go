package main

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	PLUS_SIGN_REGEX = `(?i)^+`
	NON_DIGIT_CHAR  = `(?i)\D`
)

func Normalise(p, c string) (*Result, error) {
	var (
		phone   string
		country string
		err     error
	)

	phone, err = normalisePhoneNumber(p)
	if err != nil {
		return nil, err
	}

	country, err = normaliseCountry(c)
	if err != nil {
		return nil, err
	}

	plusSign := containsPlusSign(phone)
	fmt.Println("PlusSign ", plusSign)

	// if no country, default is USA
	iso3166 := getISO3166(country)

	fmt.Println("PhoneData ", iso3166)
	return nil, nil
}

func normalisePhoneNumber(phone string) (string, error) {
	reg, err := regexp.Compile(NON_DIGIT_CHAR)
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(strings.TrimSpace(phone), ""), nil
}

func normaliseCountry(country string) (string, error) {
	return strings.ToUpper(strings.TrimSpace(country)), nil
}

func containsPlusSign(phone string) bool {
	if matched, err := regexp.MatchString(PLUS_SIGN_REGEX, phone); err == nil {
		return matched
	}
	return false
}

func getISO3166(country string) *PhoneData {
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
