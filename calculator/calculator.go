package main

import (
	"fmt"
	"regexp"
)

// const Operation (
// 	Sum = '+'
// 	Subtraction = '-'
// 	Multiplication = '*'
// 	Division = '/'
// )

func main() {
	var expression string
	fmt.Println("Type the expression")
	fmt.Scanf("%s", &expression)
	
	isInvalid, err := regexp.MatchString(`[^0-9\(\)\+\-\*\/]+`, expression)
	if err != nil || isInvalid {
		fmt.Println("The Expression is invalid")
		return
	}

}

// func calculateSums(expression string) float32 {
// 	var parts []string
// 	var operations []Operation
// 	parenthesesCount := 0

// 	for start, end := 0; end < len(expression); end++ {

// 	}
// }

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