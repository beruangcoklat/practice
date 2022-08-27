package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	tests := []struct {
		num1     string
		num2     string
		expected string
	}{
		{
			num1:     "99",
			num2:     "999",
			expected: "1098",
		},
		{
			num1:     "199",
			num2:     "19",
			expected: "218",
		},
		{
			num1:     "10",
			num2:     "0",
			expected: "10",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, add(test.num1, test.num2))
	}
}
