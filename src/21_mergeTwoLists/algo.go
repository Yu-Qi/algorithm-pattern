package algo

import (
	// "fmt"

	. "../ListNode"
)
func append(mylist *ListNode, val int) *ListNode{
	new_node := &ListNode{val,nil}
	mylist.Next = new_node
	mylist = new_node	
	return mylist
}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    res := &ListNode{}
	cur := res
	for list1 != nil && list2!=nil{
		if list1.Val <= list2.Val{
			cur = append(cur, list1.Val)
			list1 = list1.Next
		} else{
			cur = append(cur, list2.Val)
			list2 = list2.Next
		}
	}
	for list1 != nil{
		cur = append(cur, list1.Val)
		list1 = list1.Next
	}
	for list2 != nil{
		cur = append(cur, list2.Val)
		list2 = list2.Next
	}
	return res.Next
}