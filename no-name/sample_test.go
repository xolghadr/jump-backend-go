package noname

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	testSlice := []int{1, 2, 3}
	AddElement(&testSlice, 4)
	assert.EqualValues(t, []int{1, 2, 3, 4}, testSlice)
	m := FindMin(&testSlice)
	assert.Equal(t, m, 1)
	ReverseSlice(&testSlice)
	assert.EqualValues(t, []int{4, 3, 2, 1}, testSlice)
	SwapElements(&testSlice, 3, 2)
	assert.EqualValues(t, []int{4, 3, 1, 2}, testSlice)
}
