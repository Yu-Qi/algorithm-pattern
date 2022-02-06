package algo
import (
		. "../ListNode"
	)
func deleteDuplicates(head *ListNode) *ListNode {
	node := head
	for node != nil && node.Next != nil{
		if node.Val == node.Next.Val{
			node.Next = node.Next.Next
		}else{
			node = node.Next
		}
	}
	return head
}