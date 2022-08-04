package addtwonumbers

import (
	"fmt"
	"math"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// kinda hacky, but blame go, not me
func (ln *ListNode) SetLastToNil() {
	len := ln.Len() - 1
	temp := ln
	for i := 0; i < len; i++ {
		temp = temp.Next
	}

	temp.Next = nil
}

func (ln *ListNode) Insert(val int) {
	temp := ln
	for temp.Next != nil {
		temp = temp.Next
	}

	temp.Val = val
	temp.Next = &ListNode{}
}

func (ln *ListNode) InsertMany(vals ...int) {
	for _, v := range vals {
		ln.Insert(v)
	}
}

func (ln *ListNode) InsertFromSlice(vals []int) {
	for _, v := range vals {
		ln.Insert(v)
	}
}

func (ln *ListNode) SplitAndInsert(val int) {
	digits := digitsOfReversed(val)
	for _, d := range digits {
		ln.Insert(d)
	}
}

func (ln *ListNode) Traverse() []int {
	ln.SetLastToNil()
	vals := []int{}
	temp := ln
	for temp != nil {
		vals = append(vals, temp.Val)
		temp = temp.Next
	}

	fmt.Println("Linked List:", vals)
	return vals
}

func (ln *ListNode) Len() (len int) {
	for temp := ln; temp != nil; temp = temp.Next {
		if temp.Next != nil {
			len++
		}
	}
	return
}

func numberOfDigits(n int) (digits int) {
	for ; n != 0; n /= 10 {
		digits++
	}
	return
}

func digitsOf(n int) []int {
	numOfDigits := numberOfDigits(n)
	digits := make([]int, numOfDigits)

	for d, i := numOfDigits, 0; d != 0; d, i = d-1, i+1 {
		powOf := int(math.Pow10(d - 1))
		digit := n / powOf
		digits[i] = digit
		n -= digit * powOf
	}

	return digits
}

func digitsOfReversed(n int) []int {
	numOfDigits := numberOfDigits(n)
	digits := make([]int, numOfDigits)

	for d := numOfDigits; d != 0; d-- {
		powOf := int(math.Pow10(d - 1))
		digit := n / powOf
		digits[d-1] = digit
		n -= digit * powOf
	}

	return digits
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		longest     *ListNode
		shortest    *ListNode
		longestLen  int
		shortestLen int
	)

	ll := &ListNode{}

	sumSingleDigit := func(a, b, r int) (rn int) {
		partial := a + b + r
		if partial > 9 {
			rn = 1
			ll.Insert(partial - 10)
		} else {
			rn = 0
			ll.Insert(partial)
		}
		return
	}

	l1Len := l1.Len()
	l2Len := l2.Len()

	if l1Len >= l2Len {
		longest = l1
		shortest = l2
		longestLen = l1Len
		shortestLen = l2Len
	} else {
		longest = l2
		shortest = l1
		longestLen = l2Len
		shortestLen = l1Len
	}

	for i, k, carry := 0, 0, 0; i <= longestLen; i, k = i+1, k+1 {
		if i < longestLen {
			if k < shortestLen {
				carry = sumSingleDigit(longest.Val, shortest.Val, carry)
				longest = longest.Next
				shortest = shortest.Next
			} else {
				carry = sumSingleDigit(longest.Val, 0, carry)
				longest = longest.Next
			}
		} else {
			// maybe add last carry if not 0
			if carry == 1 {
				ll.Insert(carry)
			}
		}
	}

	return ll
}
