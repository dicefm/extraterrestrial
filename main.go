package main

import (
	"fmt"

	"github.com/dicefm/extraterrestrial/phone"
)

func main() {
	val, err := phone.Normalise("+852 6509-8900", "USA")
	fmt.Println("~ Normaliser.", val, err)
}
