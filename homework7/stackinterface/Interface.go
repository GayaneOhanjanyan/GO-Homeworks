package stackinterface

import "fmt"

const n = 100

// ArrayStack structura
type ArrayStack struct {
	top   int
	array [n]int
}

// Stack interface
type Stack interface {
	Push(value int) bool
	Pop() (int, bool)
}

// Push element to stack
func (s *ArrayStack) Push(value int) bool {

	if s.top >= n-1 {
		fmt.Println("Stack overflow")
		return false
	}
	s.top++
	s.array[s.top] = value
	return true

}

//Pop element from stack
func (s *ArrayStack) Pop() (int, bool) {
	if s.top <= 0 {
		fmt.Println("Stack is empty")
		return 0, false
	}
	fmt.Println("Poped element is", s.array[s.top])
	s.top--

	return s.array[s.top], true

}

// PrintStack function prints stack's elements
func (s *ArrayStack) PrintStack() {
	if s.top > 0 {
		fmt.Println("Stack elements are ")
		for i := s.top; i >= 1; i-- {
			fmt.Println(s.array[i])
		}
	}
	fmt.Println("Stack is empty")
}
