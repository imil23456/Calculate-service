package calculator

import (
	"errors"
	"strconv"
	"unicode"
)

func isOperator(r rune) bool {
	operators := []rune{'+', '-', '*', '/'}
	for _, op := range operators {
		if r == op {
			return true
		}
	}
	return false
}

func addToSlice(numStr string, nums *[]float64) error {
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return err
	}
	*nums = append(*nums, num)
	return nil
}

func evaluate(nums []float64, ops []rune) (float64, error) {
	if len(ops) == 0 {
		return nums[0], nil
	}
	results := append([]float64{}, nums...)

	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case '*':
			results[i] = results[i] * results[i+1]
			results = append(results[:i+1], results[i+2:]...)
			ops = append(ops[:i], ops[i+1:]...)
			i--
		case '/':
			if results[i+1] == 0 {
				return 0, errors.New("division by zero")
			}
			results[i] = results[i] / results[i+1]
			results = append(results[:i+1], results[i+2:]...)
			ops = append(ops[:i], ops[i+1:]...)
			i--
		}
	}

	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case '+':
			results[i] = results[i] + results[i+1]
			results = append(results[:i+1], results[i+2:]...)
			ops = append(ops[:i], ops[i+1:]...)
			i--
		case '-':
			results[i] = results[i] - results[i+1]
			results = append(results[:i+1], results[i+2:]...)
			ops = append(ops[:i], ops[i+1:]...)
			i--
		}
	}

	return results[0], nil
}

func Calc(expr string) (float64, error) {
	if len(expr) == 0 {
		return 0, errors.New("empty expression")
	}

	var numbers []float64
	var operators []rune
	var currentNum string

	i := 0
	for i < len(expr) {
		r := rune(expr[i])

		if i == 0 && r == '-' {
			currentNum += string(r)
			i++
			continue
		}

		if unicode.IsDigit(r) || r == '.' {
			currentNum += string(r)
		} else if r == '(' {
			subExpr := ""
			j := i + 1
			openBrackets := 1
			for j < len(expr) {
				if expr[j] == '(' {
					openBrackets++
				}
				if expr[j] == ')' {
					openBrackets--
				}
				if openBrackets == 0 {
					break
				}
				subExpr += string(expr[j])
				j++
			}
			if openBrackets != 0 {
				return 0, errors.New("unclosed brackets")
			}

			result, err := Calc(subExpr)
			if err != nil {
				return 0, err
			}
			numbers = append(numbers, result)
			i = j
		} else if isOperator(r) {
			if currentNum != "" {
				err := addToSlice(currentNum, &numbers)
				if err != nil {
					return 0, err
				}
				currentNum = ""
			}
			operators = append(operators, r)
		} else {
			return 0, errors.New("invalid expression")
		}
		i++
	}

	if currentNum != "" {
		err := addToSlice(currentNum, &numbers)
		if err != nil {
			return 0, err
		}
	}

	if len(numbers) == 0 || len(numbers) == len(operators) {
		return 0, errors.New("invalid expression")
	}

	return evaluate(numbers, operators)
}
