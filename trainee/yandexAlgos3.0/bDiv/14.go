package main

//type Stack struct {
//	top  *Node
//	size int
//}
//
//type Node struct {
//	prev *Node
//	val  int
//}
//
//func (s *Stack) Push(val int) {
//	var newNode = &Node{val: val, prev: s.top}
//	s.top = newNode
//	s.size++
//}
//
//func (s *Stack) Back() int {
//	var res = s.top.val
//	return res
//}
//
//func (s *Stack) Pop() int {
//	var res = s.top.val
//	s.top = s.top.prev
//	s.size--
//	return res
//}
//
//func fourteenth() {
//	var total int
//	var stack = &Stack{}
//	fmt.Scan(&total)
//	var count = 1
//	for i := 0; i < total; i++ {
//		var temp int
//		fmt.Scan(&temp)
//		stack.Push(temp)
//		for stack.size != 0 {
//			if stack.Back() != count {
//				break
//			}
//			stack.Pop()
//			count++
//		}
//	}
//	if stack.size == 0 {
//		fmt.Println("YES")
//	} else {
//		fmt.Println("NO")
//	}
//}
