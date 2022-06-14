package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverFromPreorder(traversal string) *TreeNode {
	if len(traversal) == 1 {
		intVal, _ := strconv.Atoi(traversal)
		root := TreeNode{Val: intVal, Left: nil, Right: nil}
		return &root
	}
	var nodesVisited = make(map[int]*TreeNode)

	for i, v := range traversal {
		if v != '-' {
			node := TreeNode{}
			intVal, _ := strconv.Atoi(string(v))
			node.Val = intVal
			j := i
			count := 0
			for j < len(traversal) {

				if traversal[j] != '-' {
					leftnode := TreeNode{}
					leftnode.Val, _ = strconv.Atoi(string(traversal[j]))

				} else {
					count++
				}
				j++
			}
		}

	}
}

func main() {
	input := "1-2--3--4-5--6--7"

	fmt.Println(recoverFromPreorder((input)))
}
