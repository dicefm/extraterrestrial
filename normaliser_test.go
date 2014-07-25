package main

import (
	"fmt"
	"testing"
)

func Test_Normaliser(t *testing.T) {
	res, err := Normalise("0712345678", "GB")
	fmt.Println(res.PhoneNumber, err)
	if res != nil || err != nil {
		t.Error(err)
	}
}
