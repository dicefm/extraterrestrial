package main

import "fmt"

type PhoneData struct {
	Alpha2             string
	Alpha3             string
	CountryCode        string
	CountryName        string
	MobileBeginsWith   []string
	PhoneNumberLengths []int
}

type Result struct {
	PhoneNumber string
	PhoneData   *PhoneData
}

func NewResult(phone string, data *PhoneData) *Result {
	return &Result{
		PhoneNumber: phone,
		PhoneData:   data,
	}
}

func (r Result) WithPlusSign() *Result {
	return NewResult(fmt.Sprintf("+%s", r.PhoneNumber), r.PhoneData)
}

func (r Result) AsPhoneResult() *PhoneResult {
	return NewPhoneResult(r.PhoneNumber, r.PhoneData.Alpha3)
}

type PhoneResult struct {
	PhoneNumber string
	Country     string
}

func NewPhoneResult(phone, country string) *PhoneResult {
	return &PhoneResult{
		PhoneNumber: phone,
		Country:     country,
	}
}
