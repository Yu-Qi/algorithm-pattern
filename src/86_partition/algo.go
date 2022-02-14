package algo

import (
	. "../ListNode"
)
func merge(s_head, s_tail, b_head *ListNode) *ListNode{
	if s_head == nil{
		return b_head
	}else{
		s_tail.Next = b_head
		return s_head
	}
}

func partition(head *ListNode, x int) *ListNode {
	small_head := &ListNode{}
	big_head := &ListNode{}
	small_tail := small_head
	big_tail := big_head

	for head != nil{
		if head.Val < x{
			small_tail.Next = &ListNode{head.Val,nil}
			small_tail = small_tail.Next 
		} else{
			big_tail.Next = &ListNode{head.Val,nil}
			big_tail = big_tail.Next 
		}
		head = head.Next
	}
	return merge(small_head.Next, small_tail, big_head.Next)
}