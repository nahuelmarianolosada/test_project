package coderank

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBirthdayCakeCandles(t *testing.T) {
	testArr := []int32 {18, 90, 90, 13, 90, 75, 90, 8, 90, 43}
	assert.Equal(t, int32(5), BirthdayCakeCandles(testArr))
}
