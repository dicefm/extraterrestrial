package main

import (
	"fmt"
	"testing"
)

func Test_Normaliser(t *testing.T) {
	res, err := Normalise("+44712345678", "GB")
	fmt.Println(res, err)
	if res == nil || err != nil {
		t.Error(err)
	}
}
