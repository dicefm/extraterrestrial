package phone

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	PLUS_SIGN_REGEX       = `(?i)^\+`
	NON_DIGIT_CHAR        = `(?i)\D`
	ZERO_DIGIT_CHARS      = `^0+`
	EIGHT_DIGIT_CHARS     = `^8+`
	RUS_START_DIGIT_CHARS = `^89`
)

var (
	NON_LEADING_ZERO_COUNTRIES = []string{"GAB", "CIV", "COG"}

	ErrNotFound  = fmt.Errorf("No valid format found.")
	ErrPhoneMiss = fmt.Errorf("Unable to locate data from phone number.")
)

func Normalise(p, c string) (*PhoneResult, error) {
	var (
		phone   string
		country string
		err     error
	)

	hasPlusSign := containsPlusSign(p)

	phone, err = normalisePhoneNumber(p)
	if err != nil {
		return nil, err
	}

	country, err = normaliseCountry(c)
	if err != nil {
		return nil, err
	}

	// if no country, default is USA
	var iso3166 *PhoneData
	if country != "" {
		iso3166 = GetISO3166(country)
	}

	var result *Result
	if iso3166 != nil {
		result, err = normaliseWithCountry(phone, iso3166, hasPlusSign)
	} else {
		result, err = normaliseWithoutCountry(phone, hasPlusSign)
	}

	if err != nil {
		return nil, err
	}

	if result.PhoneData == nil {
		return nil, ErrPhoneMiss
	}

	if IsValidPhoneISO3166(result.PhoneNumber, result.PhoneData) {
		return result.WithPlusSign().AsPhoneResult(), nil
	}

	return result.AsPhoneResult(), ErrNotFound
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

func containsString(haystack []string, needle string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func containsInt(haystack []int, needle int) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func removeZeros(phone string) (string, error) {
	reg, err := regexp.Compile(ZERO_DIGIT_CHARS)
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(phone, ""), nil
}

func isRussianPhoneNum(phone, country string) bool {
	if country == "RUS" && len(phone) == 11 {
		if matched, err := regexp.MatchString(RUS_START_DIGIT_CHARS, phone); err == nil {
			return matched
		}
		return false
	}
	return false
}

func cleanRussianPhoneNum(phone string) (string, error) {
	reg, err := regexp.Compile(EIGHT_DIGIT_CHARS)
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(phone, ""), nil
}

func normaliseWithCountry(phone string, data *PhoneData, hasPlusSign bool) (*Result, error) {
	var err error
	if !containsString(NON_LEADING_ZERO_COUNTRIES, data.CountryData.Alpha3) {
		if phone, err = removeZeros(phone); err != nil {
			return nil, err
		}
	}

	if isRussianPhoneNum(phone, data.CountryData.Alpha3) {
		if phone, err = cleanRussianPhoneNum(phone); err != nil {
			return nil, err
		}
	}

	if !hasPlusSign {
		lenOfPhoneNum := len(phone)
		if containsInt(data.PhoneNumberLengths, lenOfPhoneNum) {
			phone = fmt.Sprintf("%s%s", data.CountryData.CountryCode, phone)
		}
	}

	return NewResult(phone, data), nil
}

func normaliseWithoutCountry(phone string, hasPlusSign bool) (*Result, error) {
	// Let's try and brute force it anyway
	if possible, err := GetISO3166ByPhone(phone); err == nil && possible != nil {
		return NewResult(phone, possible), nil
	}

	if hasPlusSign {
		return nil, ErrPhoneMiss
	}

	// We really have no idea what this phone number is!
	var (
		data          = GetDefaultISO3166()
		lenOfPhoneNum = len(phone)
	)
	if containsInt(data.PhoneNumberLengths, lenOfPhoneNum) {
		phone = fmt.Sprintf("1%s", phone)
	}

	return NewResult(phone, data), nil
}
