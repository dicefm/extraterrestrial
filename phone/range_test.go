package phone

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

const (
	NumOfAttempts    int    = 10
	DefaultPhoneCode string = "GBR"
)

func Test_RangeTestAllTheNumbersWithCountryCode(t *testing.T) {
	numbers := [][]string{}

	for i := 0; i < NumOfAttempts; i++ {
		fmt.Println("Running test : ", i)
		numbers = append(numbers, test("+%v%v", true))
	}
}

func Test_RangeTestAllTheNumbersWithoutCountryCode(t *testing.T) {
	numbers := [][]string{}

	for i := 0; i < NumOfAttempts; i++ {
		fmt.Println("Running test : ", i)
		numbers = append(numbers, test("%v%v", false))
	}
}

func test(format string, validate bool) []string {
	numbers := []string{}
	for _, v := range ISO3166_PhoneData {
		code := v.CountryData.CountryCode

		for _, x := range v.MobileBeginsWith {
			for _, y := range v.PhoneNumberLengths {
				num := genPhoneNumber(x, y)
				x, err := NormalisePhone(fmt.Sprintf(format, code, num))
				numbers = append(numbers, x)
				if err != nil {
					if validate || (!validate && err != ErrNotFound) {
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
	data, err := Normalise(p, "")
	if err == nil {
		return data.PhoneNumber, nil
	} else {
		// If not try an normalise on the default locale
		data, err = Normalise(p, DefaultPhoneCode)
		if err == nil {
			return data.PhoneNumber, nil
		}
		// Return the error.
		return "", err
	}
}
