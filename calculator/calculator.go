package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
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
	
	result := calculateSums(expression)
	fmt.Println(result)
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

func calculateSums(expression string) float64 {
	parts, operations := breakParts(expression, []byte {Sum, Subtraction})
	convertedValue, _ := strconv.ParseFloat(parts[0], 32)
	result := convertedValue
	for i := 0; i < len(operations); i++ {
		convertedValue, _ = strconv.ParseFloat(parts[i + 1], 32)
		switch operations[i] {
			case Sum:
				result += convertedValue
			case Subtraction:
				result -= convertedValue
		}
	}
	// for i := 0; i < len(parts); i++ {
	// 	fmt.Printf("%s ", parts[i])
	// }
	// for i := 0; i < len(operations); i++ {
	// 	fmt.Printf("%c ", operations[i])
	// }
	return result
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

// func calculateMultiplications(expression string) float32 {
// 	var parts []string
// 	var operations []Operation
	
// }

// func resolveParentheses(expression string) float32 {
// 	var parts []string
// 	var operations []Operation
	
// }

// func main() {
// 	var n1, n2 float32
// 	var operation rune

// 	fmt.Println("Type the first number")
// 	fmt.Scanf("%v\n", &n1)
// 	fmt.Scanf("%v\n", &n2)
// 	fmt.Println("Type the operation")
// 	fmt.Scanf("%c", &operation)

// 	var result float32
// 	switch operation {
// 		case '+':
// 			result = n1 + n2
// 		case '-':
// 			result = n1 - n2
// 		case '*':
// 			result = n1 * n2
// 		case '/':
// 			if n2 == 0 {
// 				fmt.Println("Division by zero")
// 				return
// 			}
// 			result = n1 / n2
// 		default:
// 			fmt.Println("This operation is not supported")
// 			return
// 	}

// 	fmt.Printf("Result: %v\n", result)
// }