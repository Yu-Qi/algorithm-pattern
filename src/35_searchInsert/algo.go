package algo

func searchInsert(nums []int, target int) int {
	start:=0
	end := len(nums)-1
	for start+1 < end{
		middle := (end+start)/2
		if nums[middle] == target{
			return middle
		} else if nums[middle] < target{
			start = middle+1
		} else {
			end = middle-1
		}
	}
	if nums[start] >=target{
		return start
	} else if nums[end] < target{
		return end+1
	} else{
		return start+1
	}
}