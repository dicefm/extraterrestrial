package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/dicefm/extraterrestrial/phone"
)

var (
	DefaultPhoneCode string = "GBR"
)

func main() {
	val, _ := phone.Normalise("+852 6509-8900", "USA")
	fmt.Println("~ Normaliser.", val)

	rand.Seed(time.Now().UnixNano())

	numbers := [][]string{}

	for i := 0; i < 50; i++ {
		fmt.Println("Running test : ", i)
		numbers = append(numbers, test("+%v%v", true))
	}

	for i := 0; i < 50; i++ {
		fmt.Println("Running test : ", i)
		numbers = append(numbers, test("%v%v", false))
	}

	fmt.Println("Total numbers checked: ", len(numbers[0])*len(numbers))
}

func test(format string, validate bool) []string {
	numbers := []string{}
	for _, v := range phone.ISO3166_PhoneData {
		code := v.CountryData.CountryCode

		for _, x := range v.MobileBeginsWith {
			for _, y := range v.PhoneNumberLengths {
				num := genPhoneNumber(x, y)
				x, err := NormalisePhone(fmt.Sprintf(format, code, num))
				numbers = append(numbers, x)
				if err != nil {
					if validate || (!validate && err != phone.ErrNotFound) {
						fmt.Println(x)
						panic(err)
					}
				}
			}
		}
	}
	return numbers
}

func genPhoneNumber(begin string, length int) string {
	diff := length - len(begin)
	rest := []string{}
	for i := 0; i < diff; i++ {
		rnd := rand.Intn(10)
		rest = append(rest, fmt.Sprintf("%v", rnd))
	}
	return fmt.Sprintf("%s%s", begin, strings.Join(rest, ""))
}

func NormalisePhone(p string) (string, error) {
	// Try to normalise just on phone number
	data, err := phone.Normalise(p, "")
	if err == nil {
		return data.PhoneNumber, nil
	} else {
		// If not try an normalise on the default locale
		data, err = phone.Normalise(p, DefaultPhoneCode)
		if err == nil {
			return data.PhoneNumber, nil
		}
		// Return the error.
		return "", err
	}
}
