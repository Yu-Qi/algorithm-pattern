package algo

import (
	"testing"

	"github.com/test-go/testify/assert"
)

func TestJump(t *testing.T) {
	assert.Equal(t, 2, jump([]int{1, 2, 3}))
	assert.Equal(t, 0, jump([]int{0}))
	assert.Equal(t, 2, jump([]int{2, 3, 0, 1, 4}))
	assert.Equal(t, 1, jump([]int{2, 1}))
}
