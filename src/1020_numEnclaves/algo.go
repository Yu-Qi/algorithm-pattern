package algo

func numEnclaves(grid [][]int) int {
	// flush
	for i:=0; i<len(grid);i++{
	    dfs(&grid, i, 0)
	    dfs(&grid, i, len(grid[0])-1)
	}
	for j:=0; j<len(grid[0]);j++{
		dfs(&grid, 0, j)
		dfs(&grid, len(grid)-1, j)
	}

	move := 0
	for i:=0; i<len(grid);i++{
		for j:=0;j<len(grid[0]);j++{
			if grid[i][j]==1{
				move  += dfs(&grid,i,j)
			} 
		}
	}
	return move
}


func dfs(grid *[][]int, i, j int)int{
    if i>=len(*grid) || j>= len((*grid)[0]) || i<0 || j<0{
		return 0
	}
	if (*grid)[i][j]== 1{
		(*grid)[i][j] = 0
		return dfs(grid,i-1,j)+ dfs(grid,i+1,j)+ dfs(grid,i,j-1)+ dfs(grid,i,j+1)+1
	}
	return 0
}