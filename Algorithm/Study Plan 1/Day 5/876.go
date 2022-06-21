package main

import (
	"fmt"
	"math"
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
func middleNode(head *ListNode) *ListNode {

	nodeid := 0
	node := head
	for node.Next != nil {
		nodeid++
		node = node.Next
	}
	nodeid++
	fmt.Println(nodeid)

	mid := math.Ceil(float64(nodeid / 2))

	i := 0
	node = head
	for i <= int(mid) {
		if i == int(mid) {
			return node
		}
		node = node.Next
		i++
	}
	return node
}

func main() {
	elm5 := ListNode{Val: 5, Next: nil}
	elm4 := ListNode{Val: 4, Next: &elm5}
	elm3 := ListNode{Val: 3, Next: &elm4}
	elm2 := ListNode{Val: 2, Next: &elm3}

	elm1 := ListNode{Val: 1, Next: &elm2}

	fmt.Println(middleNode(&elm1))
}
