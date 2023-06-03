package algo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenLock(t *testing.T) {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	actual := openLock(deadends, "0202")
	fmt.Println("actual: ", actual)
	assert.Equal(t, 6, actual)
}
