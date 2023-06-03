package algo

import (
	"fmt"
	"testing"
)

func TestNextGreaterElements(t *testing.T) {
	actual := nextGreaterElements([]int{1, 2, 3, 4, 3})

	fmt.Println("actual: ", actual)
}

func TestNextGreaterElements2(t *testing.T) {
	actual := nextGreaterElements([]int{1, 2, 1})

	fmt.Println("actual: ", actual)
}
