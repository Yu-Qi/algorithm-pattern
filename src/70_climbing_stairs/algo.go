package algo

// "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	mem := []int{1, 2}
	for i := 2; i < n; i++ {
		mem = append(mem, mem[i-1]+mem[i-2])
	}
	return mem[n-1]
}
