package algo
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}
	if p.Val == root.Val || q.Val == root.Val{
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil && right != nil{
		return right
	} else if right == nil && left != nil{
		return left
    } else if left == nil && right == nil{
        return nil
    } else if left.Val != right.Val{
		return root
	}
	return nil
}