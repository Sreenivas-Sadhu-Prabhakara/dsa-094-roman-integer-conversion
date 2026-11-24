package main

import (
	"fmt"
	"strings"
)

var romanValues = map[byte]int{
	'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
}

// RomanToInt parses a Roman numeral: a symbol smaller than its right
// neighbor subtracts (IV = 4), otherwise it adds.
func RomanToInt(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		v := romanValues[s[i]]
		if i+1 < len(s) && v < romanValues[s[i+1]] {
			total -= v
		} else {
			total += v
		}
	}
	return total
}

// intToRomanTable pairs every additive and subtractive unit, largest first.
var intToRomanTable = []struct {
	v int
	s string
}{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

// IntToRoman renders n (1..3999) as a Roman numeral by greedily taking the
// largest table entry that fits. Returns "" outside the classic range.
func IntToRoman(n int) string {
	if n < 1 || n > 3999 {
		return ""
	}
	var b strings.Builder
	for _, e := range intToRomanTable {
		for n >= e.v {
			b.WriteString(e.s)
			n -= e.v
		}
	}
	return b.String()
}

func main() {
	fmt.Println(RomanToInt("MCMXCIV"), IntToRoman(1994))
}
