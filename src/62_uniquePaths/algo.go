func uniquePaths(m int, n int) int {
    metrix := make([][]int, m)
	for i:=0;i<m;i++{
		row := make([]int, n)
		row[0] = 1
		metrix[i] = row
	}
	for i:=0;i<n;i++{
		metrix[0][i] = 1
	}

	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			metrix[i][j] = metrix[i-1][j]+ metrix[i][j-1]
		}
	}
	return metrix[m-1][n-1]
}