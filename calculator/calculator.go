package main

import "fmt"

func main() {
	var expression string
	fmt.Println("Type the expression")
	fmt.Scanf("%s", &expression)
	fmt.Printf("%s", expression)
}
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