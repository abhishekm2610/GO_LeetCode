package main

import "fmt"

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
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	ansHead := ListNode{head.Val, nil}
	curr := &ansHead
	for head.Next != nil {
		head = head.Next

		newHead := ListNode{head.Val, curr}

		curr = &newHead

	}
	newHead := ListNode{head.Val, curr.Next}
	fmt.Println(newHead, newHead.Next)
	return &newHead

}
func main() {
	L1N5 := ListNode{5, nil}
	L1N4 := ListNode{4, &L1N5}
	L1N3 := ListNode{3, &L1N4}
	L1N2 := ListNode{2, &L1N3}
	L1N1 := ListNode{1, &L1N2}
	head := reverseList(&L1N1)
	fmt.Println(head)
	for true {
		fmt.Println(head.Val)
		head = head.Next
		if head.Next == nil {
			break
		}
	}
}
