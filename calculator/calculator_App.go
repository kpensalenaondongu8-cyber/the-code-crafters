package main

import "fmt"

func main() {

	var num1 int
	var operator string
	var num2 int

	fmt.Println("Enter First Number:")
	fmt.Scan(&num1)
	fmt.Println("Enter Operator:")
	fmt.Scan(&operator)
	fmt.Println("Enter Last Number:")
	fmt.Scan(&num2)

	if operator == "-" {
		fmt.Println(num1 - num2)
	} else if operator == "/" {
		fmt.Println(num1 / num2)
	} else if operator == "*" {
		fmt.Println(num1 * num2)
	} else if operator == "+" {
		fmt.Println(num1 + num2)
	} else {
		fmt.Println("Invalid Syntax:")

	}

}
