package main

func fifteenth() {
	var total int
	var stack = &Stack{}
	fmt.Scan(&total)
	res := make([]int, total)
	for i := 0; i < total; i++ {
		var temp int
		fmt.Scan(&temp)
		if i != 0 {
			val, index := stack.Back()
			for val > temp {
				stack.Pop()
				res[index] = i
				if stack.size > 0 {
					val, index = stack.Back()
				} else {
					break
				}
			}
		}
		stack.Push(temp, i)
	}
	for stack.size != 0 {
		_, index := stack.Pop()
		res[index] = -1
	}
	for _, val := range res {
		fmt.Printf("%d ", val)
	}
}

type Stack struct {
	top  *Node
	size int
}
type Node struct {
	prev  *Node
	val   int
	index int
}

func (s *Stack) Push(val, index int) {
	newNode := &Node{val: val, index: index, prev: s.top}
	s.top = newNode
	s.size++
}

func (s *Stack) Pop() (val, index int) {
	resVal := s.top.val
	resIndex := s.top.index
	s.top = s.top.prev
	s.size--
	return resVal, resIndex
}

func (s *Stack) Back() (val, index int) {
	return s.top.val, s.top.index
}
