package main

import (
	"fmt"

	"github.com/abhishekgautam2808/stackprogram/stack"
)

func main() {

	var s stack.Stack

	s = s.Push(1)
	s = s.Push(2)
	s = s.Push(3)
	s = s.Pop()
	fmt.Println(s)
}
