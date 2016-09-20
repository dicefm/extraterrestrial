package phone

import (
	"fmt"
	"testing"
	"testing/quick"
)

func Test_InvalidPhone_UK(t *testing.T) {
	var (
		number = "InvalidPhone"
		_, err = Normalise(number, "GBR")
	)

	if err == nil {
		t.Error("Error shouldn't be nil")
	}
}

func Test_Phone_GB_with_country_and_leading_zero(t *testing.T) {
	makeNum := func(a uint) (string, string) {
		mod := a % 1
		return fmt.Sprintf("+44748%v09%v54%v", mod, mod, mod),
			fmt.Sprintf("+440000748%v09%v54%v", mod, mod, mod)
	}
	prepare := func(a string) string {
		res, _ := normalisePhoneNumber(a)
		return fmt.Sprintf("+%s", res)
	}

	f := func(x uint) *PhoneResult {
		a, _ := makeNum(x)
		return NewPhoneResult(prepare(a), "GBR")
	}
	g := func(x uint) *PhoneResult {
		_, b := makeNum(x)
		res, _ := Normalise(b, "GBR")
		return res
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Phone_GB_without_country(t *testing.T) {
	makeNum := func(a uint) string {
		mod := a % 1
		return fmt.Sprintf("+44748%v09%v54%v", mod, mod, mod)
	}
	prepare := func(a string) string {
		res, _ := normalisePhoneNumber(a)
		return fmt.Sprintf("+%s", res)
	}

	f := func(x uint) *PhoneResult {
		return NewPhoneResult(prepare(makeNum(x)), "GBR")
	}
	g := func(x uint) *PhoneResult {
		res, _ := Normalise(makeNum(x), "")
		return res
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Phone_HKG_with_country(t *testing.T) {
	makeNum := func(a uint) string {
		mod := a % 10
		return fmt.Sprintf("+852 65%v9-890%v", mod, mod)
	}
	prepare := func(a string) string {
		res, _ := normalisePhoneNumber(a)
		return fmt.Sprintf("+%s", res)
	}

	f := func(x uint) *PhoneResult {
		return NewPhoneResult(prepare(makeNum(x)), "HKG")
	}
	g := func(x uint) *PhoneResult {
		res, _ := Normalise(makeNum(x), "HKG")
		return res
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Phone_1_USA_with_country(t *testing.T) {
	makeNum := func(a uint) string {
		mod := a % 10
		return fmt.Sprintf("(817) 5%v9-890%v", mod, mod)
	}
	prepare := func(a string) string {
		res, _ := normalisePhoneNumber(a)
		return fmt.Sprintf("+1%s", res)
	}

	f := func(x uint) *PhoneResult {
		return NewPhoneResult(prepare(makeNum(x)), "USA")
	}
	g := func(x uint) *PhoneResult {
		res, _ := Normalise(makeNum(x), "USA")
		return res
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Phone_2_USA_without_country(t *testing.T) {
	makeNum := func(a uint) string {
		mod := a % 10
		return fmt.Sprintf("(817) 5%v9-890%v", mod, mod)
	}
	prepare := func(a string) string {
		res, _ := normalisePhoneNumber(a)
		return fmt.Sprintf("+1%s", res)
	}

	f := func(x uint) *PhoneResult {
		return NewPhoneResult(prepare(makeNum(x)), "USA")
	}
	g := func(x uint) *PhoneResult {
		res, _ := Normalise(makeNum(x), "")
		return res
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Phone_AUT_without_country(t *testing.T) {
	makeNum := func(a uint) string {
		mod := a % 10
		return fmt.Sprintf("+4366480%s9%s2%s0%s", mod, mod)
	}
	prepare := func(a string) string {
		res, _ := normalisePhoneNumber(a)
		return fmt.Sprintf("+%s", res)
	}

	f := func(x uint) *PhoneResult {
		return NewPhoneResult(prepare(makeNum(x)), "AUT")
	}
	g := func(x uint) *PhoneResult {
		res, _ := Normalise(makeNum(x), "")
		return res
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Phone_USA_HKG_missmatch(t *testing.T) {
	makeNum := func(a uint) string {
		mod := a % 10
		return fmt.Sprintf("(817) 5%v9-890%v", mod, mod)
	}

	f := func(x uint) error {
		return ErrNotFound
	}
	g := func(x uint) error {
		_, err := Normalise(makeNum(x), "HKG")
		return err
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Phone_invalid_HKG(t *testing.T) {
	makeNum := func(a uint) string {
		mod := a % 10
		return fmt.Sprintf("+1(817) 5%v9-890%v", mod, mod)
	}

	f := func(x uint) error {
		return ErrNotFound
	}
	g := func(x uint) error {
		_, err := Normalise(makeNum(x), "HKG")
		return err
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Phone_valid_USA(t *testing.T) {
	makeNum := func(a uint) string {
		mod := a % 10
		return fmt.Sprintf("+1(817) 5%v9-890%v", mod, mod)
	}
	prepare := func(a string) string {
		res, _ := normalisePhoneNumber(a)
		return fmt.Sprintf("+%s", res)
	}

	f := func(x uint) *PhoneResult {
		return NewPhoneResult(prepare(makeNum(x)), "USA")
	}
	g := func(x uint) *PhoneResult {
		res, _ := Normalise(makeNum(x), "")
		return res
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Phone_valid_USA_2(t *testing.T) {
	makeNum := func(a uint) string {
		mod := a % 10
		return fmt.Sprintf("(817) 5%v9-890%v", mod, mod)
	}
	prepare := func(a string) string {
		res, _ := normalisePhoneNumber(a)
		return fmt.Sprintf("+1%s", res)
	}

	f := func(x uint) *PhoneResult {
		return NewPhoneResult(prepare(makeNum(x)), "USA")
	}
	g := func(x uint) *PhoneResult {
		res, _ := Normalise(makeNum(x), "")
		return res
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Phone_invalid_USA_3(t *testing.T) {
	makeNum := func(a uint) string {
		mod := a % 10
		return fmt.Sprintf("6123-61%v%v", mod, mod)
	}

	f := func(x uint) error {
		return ErrNotFound
	}
	g := func(x uint) error {
		_, err := Normalise(makeNum(x), "")
		return err
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Phone_valid_USA_3(t *testing.T) {
	makeNum := func(a uint) string {
		mod := a % 10
		return fmt.Sprintf("6123-61%v%v", mod, mod)
	}
	prepare := func(a string) string {
		res, _ := normalisePhoneNumber(a)
		return fmt.Sprintf("+852%s", res)
	}

	f := func(x uint) *PhoneResult {
		return NewPhoneResult(prepare(makeNum(x)), "HKG")
	}
	g := func(x uint) *PhoneResult {
		res, _ := Normalise(makeNum(x), "HKG")
		return res
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Phone_Valid_Argentina_1(t *testing.T) {
	var (
		number    = "+54 9 11 61231234"
		phone, _  = normalisePhoneNumber(number)
		data, err = GetISO3166ByPhone(phone)
	)

	if data == nil && err == nil {
		t.Error("We got no matching for number")
	}
}

func Test_Phone_Valid_Argentina_2(t *testing.T) {
	var (
		number   = "+54 11 61231234"
		phone, _ = normalisePhoneNumber(number)
		data, _  = GetISO3166ByPhone(phone)
	)

	if data == nil {
		t.Error("No matching for number")
	} else if "Argentina" != data.CountryData.Name {
		message := fmt.Sprintf("Country should be Argentina, but got %s", data.CountryData.Name)
		t.Error(message)
	}
}

func Test_Phone_Valid_Argentina_3(t *testing.T) {
	var (
		number = "+54 9 11 61231234"
		_, err = Normalise(number, "")
	)

	if err != nil {
		t.Error(err)
	}
}
