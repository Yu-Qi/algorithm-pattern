package algo
import "testing"

func TestMinPathSum(t *testing.T){
	grid := [][]int{
		{1,3,1},{1,5,1},{4,2,1},
	}
	if minPathSum(grid)==7{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}

}
