package algo
import (
	"testing"
	"fmt"
)

func TestNumEnclaves(t *testing.T){
	grid := [][]int{
		{0,0,0,0},
		{1,0,1,0},
		{0,1,1,0},
		{0,0,0,0},
	}
	actual := numEnclaves(grid)
	if actual == 3{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}


func TestNumEnclaves2(t *testing.T){
	grid := [][]int{
		{0,1,1,0},{0,0,1,0},{0,0,1,0},{0,0,0,0},
		}
	actual := numEnclaves(grid)
	fmt.Println(actual)
	if actual == 0{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}
