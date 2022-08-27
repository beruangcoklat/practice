package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	t.Parallel()
	tests := []struct {
		number   string
		divisor  int
		expected string
	}{
		{
			number:   "12313413534672234",
			divisor:  754,
			expected: "16330787181262",
		},
		{
			number:   "1248163264128256512",
			divisor:  125,
			expected: "9985306113026052",
		},
		{
			number:   "12",
			divisor:  13,
			expected: "0",
		},
		{
			number:   "12",
			divisor:  12,
			expected: "1",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, divide(test.number, test.divisor))
	}
}
