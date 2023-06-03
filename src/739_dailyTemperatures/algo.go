package algo

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := [][]int{}

	for i := len(temperatures) - 1; i >= 0; i-- {
		if len(stack) == 0 || temperatures[i] >= stack[len(stack)-1][0] {
			for len(stack) > 0 && temperatures[i] >= stack[len(stack)-1][0] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				result[i] = 0
			} else {
				result[i] = stack[len(stack)-1][1] - i
			}
		} else {
			result[i] = stack[len(stack)-1][1] - i
		}

		stack = append(stack, []int{temperatures[i], i})
	}
	return result
}
