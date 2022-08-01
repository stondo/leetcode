package romantointeger

import (
	"testing"
)

// TestRomanToInteger3 calls romantointeger.RomanToInt with a string roman number, checking
// for a valid converstion to integer value.
func TestRomanToInteger3(t *testing.T) {
	roman := "III"
	i := RomanToInt(roman)
	if i != 3 {
		t.Fatalf(`RomanToInt("III") = %v, want 3`, i)
	}
}

// TestRomanToInteger58 calls romantointeger.RomanToInt with a string roman number, checking
// for a valid converstion to integer value.
func TestRomanToInteger58(t *testing.T) {
	roman := "LVIII"
	i := RomanToInt(roman)
	if i != 58 {
		t.Fatalf(`RomanToInt("LVIII") = %v, want 58`, i)
	}
}

// TestRomanToInteger1994 calls romantointeger.RomanToInt with a string roman number, checking
// for a valid converstion to integer value.
func TestRomanToInteger1994(t *testing.T) {
	roman := "MCMXCIV"
	i := RomanToInt(roman)
	if i != 1994 {
		t.Fatalf(`RomanToInt("MCMXCIV") = %v, want 1994`, i)
	}
}
