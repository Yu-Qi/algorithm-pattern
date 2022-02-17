package algo

func singleNumber(nums []int) int {
	bit := 0
	for _,num := range nums{
		bit ^=num
	}
	return bit
}