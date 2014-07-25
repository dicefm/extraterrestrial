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
