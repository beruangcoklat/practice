package main

import (
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	rows := []string{}
	count := -1
	for i := len(num1) - 1; i >= 0; i-- {
		count++
		carrier := 0
		rowResult := ""

		for j := 0; j < count; j++ {
			rowResult += "0"
		}

		for j := len(num2) - 1; j >= 0; j-- {
			x := byteToInt(num1[i])*byteToInt(num2[j]) + carrier
			carrier = 0
			if x > 9 {
				carrier = x / 10
				x %= 10
			}

			rowResult = strconv.Itoa(x) + rowResult
		}
		if carrier > 0 {
			rowResult = strconv.Itoa(carrier) + rowResult
		}

		rows = append(rows, rowResult)
	}

	i := 0
	result := ""
	carrier := 0
	for {
		i++
		total := carrier
		carrier = 0
		invalid := true
		for j := 0; j < len(rows); j++ {
			idx := len(rows[j]) - i
			if idx >= 0 {
				invalid = false
				total += byteToInt(rows[j][idx])
			}
		}

		if total > 9 {
			carrier = total / 10
			total %= 10
		}

		result = strconv.Itoa(total) + result

		if invalid {
			break
		}
	}

	if carrier > 0 {
		result = strconv.Itoa(carrier) + result
	}

	result = strings.TrimLeft(result, "0")
	if result == "" {
		result = "0"
	}

	return result
}
