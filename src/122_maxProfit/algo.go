package algo

func maxProfit(prices []int) int {

	return maxProfitMonotonicStack(prices)
}

func maxProfitMonotonicStack(prices []int) int {
	stack := [][]int{}
	profit := 0

	// prepare monotonic stack
	for i, val := range prices {
		if i == 0 {
			stack = append(stack, []int{i, val})
		} else {
			if val < stack[len(stack)-1][1] {
				stack[len(stack)-1] = []int{i, val}
			} else if val > stack[len(stack)-1][1] {
				stack = append(stack, []int{i, val})
			}
		}
	}
	if len(stack) <= 1 {
		return 0
	}

	for i := len(prices) - 1; i >= 0; i-- {
		if len(stack) == 0 {
			break
		}
		if i > stack[len(stack)-1][0] && prices[i] > stack[len(stack)-1][1] {
			profit += (prices[i] - stack[len(stack)-1][1])
		} else if i <= stack[len(stack)-1][0] {
			stack = stack[:len(stack)-1]
		}
	}
	return profit
}
