package algo

func jump(nums []int) int {
	count := 0
	i := 0
	if len(nums) == 1 {
		return 0
	}

	for count < len(nums) {
		max_value := 0
		max_index := 0
		if i+nums[i] >= len(nums)-1 {
			return count + 1
		}
		for j := 1; j <= nums[i]; j++ {
			value := nums[i+j] + (i + j)
			if value >= max_value {
				max_value = value
				max_index = i + j
			}
		}
		count++
		i = max_index
		if max_index >= len(nums)-1 {
			return count
		}

	}
	return count
}
