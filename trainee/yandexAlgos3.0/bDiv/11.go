package main

//type Stack struct {
//	root *Node
//	size int
//}
//
//type Node struct {
//	Val  int
//	Next *Node
//}

//func (s *Stack) Push(val int) {
//	if s.root == nil {
//		var n = &Node{val, nil}
//		s.root = n
//		s.size++
//		return
//	}
//	temp := s.root
//	for temp.Next != nil {
//		temp = temp.Next
//	}
//	var node = &Node{val, nil}
//	temp.Next = node
//	s.size++
//}

//func (s *Stack) Pop() {
//	temp := s.root
//	if s.root == nil {
//		fmt.Println("error")
//		return
//	}
//	if s.size == 1 {
//		fmt.Println(s.root.Val)
//		s.root = nil
//		s.size--
//		return
//	}
//	for {
//		if temp.Next != nil && temp.Next.Next == nil {
//			break
//		}
//		temp = temp.Next
//	}
//	res := temp.Next.Val
//	temp.Next = nil
//	s.size--
//	fmt.Println(res)
//}

//func (s *Stack) Clear() {
//	for s.size != 0 {
//		var temp = s.root
//		for {
//			if temp.Next != nil && temp.Next.Next == nil {
//				break
//			}
//			temp.Next = nil
//			break
//		}
//		s.size--
//	}
//	s.root = nil
//	fmt.Println("ok")
//}

//func (s *Stack) Back() {
//	if s.root == nil {
//		fmt.Println("error")
//		return
//	}
//	temp := s.root
//	for temp.Next != nil {
//		temp = temp.Next
//	}
//	fmt.Println(temp.Val)
//}

//func eleven() {
//	var stack = &Stack{}
//	for {
//		var choice string
//		_, _ = fmt.Scan(&choice)
//		switch choice {
//		case "push":
//			var temp int
//			_, _ = fmt.Scan(&temp)
//			stack.Push(temp)
//			fmt.Println("ok")
//		case "back":
//			stack.Back()
//		case "pop":
//			stack.Pop()
//		case "size":
//			fmt.Println(stack.size)
//		case "clear":
//			stack.Clear()
//		case "exit":
//			fmt.Println("bye")
//			return
//		}
//	}
//}
