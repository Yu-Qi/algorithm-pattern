package ListNode

type ListNode struct {
     Val int
     Next *ListNode
 }

func ArrToListNode(nums []int) *ListNode{
	if len(nums)==0{
		return nil
	}
	head := ListNode{nums[0], nil}
	cur_node := head
	for i:=1; i<len(nums);i++{
		new_node := ListNode{nums[i],nil}
		cur_node.Next = &new_node
		cur_node = new_node
	}
	return &head
}