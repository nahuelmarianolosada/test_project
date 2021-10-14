package coderank

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAVeryBigSum(t *testing.T) {
	testArray := []int64{1000000001, 1000000002, 1000000003, 1000000004,1000000005}
 	var expectedSum int64 = 5000000015

	assert.Equal(t, expectedSum, AVeryBigSum(testArray))
}