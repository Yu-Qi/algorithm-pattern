package algo
import (
	"testing"
	"fmt"
)

func TestNumIslands(t *testing.T){
	grid := [][]byte{
		{'1','1','1','1','0'},
		{'1','1','0','1','0'},
		{'1','1','0','0','0'},
		{'0','0','0','0','0'},
	}
	actual := numIslands(grid)
	fmt.Println(actual)
	if actual == 1 {
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}

func TestNumIslands2(t *testing.T){
	grid := [][]byte{
		{'1','1','0','0','0'},
		{'1','1','0','0','0'},
		{'0','0','1','0','0'},
		{'0','0','0','1','1'},
	}
	actual := numIslands(grid)
	fmt.Println(actual)
	if actual == 3 {
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}
