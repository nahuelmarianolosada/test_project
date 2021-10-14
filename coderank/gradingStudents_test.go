package coderank

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGradingStudents(t *testing.T) {
	testGrades := []int32{
		73,
		67,
		38,
		33,
	}

	expectedGrades := []int32{
		75,
		67,
		40,
		33,
	}
	assert.Equal(t, expectedGrades, GradingStudents(testGrades))
}