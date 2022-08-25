package main

import (
	"strconv"
)

func isOperator(s string) bool {
	return s == "+" ||
		s == "-" ||
		s == "*" ||
		s == "/"
}

func getPriority(s string) int {
	if s == "*" || s == "/" {
		return 2
	}
	return 1
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}

func performOperation(operands []int, operators []string) ([]int, []string) {
	lastOperandIdx := len(operands) - 1
	lastOperatorIdx := len(operators) - 1
	operand1 := operands[lastOperandIdx-1]
	operand2 := operands[lastOperandIdx]
	operands = operands[:lastOperandIdx-1]
	operands = append(operands, calculate(operand1, operand2, operators[lastOperatorIdx]))
	operators = operators[:lastOperatorIdx]
	return operands, operators
}

func canPerformOperation(operands []int, operators []string) bool {
	if len(operands) < 2 {
		return false
	}
	if operators[len(operators)-1] == "(" {
		return false
	}
	return true
}

func evaluate(str string) int {
	data := parseInput(str)
	operators := []string{}
	operands := []int{}

	for _, s := range data {
		if s == "(" {
			operators = append(operators, s)
			continue
		}

		if s == ")" {
			for operators[len(operators)-1] != "(" {
				operands, operators = performOperation(operands, operators)
			}
			operators = operators[:len(operators)-1]
			continue
		}

		if !isOperator(s) {
			x, _ := strconv.Atoi(s)
			operands = append(operands, x)
			continue
		}

		if !canPerformOperation(operands, operators) ||
			getPriority(operators[len(operators)-1]) < getPriority(s) {
			operators = append(operators, s)
			continue
		}

		operands, operators = performOperation(operands, operators)
		operators = append(operators, s)
	}

	for len(operators) > 0 {
		operands, operators = performOperation(operands, operators)
	}

	return operands[0]
}

func parseInput(str string) []string {
	ret := []string{}
	number := ""
	for _, c := range str {
		if c >= '0' && c <= '9' {
			number += string(c)
			continue
		}

		if number != "" {
			ret = append(ret, number)
			number = ""
		}

		if c == ' ' {
			continue
		}

		ret = append(ret, string(c))
	}

	if number != "" {
		ret = append(ret, number)
	}

	return ret
}
