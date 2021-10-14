package coderank

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTimeConversion(t *testing.T) {
	testTime := "07:05:45PM"
	expectedRsp := "19:05:45"
	assert.Equal(t, expectedRsp, TimeConversion(testTime))
}
