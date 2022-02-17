package algo

func findMinNeight(mat *[][]int, i, j int)int{
	min := 1<<31
	if i-1 >=0 && (*mat)[i-1][j]<min{
		min = (*mat)[i-1][j]
	}
	if i+1 < len(*mat) &&  (*mat)[i+1][j]<min{
		min = (*mat)[i+1][j]
	}

	if j-1 >=0 && (*mat)[i][j-1]<min{
		min = (*mat)[i][j-1]
	}
	if j+1 < len((*mat)[0]) && (*mat)[i][j+1]<min{
		min = (*mat)[i][j+1]
	}
	return min
}
func updateMatrix_v0(mat [][]int) [][]int {
	// timeout
    dst_mat := make([][]int, len(mat))
	for i:=0;i<len(mat);i++{
        dst_mat[i] = make([]int, len(mat[0]))
		for j:=0;j<len(mat[0]);j++{
			dst_mat[i][j] = 1<<31
		}
	}

    change := true

	for change == true{
		change = false
		for i:=0;i<len(mat);i++{
			for j:=0;j<len(mat[0]);j++{
				if mat[i][j] == 0 && dst_mat[i][j]!=0{
					dst_mat[i][j] = 0
					change = true
				} else{
                    min := findMinNeight(&dst_mat,i,j)
					if min+1 != dst_mat[i][j]{
						dst_mat[i][j] = min +1
						change = true
					}
				}
			}
		}
	}
    return dst_mat
}

type Point struct{
	X int
	Y int
}
func updateMatrix_v1(mat [][]int) [][]int {
	queue := []Point{}
	// create
    dst_mat := make([][]int, len(mat))
	for i:=0;i<len(mat);i++{
        dst_mat[i] = make([]int, len(mat[0]))
		for j:=0;j<len(mat[0]);j++{
			if mat[i][j] == 0{
				queue = append(queue, Point{i,j})
				dst_mat[i][j] = 0
			}else{
				dst_mat[i][j] = 1<<31
			}
		}
	}
	count := 1
	for len(queue) >0{
		size := len(queue)
        for idx:=0;idx<size;idx++{
			x,y := queue[idx].X, queue[idx].Y
            if x+1<len(mat)&& dst_mat[x+1][y] == 1<<31{
				dst_mat[x+1][y] = count
				queue = append(queue, Point{x+1,y})
			}
			if x-1>=0 && dst_mat[x-1][y] == 1<<31{
				dst_mat[x-1][y] = count
				queue = append(queue, Point{x-1,y})
			}
            if y+1 < len(mat[0]) && dst_mat[x][y+1] == 1<<31{
				dst_mat[x][y+1] = count
				queue = append(queue, Point{x,y+1})
			}
			if y-1 >=0 && dst_mat[x][y-1] == 1<<31{
				dst_mat[x][y-1] = count
				queue = append(queue, Point{x,y-1})
			}									
		}
		count ++
		queue = queue[size:]
	}
    return dst_mat
}


func updateMatrix_v2(mat [][]int) [][]int {
	queue := []Point{}
	// create
	for i:=0;i<len(mat);i++{
		for j:=0;j<len(mat[0]);j++{
			if mat[i][j] == 0{
				queue = append(queue, Point{i,j})
			}else{
				// 除了0以外的值都不重要
				mat[i][j] = -1
			}
		}
	}
	direction := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
	for len(queue) >0{
		point := queue[0]
		queue = queue[1:]
		for i:=0;i<len(direction);i++{
			x := point.X + direction[i][0]
			y := point.Y + direction[i][1]
			if x>=0 && y>=0 && x<len(mat)&&y<len(mat[0])&&mat[x][y]==-1{
				mat[x][y] = mat[point.X][point.Y]+1
				queue = append(queue, Point{x,y})
			}
		}
		
    return mat
	}
}