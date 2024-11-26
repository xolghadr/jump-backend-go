package armstrong

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExtractNumbers(t *testing.T) {
	digits := decodeStringIntoDigits("123someRandomText89andSthElse03")

	assert.EqualValues(t, 123, digits[0])
	assert.EqualValues(t, 89, digits[1])
	assert.EqualValues(t, 3, digits[2])
}

func Test_CheckArmstrong(t *testing.T) {
	result := checkArmstrong(153)
	assert.True(t, result)

	result = checkArmstrong(12)
	assert.False(t, result)
}
