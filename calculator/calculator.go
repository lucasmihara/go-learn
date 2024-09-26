package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"errors"
)

const (
	Sum byte = '+'
	Subtraction = '-'
	Multiplication = '*'
	Division = '/'
)

func main() {
	var expression string
	fmt.Println("Type the expression")
	fmt.Scanf("%s", &expression)
	
	if !isValid(expression) {
		fmt.Println("The expression is invalid")
		return
	}
	
	result, err := calculateSums(expression)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func calculateSums(expression string) (float64, error) {
	parts, operations := breakParts(expression, []byte {Sum, Subtraction})
	convertedValue, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		convertedValue, err = calculateMultiplications(parts[0])
		if err != nil {
			return 0, err
		}
	}
	result := convertedValue
	for i := 0; i < len(operations); i++ {
		convertedValue, err = strconv.ParseFloat(parts[i + 1], 64)
		if err != nil {
			convertedValue, err = calculateMultiplications(parts[i + 1])
			if err != nil {
				return 0, err
			}
		}
		switch operations[i] {
			case Sum:
				result += convertedValue
			case Subtraction:
				result -= convertedValue
		}
	}
	return result, nil
}

func calculateMultiplications(expression string) (float64, error) {
	parts, operations := breakParts(expression, []byte {Multiplication, Division})
	convertedValue, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		convertedValue, err = resolveParentheses(parts[0])
		if err != nil {
			return 0, err
		}
	}
	result := convertedValue
	for i := 0; i < len(operations); i++ {
		convertedValue, err = strconv.ParseFloat(parts[i + 1], 64)
		if err != nil {
			convertedValue, err = resolveParentheses(parts[i + 1])
			if err != nil {
				return 0, err
			}
		}
		switch operations[i] {
			case Multiplication:
				result *= convertedValue
			case Division:
				if convertedValue == 0 {
					return 0, errors.New("Division by zero")
				}
				result /= convertedValue
		}
	}
	return result, nil
}

func resolveParentheses(expression string) (float64, error) {
	return calculateSums(expression[1:(len(expression) - 1)])
}

func breakParts(expression string, operators []byte) ([]string, []byte) {
	var parts []string
	var operations []byte
	parenthesesCount := 0
	
	start, end := 0, 0
	for ; end < len(expression); end++ {
		if expression[end] == '(' {
			parenthesesCount++
		} else if expression[end] == ')' {
			parenthesesCount--
		} else if parenthesesCount == 0 && (slices.Contains(operators, expression[end])) {
			parts = append(parts, expression[start:end])
			operations = append(operations, expression[end])
			start = end + 1
		}
	}
	parts = append(parts, expression[start:end])

	return parts, operations
}

func isValid(expression string) bool {
	isInvalid, err := regexp.MatchString(`[^0-9\(\)\+\-\*\/]+`, expression)
	if err != nil || isInvalid {
		return false
	}

	parenthesesCount := 0
	for i := 0; i < len(expression); i++ {
		if expression[i] == '(' {
			parenthesesCount++
		} else if expression[i] == ')' {
			parenthesesCount--
			if parenthesesCount < 0 {
				return false
			}
		}
	}
	if parenthesesCount > 0 {
		return false
	}

	return true
}