package main

import (
	"strconv"
	"strings"
)

func byteToInt(a byte) int {
	return int(a) - '0'
}

func divide(number string, divisor int) string {
	result := ""
	i := 0
	carrier := 0
	for {
		x := carrier
		count := 0
		for x < divisor && i < len(number) {
			count++
			x *= 10
			x += byteToInt(number[i])
			i++
		}

		if x < divisor {
			break
		}

		r := strconv.Itoa(x / divisor)
		for len(r) < count {
			r = "0" + r
		}
		result += r
		carrier = x % divisor
	}

	result = strings.TrimLeft(result, "0")
	if result == "" {
		result = "0"
	}
	return result
}
