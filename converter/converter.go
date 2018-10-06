package converter

import (
	"fmt"
	"math"
	"unicode"
	"unicode/utf8"
)

const (
	radixLiterals = "0123456789abcdef"
)

var (
	radixMap = map[rune]uint64{}
)

func ConvertNumberToAnotherRadix(number string, from, to uint64) (string, error) {
	if err := checkInputDataForConverting(number, from, to); err != nil {
		return "", err
	}
	return fromDecimalToBase(toDecimalByGornerScheme(number, from), to), nil
}

func checkInputDataForConverting(number string, from, to uint64) error {
	if number == "" {
		return fmt.Errorf("number string should not be empty")
	} else if from < 2 || from > uint64(len(radixLiterals)) {
		return fmt.Errorf("min radix is 2, max is %d, can not use %d", len(radixLiterals), from)
	} else if to < 2 || to > uint64(len(radixLiterals)) {
		return fmt.Errorf("min radix is 2, max is %d, can not use %d", len(radixLiterals), to)
	}

	for i, c := range number {
		r, ok := radixMap[unicode.ToLower(c)]
		if !ok {
			return fmt.Errorf("can not convert literal %q into number representation on position %d", c, i+1)
		} else if r > from {
			return fmt.Errorf("literal %q on position %d can not be used to convert number from radix %d", c, i+1, from)
		}
	}

	return nil
}

func fromDecimalToBase(number uint64, to uint64) string {
	result := ""
	for number != 0 {
		mod := number % to
		number /= to
		for k, v := range radixMap {
			if mod+1 == v {
				result = string(k) + result
			}
		}
	}
	return result
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

func init() {
	for i, c := range radixLiterals {
		radixMap[c] = uint64(i + 1)
	}
}
