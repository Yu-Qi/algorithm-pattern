package algo
import(
	"testing"
	"fmt"
)

func TestClosedIsland(t *testing.T){
	grid := [][]int{
		{1,1,1,1,1,1,1,0},
		{1,0,0,0,0,1,1,0},
		{1,0,1,0,1,1,1,0},
		{1,0,0,0,0,1,0,1},
		{1,1,1,1,1,1,1,0},
	}
	actual := closedIsland(grid)
	fmt.Println(actual)
	if actual == 2{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}

}
