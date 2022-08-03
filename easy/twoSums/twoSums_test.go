package twosums

import (
	"testing"
)

func intSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// TestCheckOnesSegment1 calls twosums.twoSum with a slice of integer and a target, checking
// if wihch number at index adds up to target.
func TestCheckOnesSegment1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	if !intSlicesEqual(res, []int{0, 1}) {
		t.Fatalf(`checkOnesSegment(1001) = %v, want false`, res)
	}
}

// TestCheckOnesSegment2 calls twosums.twoSum with a slice of integer and a target, checking
// if wihch number at index adds up to target.
func TestCheckOnesSegment2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	res := twoSum(nums, target)
	if !intSlicesEqual(res, []int{1, 2}) {
		t.Fatalf(`checkOnesSegment(1001) = %v, want false`, res)
	}
}

// TestCheckOnesSegment1 calls twosums.twoSum with a slice of integer and a target, checking
// if wihch number at index adds up to target.
func TestCheckOnesSegment3(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	res := twoSum(nums, target)
	if !intSlicesEqual(res, []int{0, 1}) {
		t.Fatalf(`checkOnesSegment(1001) = %v, want false`, res)
	}
}
