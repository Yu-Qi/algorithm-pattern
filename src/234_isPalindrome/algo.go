package algo
import (
	. "../ListNode"
)
func findMiddle(head *ListNode) *ListNode{
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
func reverseList(head *ListNode) *ListNode{
	var res *ListNode
	for head != nil{
		res = &ListNode{head.Val, res}
		head = head.Next
	}
	return res
}
func isPalindrome(head *ListNode) bool {
    middle := findMiddle(head)
	tail := middle.Next
	middle.Next = nil

	right := reverseList(tail)
	for right != nil{
		if head.Val != right.Val{
			return false
		}
		head = head.Next
		right = right.Next
	}
	return true


}
