package algo
import "testing"

func TestLargestRectangleArea(t *testing.T){
	heights := []int{2,1,5,6,2,3}
	actual := largestRectangleArea(heights)

	if actual == 10{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}

}
