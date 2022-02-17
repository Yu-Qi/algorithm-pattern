package algo

func singleNumber_v0(nums []int) int {
	// 不適用於負數
	a, b :=0,0
	for i:=0;i<len(nums);i++{
		if a^nums[i] >=a || (a^nums[i]<=a && b&nums[i]<=b){
			a = a^nums[i]
		} else if a^nums[i]<=a && b&nums[i]>b{
			b = b^nums[i]
		}
	}
	return a
}

func singleNumber(nums []int) int {
    result := 0
    for i:=0;i<64;i++{
        sum :=0
        for j:=0;j<len(nums);j++{
            sum += nums[j] >>i & 1
        }
        result |= (sum%3) << i
    }
    return result
}