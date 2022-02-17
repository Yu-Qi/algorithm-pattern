package algo

func search(nums []int, target int) int {
	middle := len(nums)/2
    if len(nums)==0{
        return -1
    }
	if nums[middle] ==target{
		return middle
	} else if nums[middle]>target{
        idx:= search(nums[:middle], target)
        if idx == -1{
            return -1
        } else{
            return idx
        }
	} else{
        idx:=search(nums[middle+1:],target)
        if idx == -1{
            return -1
        } else{
            return idx+middle+1
        }        
	}
}