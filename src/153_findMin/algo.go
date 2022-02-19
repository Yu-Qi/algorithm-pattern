package algo

func findMin(nums []int) int {
	start:=0
	end:= len(nums)-1
	for start+1<end{
		middle := (start+end)/2
		if nums[middle] >= nums[end]{
			start = middle
        } else if nums[middle] >= nums[start]{
			end = middle
        } else{
            end = middle
        }
	}
    if nums[start]<nums[end]{
        return nums[start]
    }else{
        return nums[end]
    }
}