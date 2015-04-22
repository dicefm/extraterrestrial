package phone

import (
	"fmt"
	"testing"
	"testing/quick"
)

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
		return fmt.Sprintf("4366480%v9%v2%v0%v", mod, mod, mod, mod)
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

func Test_Phone_IRL_without_country(t *testing.T) {
	makeNum := func(a uint) string {
		mod := a % 1
		return fmt.Sprintf("35387%v09%v54%v", mod, mod, mod)
	}
	prepare := func(a string) string {
		res, _ := normalisePhoneNumber(a)
		return fmt.Sprintf("+%s", res)
	}

	f := func(x uint) *PhoneResult {
		return NewPhoneResult(prepare(makeNum(x)), "IRL")
	}
	g := func(x uint) *PhoneResult {
		res, _ := Normalise(makeNum(x), "")
		return res
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_Phone_IRL_with_country(t *testing.T) {
	makeNum := func(a uint) string {
		mod := a % 1
		return fmt.Sprintf("+35387%v09%v54%v", mod, mod, mod)
	}
	prepare := func(a string) string {
		res, _ := normalisePhoneNumber(a)
		return fmt.Sprintf("+%s", res)
	}

	f := func(x uint) *PhoneResult {
		return NewPhoneResult(prepare(makeNum(x)), "IRL")
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
