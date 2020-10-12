package stack

import (
	"fmt"
	"testing"
)

func compare(a, b []int) bool {
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
func TestPush(t *testing.T) {
	var s Stack
	var expected = []int{10}
	res := s.Push(10)
	if compare(res, expected) == false {
		t.Fatalf("expected result to be %v, got %v", expected, res)
	}
}

func TestPushMax(t *testing.T) {
	var s Stack
	var expected = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 11; i++ {
		s = s.Push(i)
	}

	if compare(s, expected) == false {
		fmt.Println(s)
		t.Fatalf("Maximum limit of stack exceeded")
	}
}

func TestPop(t *testing.T) {
	var s Stack
	var expected = []int{0, 1, 2, 3}
	for i := 0; i < 5; i++ {
		s = s.Push(i)
	}
	s = s.Pop()
	if compare(s, expected) == false {
		fmt.Println(s)
		t.Fatalf("expected result to be %v, got %v", expected, s)
	}
}

func TestPopLengthZero(t *testing.T) {
	var s Stack
	var expected = []int{}
	s = s.Pop()
	if compare(s, expected) != false {
		fmt.Println("no element available to pop")
	}
}
