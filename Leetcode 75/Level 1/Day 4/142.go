package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	addresses := make(map[*ListNode]int)

	for head != nil {
		if _, exists := addresses[head]; exists {
			return head
		} else {
			addresses[head] = 1
		}
		head = head.Next
	}
	return head
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

	ans := detectCycle(head)

	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}

}
