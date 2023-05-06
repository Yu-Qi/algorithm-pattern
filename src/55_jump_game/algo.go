package algo

// "fmt"

func canJump(nums []int) bool {
	max_idx := 0

	for i := 0; i < len(nums); {
		num := nums[i]
		max := 0
		max_idx = 0
		if num == 0 {
			return i == 0 && len(nums) == 1
		}
		for j := 1; j <= num; j++ {
			if i+j >= len(nums)-1 {
				return true
			}
			if nums[i+j]+(i+j) >= max {
				max = nums[i+j] + (i + j)
				max_idx = i + j
			}
			// fmt.Printf("i=%d,j=%d,nums[i+j]=%d,max=%d,max_idx=%d\n",i,j,nums[i+j],max,max_idx)
		}
		if max_idx == i {
			return false
		}
		i = max_idx
	}
	return max_idx+1 >= len(nums)
}
