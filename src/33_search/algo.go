package algo

func search(nums []int, target int) int {
	start :=0
	end := len(nums)-1
	for start+1<end{
		middle := (start+end)/2
		if nums[middle] == target{
			return middle
		} else if nums[end] >= target && nums[middle] < target{
			start = middle
		} else{
			end = middle
		}
	}
	if nums[start] == target{
		return start
	} else if nums[end] == target{
		return end
	}
	return -1
}