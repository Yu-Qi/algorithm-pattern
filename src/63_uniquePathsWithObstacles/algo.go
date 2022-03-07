package algo


func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// init
	metrix := make([][]int, len(obstacleGrid))
	for i:=0;i<len(obstacleGrid);i++{
		row := make([]int, len(obstacleGrid[0]))
		row[0] = 1
		metrix[i] = row
	}
	
	for i:=0;i<len(obstacleGrid[0]);i++{
		metrix[0][i] = 1
	}
	if obstacleGrid[0][0] == 1{
		metrix[0][0] = 0
	} else{
		metrix[0][0] = 1
	}
	// 判斷
	for i:=1;i<len(obstacleGrid);i++{
		if obstacleGrid[i][0] == 1 {
			metrix[i][0] = 0
		}else{
			metrix[i][0] = metrix[i-1][0]
		}
	}
	for j:=1;j<len(obstacleGrid[0]);j++{
		if obstacleGrid[0][j] == 1 {
			metrix[0][j] = 0
		} else{
			metrix[0][j] = metrix[0][j-1]
		}
	}	
	for i:=1;i<len(obstacleGrid);i++{
		for j:=1;j<len(obstacleGrid[0]);j++{
			if obstacleGrid[i][j] == 1 {
				metrix[i][j] = 0
			} else{
				metrix[i][j] = metrix[i-1][j] + metrix[i][j-1]
			}
		}
	}	
	return metrix[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}