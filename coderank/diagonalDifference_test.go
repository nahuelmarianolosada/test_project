package coderank

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDiagonalDifference(t *testing.T) {
	testMatrix := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	var expectedDiff int32 = 2
	assert.Equal(t, expectedDiff, DiagonalDifference(testMatrix))
}