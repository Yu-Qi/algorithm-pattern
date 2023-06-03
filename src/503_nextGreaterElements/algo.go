package algo

func nextGreaterElements(nums []int) []int {
	stack := []int{}
	result := make([]int, len(nums))
	size := len(nums)
	for i := 2*size - 1; i >= 0; i-- {

		if len(stack) == 0 || nums[i%size] >= stack[len(stack)-1] {
			for len(stack) > 0 && nums[i%size] >= stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				result[i%size] = -1
			} else {
				result[i%size] = stack[len(stack)-1]
			}
		} else {
			result[i%size] = stack[len(stack)-1]
		}

		if len(stack) == 0 || nums[i%size] < stack[len(stack)-1] {
			stack = append(stack, nums[i%size])
		} else {
			for len(stack) > 0 && nums[i%size] >= stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, nums[i%size])
		}

	}

	return result
}
