package integertoroman

import (
	"log"
	"math"
	"strings"
)

var intToRomanKnownMap = map[int]string{
	1:    "I",
	2:    "II",
	3:    "III",
	4:    "IV",
	5:    "V",
	6:    "VI",
	7:    "VII",
	8:    "VIII",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func numberOfDigits(n int) (digits int) {
	for ; n != 0; n /= 10 {
		digits++
	}
	return
}

func IntToRoman(num int) string {
	result := ""
	digits := numberOfDigits(num)
	log.Println("digits", digits)

	for d := digits; d != 0; d-- {
		log.Println("d", d)

		powOf := int(math.Pow10(d - 1))

		splitted := (num / powOf)
		log.Println("splitted", splitted)

		if splitted != 0 {
			if maybeRoman, ok := intToRomanKnownMap[splitted*powOf]; ok {
				log.Println("maybeRoman", maybeRoman)
				result += maybeRoman
			} else {
				if roman, ok := intToRomanKnownMap[powOf]; ok {
					log.Println("roman", roman)
					if splitted > 5 {
						result += intToRomanKnownMap[5*powOf] + strings.Repeat(roman, splitted-5)
					} else {
						result += strings.Repeat(roman, splitted)
					}
				}
			}
		}

		num -= splitted * powOf
		log.Println("num", num)
	}

	return result
}
