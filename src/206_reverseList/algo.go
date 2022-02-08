package algo

import (
	// "container/list"
	// "fmt"

	. "../ListNode"
)

func reverseListv0(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	hashmap := map[int]*ListNode{}
	idx := 0
	// queue := list.New()
	for head != nil {
		hashmap[idx] = head
		idx++
		// queue.PushBack(head)
		head = head.Next
	}
	idx--
	res := &ListNode{hashmap[idx].Val, nil}
	cur := res
	idx--
	for ; idx >= 0; idx-- {
		cur.Next = &ListNode{hashmap[idx].Val, nil}
		cur = cur.Next
	}
	return res
}

func reverseList(head *ListNode) *ListNode {
	var res *ListNode
	for head != nil {
		res = &ListNode{head.Val, res}
		head = head.Next
	}
	return res
}
