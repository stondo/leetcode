package checkOnesSegment

import (
	"testing"
)

// TestCheckOnesSegment1001 calls checkOnesSegment.checkOnesSegment with a string, checking
// if it contains at most one contiguous segment of ones.
func TestCheckOnesSegment1001(t *testing.T) {
	s := "1001"
	b := checkOnesSegment(s)
	if b {
		t.Fatalf(`checkOnesSegment(1001) = %v, want false`, b)
	}
}

// TestCheckOnesSegment110 calls checkOnesSegment.checkOnesSegment with a string, checking
// if it contains at most one contiguous segment of ones.
func TestCheckOnesSegment110(t *testing.T) {
	s := "110"
	b := checkOnesSegment(s)
	if !b {
		t.Fatalf(`checkOnesSegment(110) = %v, want true`, b)
	}
}
