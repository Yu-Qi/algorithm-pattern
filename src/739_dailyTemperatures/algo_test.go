package algo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDailyTemperatures1(t *testing.T) {
	actual := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	fmt.Println("actual: ", actual)
	assert.Equal(t, []int{1, 1, 4, 2, 1, 1, 0, 0}, actual)
}

func TestDailyTemperatures2(t *testing.T) {
	actual := dailyTemperatures([]int{30, 40, 50, 60})
	fmt.Println("actual: ", actual)
	assert.Equal(t, []int{1, 1, 1, 0}, actual)
}

func TestDailyTemperatures3(t *testing.T) {
	actual := dailyTemperatures([]int{30})
	fmt.Println("actual: ", actual)
	assert.Equal(t, []int{0}, actual)
}

func TestDailyTemperatures4(t *testing.T) {
	actual := dailyTemperatures([]int{3, 2, 1})
	fmt.Println("actual: ", actual)
	assert.Equal(t, []int{0, 0, 0}, actual)
}
