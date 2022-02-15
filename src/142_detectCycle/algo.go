package algo


import (
	. "../ListNode"
)

func hasCycle(head *ListNode) (bool,*ListNode){
	if head == nil{
		return false,nil
	}
	slow := head 
	fast := head.Next
	for (fast != nil && fast.Next !=nil) && (slow!=fast){
		fast = fast.Next.Next
		slow = slow.Next 
	}
	return slow ==fast,fast
}

func detectCycle(head *ListNode) *ListNode {
	isCycle,node := hasCycle(head)
	if !isCycle{
		return nil
	}else{
		// 算算數學
		slow := head
		fast := node.Next //不知道為啥要從下一個點開始
		for slow != fast{
			slow = slow.Next
			fast = fast.Next
		}
		return slow
	}
	return nil
}