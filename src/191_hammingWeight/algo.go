package algo

func hammingWeight(num uint32) int {
	var count uint32
	for i:=0;i<32;i++{
        if (num >> 1) & 1 == 1{
        }
        count += num & 1
        num = num>>1
	}
    return int(count)
}