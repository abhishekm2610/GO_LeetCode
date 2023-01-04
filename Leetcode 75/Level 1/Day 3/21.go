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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	head := ListNode{}
	if list1.Val <= list2.Val {
		head.Val = list1.Val
		list1 = list1.Next

	} else {
		head.Val = list2.Val
		list2 = list2.Next

	}
	head.Next = nil
	curr := &head
	for true {
		newNode := ListNode{}

		if list1 == nil {
			newNode = *list2
			curr.Next = &newNode
			curr = &newNode
			break

		} else if list2 == nil {
			newNode = *list1
			curr.Next = &newNode
			curr = &newNode
			break
		}
		if list1.Val <= list2.Val {
			newNode.Val = list1.Val
			list1 = list1.Next

		} else {
			newNode.Val = list2.Val
			list2 = list2.Next

		}
		newNode.Next = nil
		curr.Next = &newNode
		curr = &newNode

	}
	return &head
}
func main() {
	// L1N3 := ListNode{4, nil}
	// L1N2 := ListNode{2, &L1N3}
	// L1N1 := ListNode{1, &L1N2}
	// L2N3 := ListNode{4, nil}
	// L2N2 := ListNode{3, &L2N3}
	// L2N1 := ListNode{1, &L2N2}
	head := mergeTwoLists(nil, nil)
	// head := &L2N1
	fmt.Println(head)
	// for true {
	// 	fmt.Println(head.Val)
	// 	head = head.Next
	// 	if head.Next == nil {
	// 		break
	// 	}
	// }
}
