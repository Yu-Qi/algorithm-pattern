package algo

import (
	. "../ListNode"
)
func bubblesort_v0(head *ListNode)*ListNode{
	// Time Limit Exceeded -> {2,3,4,5,...,50000,1}
	sorted := false
	for sorted == false{
		sorted = true
		cur := head
		for cur != nil && cur.Next != nil{
			if cur.Val > cur.Next.Val{
				tmp := cur.Val
				cur.Val = cur.Next.Val
				cur.Next.Val = tmp
				sorted = false
			}
			cur = cur.Next
		}
	}
	return head
}
func findMiddle(head *ListNode) *ListNode{
	slow := head
	fast := head.Next
	for fast !=nil && fast.Next!=nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

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

func mergesort(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
        return head
    }
	middle := findMiddle(head)
	tail := middle.Next
	middle.Next = nil //切斷兩邊
	left := mergesort(head)
	right := mergesort(tail)
	return mergeTwoLists(left, right)
}

func sortList(head *ListNode) *ListNode {
    return mergesort(head)
}