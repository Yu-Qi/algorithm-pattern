package algo

import (
	. "../ListNode"
)
func hasCycle(head *ListNode) bool {
	if head == nil{
		return false
	}
	slow := head 
	fast := head.Next
	for (fast != nil && fast.Next !=nil) && (slow!=fast){
		fast = fast.Next.Next
		slow = slow.Next 
	}
	return slow ==fast
}