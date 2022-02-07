package algo

import (
	// "fmt"
	. "../ListNode"
)

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}
    hashmap := map[int]int{}
    var entry int
	head2 := head
	for head != nil{
		entry = head.Val
		_, isExist := hashmap[entry]
		if isExist{
			hashmap[entry] +=1
		}else{
			hashmap[entry] = 1 
		}
        head = head.Next
	}
	var res,cur *ListNode
	for head2 != nil{
		if value, _ := hashmap[head2.Val]; value ==1{
			if res == nil{
				res = &ListNode{head2.Val, nil}
				cur = res
			} else{
				new_node := &ListNode{head2.Val, nil}
				cur.Next = new_node
				cur = new_node
			}
		}
		head2 = head2.Next
	}
	return res
}