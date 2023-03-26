package leetcode

//Given the head of a singly linked list, reverse the list, and return the reversed list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	curr, prev := head, head
	for curr.Next != nil {
		curr = curr.Next
		curr.Next = prev
		prev = curr
	}
	head = prev
	return head
}
