package algo

import (
	// "fmt"

	. "../ListNode"
)

func reverse(nodes *[]*ListNode, left int, right int) {
	if left == -1 && right == -1 {
		return
	}
	size := -1
	if right-left == 1 {
		size = 1
	} else {
		size = (right - left) / 2
	}
	for i := 0; i < size; i++ {
		tmp := (*nodes)[i+left-1]
		(*nodes)[i+left-1] = (*nodes)[right-i-1]
		(*nodes)[right-i-1] = tmp
	}
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	nodes := []*ListNode{}
	idx := 0

	for head != nil {
		nodes = append(nodes, head)
		idx++
		head = head.Next
	}
	reverse(&nodes, left, right)
	var res *ListNode
	for i := len(nodes) - 1; i >= 0; i-- {
		res = &ListNode{nodes[i].Val, res}
	}
	return res
}
