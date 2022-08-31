package main

import (
	"fmt"
	"strconv"
)

var (
	factorialCache map[int]int
	print          = true
)

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	val, ok := factorialCache[n]
	if ok {
		return val
	}
	result := n * factorial(n-1)
	factorialCache[n] = result
	return result
}

func binominalCoefficient(n, k int) int {
	return factorial(n) / (factorial(k) * factorial(n-k))
}

func pascalTriangle(n int) {
	factorialCache = make(map[int]int)
	for y := 0; y < n; y++ {
		for x := 0; x <= y; x++ {
			num := binominalCoefficient(y, x)
			printf("%4s", strconv.Itoa(num))
		}
		printf("\n")
	}
}

func pascalTriangle2(n int) {
	prevRow := []int{}
	for y := 0; y < n; y++ {
		row := []int{}

		for x := 0; x <= y; x++ {
			var num int
			if x == 0 || x == y {
				num = 1
			} else {
				num = prevRow[x] + prevRow[x-1]
			}

			row = append(row, num)
			printf("%4s", strconv.Itoa(num))
		}

		prevRow = row
		printf("\n")
	}
}

func main() {
	pascalTriangle(10)
	fmt.Println()
	fmt.Println("================================")
	fmt.Println()
	pascalTriangle2(10)
}

func printf(format string, any ...interface{}) {
	if !print {
		return
	}
	fmt.Printf(format, any...)
}
