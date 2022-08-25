package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		param    string
		expected int
	}{
		{
			param:    "1 + 2 / 2 * 3",
			expected: 4,
		},
		{
			param:    "(1 + 2) / 2 * 3",
			expected: 3,
		},
		{
			param:    "(1+((2+1)*2))/2*3",
			expected: 9,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, evaluate(test.param))
	}
}
