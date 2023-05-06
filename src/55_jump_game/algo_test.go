package algo

import (
	"testing"

	"github.com/test-go/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	assert.Equal(t, true, canJump([]int{2, 3, 1, 1, 4}))
	assert.Equal(t, true, canJump([]int{1, 1, 2, 2, 0, 1, 1}))
	assert.Equal(t, true, canJump([]int{2, 0, 0}))
	assert.Equal(t, true, canJump([]int{2}))
	assert.Equal(t, true, canJump([]int{0}))
	assert.Equal(t, false, canJump([]int{0, 1}))
}
