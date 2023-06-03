package algo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	actual := maxProfit(prices)
	fmt.Println("actual: ", actual)

	assert.Equal(t, 7, actual)
}

// 7, 1,5 ,3,
