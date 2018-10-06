package converter

import (
	"fmt"
	"strings"
	"testing"
)

func TestRadixMapInitialization(t *testing.T) {
	if len(radixMap) != len(radixLiterals) {
		t.Errorf("elements count in radixMap is not equal to lenght of radixes literal, got %d, want %d", len(radixMap), len(radixLiterals))
	}

	for k, v := range radixMap {
		if !strings.ContainsRune(radixLiterals, k) {
			t.Errorf("unexpected symbol %q in radixMap", k)
		}

		if v < 1 || v > uint64(len(radixLiterals)) {
			t.Errorf("too big value in radixMap, max is %d, got %d", len(radixLiterals), v)
		}
	}
}

func TestCheckInputDataForConverting(t *testing.T) {
	data := map[struct {
		num      string
		from, to uint64
	}]error{
		{"deadbeef", uint64(len(radixLiterals) + 1), 10}:  fmt.Errorf("min radix is 2, max is %d, can not use %d", len(radixLiterals), len(radixLiterals)+1),
		{"deadbeef", uint64(len(radixLiterals) + 10), 10}: fmt.Errorf("min radix is 2, max is %d, can not use %d", len(radixLiterals), len(radixLiterals)+10),
		{"deadbeef", 16, uint64(len(radixLiterals) + 1)}:  fmt.Errorf("min radix is 2, max is %d, can not use %d", len(radixLiterals), len(radixLiterals)+1),
		{"deadbeef", 16, uint64(len(radixLiterals) + 10)}: fmt.Errorf("min radix is 2, max is %d, can not use %d", len(radixLiterals), len(radixLiterals)+10),
		{"deadbeef", 2, 10}:     fmt.Errorf("literal %q on position %d can not be used to convert number from radix %d", 'd', 1, 2),
		{"101042", 2, 10}:       fmt.Errorf("literal %q on position %d can not be used to convert number from radix %d", '4', 5, 2),
		{"101042cafe", 10, 10}:  fmt.Errorf("literal %q on position %d can not be used to convert number from radix %d", 'c', 7, 10),
		{"42_deadbeef", 10, 10}: fmt.Errorf("can not convert literal %q into number representation on position %d", '_', 3),
		{"42+deadbeef", 10, 10}: fmt.Errorf("can not convert literal %q into number representation on position %d", '+', 3),
		{"deadbeef", 0, 10}:     fmt.Errorf("min radix is 2, max is %d, can not use %d", len(radixLiterals), 0),
		{"deadbeef", 1, 10}:     fmt.Errorf("min radix is 2, max is %d, can not use %d", len(radixLiterals), 1),
		{"", 0, 0}:              fmt.Errorf("number string should not be empty"),
		{"42", 10, 10}:          nil,
		{"10101", 2, 16}:        nil,
		{"deadbeef", 16, 2}:     nil,
	}

	for arg, err := range data {
		ferr := checkInputDataForConverting(arg.num, arg.from, arg.to)
		if (ferr == nil && ferr != err) || (ferr != nil && ferr.Error() != err.Error()) {
			t.Errorf("checkInputDataForConverting(%q, %d, %d): unexpected error message, want %q, got %q", arg.num, arg.from, arg.to, err, ferr)
		}
	}
}

func TestToDecimalByGornerScheme(t *testing.T) {
	data := map[struct {
		num string
		r   uint64
	}]uint64{
		{"101", 2}:   5,
		{"111", 2}:   7,
		{"42", 5}:    22,
		{"1234", 10}: 1234,
		{"cac", 13}:  2170,
		{"dead", 16}: 57005,
	}

	for arg, r := range data {
		if fr := toDecimalByGornerScheme(arg.num, arg.r); fr != r {
			t.Errorf("toDecimalByGornerScheme(%q, %d): wrong result, want %d, got %d", arg.num, arg.r, r, fr)
		}
	}
}

func TestFromDecimalToBase(t *testing.T) {
	data := map[struct {
		num uint64
		r   uint64
	}]string{
		{5, 2}:      "101",
		{7, 2}:      "111",
		{22, 5}:     "42",
		{1234, 10}:  "1234",
		{2170, 13}:  "cac",
		{57005, 16}: "dead",
	}

	for arg, r := range data {
		if fr := fromDecimalToBase(arg.num, arg.r); fr != r {
			t.Errorf("fromDecimalToBase(%d, %d): wrong result, want %q, got %q", arg.num, arg.r, r, fr)
		}
	}
}

func TestConvertNumberToAnotherRadix(t *testing.T) {
	data := map[struct {
		num      string
		from, to uint64
	}]struct {
		result string
		err    error
	}{
		{"deadbeef", uint64(len(radixLiterals) + 1), 10}:  {"", fmt.Errorf("min radix is 2, max is %d, can not use %d", len(radixLiterals), len(radixLiterals)+1)},
		{"deadbeef", uint64(len(radixLiterals) + 10), 10}: {"", fmt.Errorf("min radix is 2, max is %d, can not use %d", len(radixLiterals), len(radixLiterals)+10)},
		{"deadbeef", 16, uint64(len(radixLiterals) + 1)}:  {"", fmt.Errorf("min radix is 2, max is %d, can not use %d", len(radixLiterals), len(radixLiterals)+1)},
		{"deadbeef", 16, uint64(len(radixLiterals) + 10)}: {"", fmt.Errorf("min radix is 2, max is %d, can not use %d", len(radixLiterals), len(radixLiterals)+10)},
		{"deadbeef", 2, 10}:     {"", fmt.Errorf("literal %q on position %d can not be used to convert number from radix %d", 'd', 1, 2)},
		{"101042", 2, 10}:       {"", fmt.Errorf("literal %q on position %d can not be used to convert number from radix %d", '4', 5, 2)},
		{"101042cafe", 10, 10}:  {"", fmt.Errorf("literal %q on position %d can not be used to convert number from radix %d", 'c', 7, 10)},
		{"42_deadbeef", 10, 10}: {"", fmt.Errorf("can not convert literal %q into number representation on position %d", '_', 3)},
		{"42+deadbeef", 10, 10}: {"", fmt.Errorf("can not convert literal %q into number representation on position %d", '+', 3)},
		{"deadbeef", 0, 10}:     {"", fmt.Errorf("min radix is 2, max is %d, can not use %d", len(radixLiterals), 0)},
		{"deadbeef", 1, 10}:     {"", fmt.Errorf("min radix is 2, max is %d, can not use %d", len(radixLiterals), 1)},
		{"", 0, 0}:              {"", fmt.Errorf("number string should not be empty")},
		{"42", 10, 10}:          {"42", nil},
		{"10101", 2, 16}:        {"15", nil},
		{"deadbeef", 16, 2}:     {"11011110101011011011111011101111", nil},
		{"12357", 8, 2}: {"1010011101111", nil},
		{"1426", 7, 13}: {"340", nil},
	}

	for arg, res := range data {
		num, err := ConvertNumberToAnotherRadix(arg.num, arg.from, arg.to)
		if res.result != num || (err != nil && res.err.Error() != err.Error()) {
			t.Errorf("ConvertNumberToAnotherRadix(%q, %d, %d): unexpected result, want %v, got %v", arg.num, arg.from, arg.to, res, struct{string; error}{num, err})
		}
	}
}
