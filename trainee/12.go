package main

//func twelve() {
//	var str string
//	var stack = &Stack{}
//	fmt.Scan(&str)
//	for _, val := range str {
//		if val == '(' || val == '[' || val == '{' {
//			stack.Push(val)
//		} else if val == ')' || val == ']' || val == '}' {
//			back := stack.Back()
//			if val == ')' && back == '(' {
//				stack.Pop()
//			} else if val == ']' && back == '[' {
//				stack.Pop()
//			} else if val == '}' && back == '{' {
//				stack.Pop()
//			} else {
//				fmt.Println("no")
//				return
//			}
//		}
//	}
//	if stack.size != 0 {
//		fmt.Println("no")
//	} else {
//		fmt.Println("yes")
//	}
//}
//
//type Stack struct {
//	root *Node
//	size int
//}
//
//type Node struct {
//	Val  int32
//	Next *Node
//}
//
//func (s *Stack) Push(value int32) {
//	n := &Node{value, s.root}
//	s.root = n
//	s.size++
//}
//
//func (s *Stack) Back() (res int32) {
//	if s.size == 0 {
//		return res
//	}
//	return s.root.Val
//}
//func (s *Stack) Pop() {
//	if s.size == 0 {
//		return
//	}
//	n := s.root
//	s.root = n.Next
//	s.size--
//}
