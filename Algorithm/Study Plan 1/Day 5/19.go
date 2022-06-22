package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := 0
	node := head
	previousNode := head
	removedNode := head
	for node.Next != nil {
		fast++
		if fast == n-1 {
			previousNode = node
			removedNode = previousNode.Next
		}
		node = node.Next
	}
	previousNode.Next = removedNode.Next
	return head
}

func main() {
	elm5 := ListNode{Val: 5, Next: nil}
	elm4 := ListNode{Val: 4, Next: &elm5}
	elm3 := ListNode{Val: 3, Next: &elm4}
	elm2 := ListNode{Val: 2, Next: &elm3}

	elm1 := ListNode{Val: 1, Next: &elm2}
	ip := 2
	fmt.Println(removeNthFromEnd(&elm1, ip))
}
