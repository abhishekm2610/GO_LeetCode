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

	fast := head
	// node := head
	// previousNode := head
	slow := head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return head
}

func main() {
	// elm5 := ListNode{Val: 5, Next: nil}
	// elm4 := ListNode{Val: 4, Next: &elm5}
	// elm3 := ListNode{Val: 3, Next: &elm4}
	// elm2 := ListNode{Val: 2, Next: &elm3}

	elm1 := ListNode{Val: 1, Next: nil}
	ip := 1
	fmt.Println(removeNthFromEnd(&elm1, ip))
	// node := removeNthFromEnd(&elm1, ip)
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(node.Val)
	// 	node = node.Next
	// }
}

// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	fast := 1
// 	node := head
// 	previousNode := head
// 	removedNode := head
// 	for true {
// 		if fast >= n-1{
// 			previousNode = node
// 			removedNode = previousNode.Next
// 		}
//         		fast++
// node = node.Next
//         if node.Next== nil{
//             previousNode = node
// 			removedNode = previousNode.Next
//             break
//             }

// 	}

// 	previousNode.Next = removedNode.Next
// 	return head
// }

// func removeNthFromEnd(head *ListNode, n int) *ListNode {

// 	fast := head
// 	// node := head
// 	// previousNode := head
// 	slow := head

// 	for i := 0; i < n; i++ {
// 		fast = fast.Next
// 	}

// 	if fast == nil {
// 		return head.Next
// 	}
// 	for fast.Next != nil {
// 		fast = fast.Next
// 		slow = slow.Next
// 	}

// 	slow.Next = slow.Next.Next
// 	return head
// }
