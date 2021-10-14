package coderank

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetPositiveNegativeAndZero(t *testing.T) {
	testArr := []int32{-4, 3, -9, 0, 4, 1}

	pos, neg, zeros := GetPositiveNegativeAndZero(testArr)
	assert.Equal(t, 3.0, pos)
	assert.Equal(t, 2.0, neg)
	assert.Equal(t, 1.0, zeros)

	plusMinus(testArr)
}
