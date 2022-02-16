package algo

func largestRectangleArea_v0(heights []int) int {
	// 當是一個超級大的正長方形，會Time Limit Exceeded
	area := 0
	for i:=0;i<len(heights);i++{
		v := dfs(&heights, i, heights[i])
		if v >area{
			area = v
		}
	}
	return area
}
func dfs(heights *[]int, idx, target int) int{
	// left
	left_idx := idx-1
	left := 0
	for left_idx >=0 && (*heights)[left_idx] >= target{
		left += target
		left_idx --
	}
	right_idx := idx+1
	right := 0
	for right_idx < len(*heights) && (*heights)[right_idx] >= target{
		right += target
		right_idx++
	}
	return left+right+target
}

func largestRectangleArea(heights []int) int {
	// 當是一個超級大的正長方形，會Twime Limit Exceeded
	area := 0
	for i:=0;i<len(heights);i++{
		v := dfs(&heights, i, heights[i])
		if v >area{
			area = v
		}
	}
	return area
}