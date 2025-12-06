package main

import "fmt"

func main() {

	var a, b, result float64
	var operator string
	fmt.Println("Enter the first number: ")
	fmt.Scanln(&a)
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&b)
	fmt.Println("Enter the operator[+,-,*,/]: ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		result = a + b
		fmt.Println(result)
	case "-":
		result = a - b
		fmt.Println(result)
	case "*":
		result = a * b
		fmt.Println(result)
	case "/":
		if b == 0 {
			fmt.Println("Divide by zero is not allow")
			return
		}
		result = a / b
		fmt.Println(result)
	default:
		fmt.Println("Invalid Operator")
	}

}
