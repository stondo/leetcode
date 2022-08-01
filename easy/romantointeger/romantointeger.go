package romantointeger

import "fmt"

var romanToIntKnownMap = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

/*
	I can be placed before V (5) and X (10) to make 4 and 9.
	X can be placed before L (50) and C (100) to make 40 and 90.
	C can be placed before D (500) and M (1000) to make 400 and 900.
*/

func RomanToInt(s string) int {
	result := 0
	sLen := len(s)

	for i := 0; i < sLen; {
		var a, b string
		a = fmt.Sprintf("%c", s[i])

		if i+1 < sLen {
			b = fmt.Sprintf("%c", s[i+1])
		}

		if val, ok := romanToIntKnownMap[a+b]; ok {
			result += val
			i += 2
		} else {
			result += romanToIntKnownMap[a]
			i++
		}
	}

	return result
}
