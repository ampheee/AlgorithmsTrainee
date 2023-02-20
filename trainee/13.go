package main

import (
	"bufio"
	"fmt"
	"os"
	"text/scanner"
)

type Stack struct {
	top  *Node
	size int
}

type Node struct {
	prev *Node
	val  int
}

func thirteenth() {
	var spaceCount int
	var stack = &Stack{}
	f, _ := os.Open("input.txt")
	myScanner := bufio.NewReader(f)
	var r rune
	for spaceCount <= 2 && r != 10 && r != scanner.EOF {
		r, _, _ = myScanner.ReadRune()
		if r == ' ' {
			spaceCount++
			continue
		} else {
			spaceCount = 0
		}
		if isDigit(r) {
			stack.Push(int(r - 48))
		} else if isOperator(r) {
			val1 := stack.Pop()
			val2 := stack.Pop()
			if r == '*' {
				stack.Push(val1 * val2)
			} else if r == '+' {
				stack.Push(val1 + val2)
			} else if r == '-' {
				stack.Push(val2 - val1)
			}
		}
	}
	if stack.size == 1 {
		fmt.Println(stack.Pop())
	}
}

func (s *Stack) Pop() int {
	val := s.top.val
	s.top = s.top.prev
	s.size--
	return val
}

func (s *Stack) Push(r int) {
	var node = &Node{val: r, prev: s.top}
	s.top = node
	s.size++
}

func isDigit(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func isOperator(r rune) bool {
	if r == '*' || r == '-' || r == '+' {
		return true
	}
	return false
}
