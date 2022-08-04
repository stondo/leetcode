package addtwonumbers

import (
	"testing"
)

func slicesEqual(a, b []int) bool {
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

// TestaddTwoNumbers1 calls addtwonumbers.addTwoNumbers with 2 Linked List of type int, checking
// for a valid sum value.
func TestAddTwoNumbers1(t *testing.T) {
	a := 342
	b := 465

	l1 := &ListNode{}
	l1.SplitAndInsert(a)

	l2 := &ListNode{}
	l2.SplitAndInsert(b)

	expected := []int{7, 0, 8}
	ll := addTwoNumbers(l1, l2)

	if !slicesEqual(ll.Traverse(), expected) {
		t.Fatalf(`addTwoNumbers(l1, l2) = %v, want []int{7, 0, 8}`, ll.Traverse())
	}
}

// TestaddTwoNumbers2 calls addtwonumbers.addTwoNumbers with 2 Linked List of type int, checking
// for a valid sum value.
func TestAddTwoNumbers2(t *testing.T) {
	a := 0
	b := 0

	l1 := &ListNode{}
	l1.SplitAndInsert(a)

	l2 := &ListNode{}
	l2.SplitAndInsert(b)

	expected := []int{0}
	ll := addTwoNumbers(l1, l2)

	if !slicesEqual(ll.Traverse(), expected) {
		t.Fatalf(`addTwoNumbers(l1, l2) = %v, want []int{0}`, ll.Traverse())
	}
}

// TestaddTwoNumbers3 calls addtwonumbers.addTwoNumbers with 2 Linked List of type int, checking
// for a valid sum value.
func TestAddTwoNumbers3(t *testing.T) {
	a := 9999999
	b := 9999

	l1 := &ListNode{}
	l1.SplitAndInsert(a)

	l2 := &ListNode{}
	l2.SplitAndInsert(b)

	expected := []int{8, 9, 9, 9, 0, 0, 0, 1}
	ll := addTwoNumbers(l1, l2)

	if !slicesEqual(ll.Traverse(), expected) {
		t.Fatalf(`addTwoNumbers(l1, l2) = %v, want []int{8, 9, 9, 9, 0, 0, 0, 1}`, ll.Traverse())
	}
}
