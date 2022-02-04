package algo

import (
	. "../Tree"
	. "../inorderTraversal"
)
func isSorted(nums []int) bool{
	for i:=1;i<len(nums);i++{
		if nums[i-1] >= nums[i]{
			return false
		}
	}
	return true
}
func isValidBST(root *TreeNode) bool {
	ordered_traversal := InorderTraversal(root)
	return isSorted(ordered_traversal)
}