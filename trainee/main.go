package main

import "fmt"

func main() {
	var stack = &Stack{}
	var total int
	fmt.Scan(&total)
	for i := 0; i < total; i++ {
		var count int
		fmt.Scan(&count)
		for j := 0; j < count; j++ {
			var val float64
			fmt.Scan(&val)
			stack.Push(val)
			
		}
	}
}

type Stack struct {
	top  *Node
	size int
}

type Node struct {
	prev *Node
	val  float64
}

func (s *Stack) Pop() float64 {
	temp := s.top.val
	s.top = s.top.prev
	s.size--
	return temp
}

func (s *Stack) Push(val float64) {
	newNode := &Node{val: val, prev: s.top}
	s.top = newNode
	s.size++
}

func (s *Stack) Back() float64 {
	return s.top.val
}
