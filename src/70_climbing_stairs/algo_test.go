package algo

import (
	"testing"

	"github.com/test-go/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	assert.Equal(t, 1, climbStairs(1))
	assert.Equal(t, 2, climbStairs(2))
	assert.Equal(t, 3, climbStairs(3))
	assert.Equal(t, 5, climbStairs(4))
	assert.Equal(t, 8, climbStairs(5))
	assert.Equal(t, 13, climbStairs(6))
}
