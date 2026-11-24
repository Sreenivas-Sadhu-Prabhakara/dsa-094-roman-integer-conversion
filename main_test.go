package main

import "testing"

func TestRomanToInt(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"III", 3}, {"LVIII", 58}, {"MCMXCIV", 1994}, {"IV", 4}, {"IX", 9}, {"M", 1000},
	}
	for _, c := range cases {
		if got := RomanToInt(c.in); got != c.want {
			t.Errorf("RomanToInt(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}

func TestIntToRoman(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{3, "III"}, {58, "LVIII"}, {1994, "MCMXCIV"}, {4, "IV"}, {9, "IX"}, {3999, "MMMCMXCIX"},
		{0, ""}, {4000, ""},
	}
	for _, c := range cases {
		if got := IntToRoman(c.in); got != c.want {
			t.Errorf("IntToRoman(%d) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestRoundTrip(t *testing.T) {
	for n := 1; n <= 3999; n++ {
		if got := RomanToInt(IntToRoman(n)); got != n {
			t.Fatalf("round trip failed at %d (got %d)", n, got)
		}
	}
}
