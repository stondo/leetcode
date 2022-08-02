package checkOnesSegment

import (
	"fmt"
)

func checkOnesSegment(s string) bool {
	zeroFound := false
	oneAfterZero := false

	for i := 1; i < len(s); i++ {
		if fmt.Sprintf("%c", s[i]) == "0" {
			zeroFound = true
		} else if fmt.Sprintf("%c", s[i]) == "1" && zeroFound {
			oneAfterZero = true
		}
	}

	return !oneAfterZero
}
