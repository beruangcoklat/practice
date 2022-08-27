package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	t.Parallel()
	tests := []struct {
		num1     string
		num2     string
		expected string
	}{
		{
			num1:     "408",
			num2:     "5",
			expected: "2040",
		},
		{
			num1:     "123456789",
			num2:     "987654321",
			expected: "121932631112635269",
		},
		{
			num1:     "123",
			num2:     "0",
			expected: "0",
		},
		{
			num1:     "666",
			num2:     "6",
			expected: "3996",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, multiply(test.num1, test.num2))
	}

}
