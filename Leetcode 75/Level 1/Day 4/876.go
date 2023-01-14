package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	count := 0
	ans := head
	for head != nil {
		count++
		if count%2 == 0 {
			ans = ans.Next
		}
		head = head.Next
	}
	return ans
}

func main() {

	list := []int{1}
	var headPointer *ListNode
	var head *ListNode

	for i := range list {
		node := ListNode{list[i], nil}

		if headPointer == nil {
			headPointer = &node
			head = headPointer
		} else {
			// node.Next = headPointer.Next
			headPointer.Next = &node
			headPointer = headPointer.Next
		}
	}

	ans := middleNode(head)

	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}

}
