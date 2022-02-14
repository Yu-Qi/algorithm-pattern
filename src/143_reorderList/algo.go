package algo

import (
	. "../ListNode"
	// "fmt"
)
func findMiddle(head *ListNode)*ListNode{
	slow := head
	fast := head.Next
	for fast != nil && fast.Next !=nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	fmt.Printf("slow:%d\n",slow.Val)
	return slow
}
func splitTwoList(head *ListNode)(*ListNode,*ListNode){
	middle := findMiddle(head)
	tail := middle.Next
	middle.Next = nil
	return head, tail
}
func reverseList(head *ListNode)*ListNode{
	var res *ListNode
	for head != nil{
		res = &ListNode{head.Val, res}
		head = head.Next
	}
	return res
}
func mergeTwoList(l1, l2 *ListNode){
	for l2 != nil{
		tmp := l2
		l2 = l2.Next
		tmp.Next = l1.Next
		l1.Next = tmp
		l1 = l1.Next.Next
	}
}
func reorderList(head *ListNode){
	if head == nil || head.Next ==nil{
		return
	}
    l1, l2 := splitTwoList(head)
	l2_reverse := reverseList(l2)
	// fmt.Printf("l1:%v,:2:%v\n",ArrFromListNode(l1),ArrFromListNode(l2_reverse))
	mergeTwoList(l1,l2_reverse)
}