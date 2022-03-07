package algo
func min(a, b int)int{
	if a<=b{
		return a
	}
	return b
}
func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i:=0;i<len(grid);i++{ 
		dp[i] = make([]int, len(grid[0]))
		for j:=0;j<len(grid[0]);j++{
			if i-1>=0 && j-1>=0{
				dp[i][j] = min(dp[i-1][j],dp[i][j-1])+grid[i][j]
			} else if i-1 <0 && j-1 >=0{
				dp[i][j] = dp[i][j-1]+grid[i][j]
			} else if j-1<0 && i-1>=0{
				dp[i][j] = dp[i-1][j]+grid[i][j]
			} else{
				dp[i][j] = grid[i][j]
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}
