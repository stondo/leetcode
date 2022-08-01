package integertoroman

import (
	"testing"
)

// TestIntegerToRoman3 calls integertoroman.IntToRoman with an int number, checking
// for a valid converstion to roman value.
func TestIntegerToRoman3(t *testing.T) {
	integer := 3
	i := IntToRoman(integer)
	if i != "III" {
		t.Fatalf(`IntToRoman(3) = %v, want III`, i)
	}
}

// TestIntegerToRoman58 calls integertoroman.IntToRoman with an int number, checking
// for a valid converstion to roman value.
func TestIntegerToRoman58(t *testing.T) {
	integer := 58
	i := IntToRoman(integer)
	if i != "LVIII" {
		t.Fatalf(`IntToRoman(58) = %v, want LVIII`, i)
	}
}

// TestIntegerToRoman1994 calls integertoroman.IntToRoman with an int number, checking
// for a valid converstion to roman value.
func TestIntegerToRoman1994(t *testing.T) {
	integer := 1994
	i := IntToRoman(integer)
	if i != "MCMXCIV" {
		t.Fatalf(`IntToRoman(1994) = %v, want MCMXCIV`, i)
	}
}
