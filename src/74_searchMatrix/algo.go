package algo

func searchMatrix(matrix [][]int, target int) bool {
	start := 0
	end := len(matrix)*len(matrix[0])-1
	size := len(matrix[0])
	for start <=end{
		middle := (start+end)/2
		x:= middle/size
		y:= middle%size
		if matrix[x][y]==target{
			return true
		}else if matrix[x][y]<target{
			start = middle+1
		} else{
			end = middle-1
		}
	}
	return false
}