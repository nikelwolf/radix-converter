package converter

import (
	"fmt"
	"math"
	"unicode"
	"unicode/utf8"

	"github.com/therecipe/qt/core"
)

const (
	radixLiterals = "0123456789abcdef"
)

var (
	radixMap map[rune]uint64 = map[rune]uint64{}
)

type RadixConverter struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(string, uint64, uint64) string `slot:"convertButtonClicked"`
}

func (rc *RadixConverter) init() {
	rc.ConnectConvertButtonClicked(rc.convertButtonClickHandler)
}

func (rc *RadixConverter) convertButtonClickHandler(number string, from, to uint64) string {
	if err := checkNumberWithRadixValid(number, from); err != nil {
		return err.Error()
	}

	result, err := convertNumberToAnotherRadix(number, from, to)
	if err != nil {
		return err.Error()
	}
	return result
}

func checkNumberWithRadixValid(number string, radix uint64) error {
	for i, c := range number {
		r, ok := radixMap[unicode.ToLower(c)]
		if !ok {
			return fmt.Errorf("can not convert literal %q into number representation on position %d", c, i + 1)
		} else if r > radix {
			return fmt.Errorf("literal %q on position %d can not be used to convert number from radix %d", c, i + 1, radix)
		}
	}

	return nil
}

func toDecimalByGornerScheme(number string, from uint64) uint64 {
	var result uint64
	for i, c := range number {
		power := utf8.RuneCountInString(number) - i - 1
		digit := radixMap[unicode.ToLower(c)] - 1
		result += digit * uint64(math.Pow(float64(from), float64(power)))
	}
	return result
}

func fromDecimalToBase(number uint64, to uint64) (string, error) {
	result := ""
	for number != 0 {
		mod := number % to
		number /= to
		for k, v := range radixMap {
			if mod + 1 == v {
				result += string(k)
			}
		}
	}
	return reverseString(result), nil
}

func convertNumberToAnotherRadix(number string, from, to uint64) (string, error) {
	return fromDecimalToBase(toDecimalByGornerScheme(number, from), to)
}

func reverseString(s string) string {
	cs := make([]rune, utf8.RuneCountInString(s))
	i := len(cs)
	for _, c := range s {
		i--
		cs[i] = c
	}
	return string(cs)
}

func init() {
	for i, c := range radixLiterals {
		radixMap[c] = uint64(i + 1)
	}
}