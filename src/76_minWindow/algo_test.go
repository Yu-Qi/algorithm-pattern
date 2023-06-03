package algo

import (
	"fmt"
	"testing"
)

func TestMinWindow1(t *testing.T) {
	actual := minWindow("ADOBECODEBANC", "ABC")
	fmt.Println("actual: ", actual)
}

// ADOBEC, 0,5
// DOBEC, 1, 5
// DOBECODEBA, 1,10, [1, 2, 1]
// ODEBA,

func TestMinWindow2(t *testing.T) {
	actual := minWindow("a", "a")
	fmt.Println("actual: ", actual)
}

func TestMinWindow3(t *testing.T) {
	actual := minWindow("a", "aa")
	fmt.Println("actual: ", actual)
}

func TestMinWindow4(t *testing.T) {
	actual := minWindow("aa", "aa")
	fmt.Println("actual: ", actual)
}
