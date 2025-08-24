package calculator

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func Calculator() {
	var result float64
	num1, operator, num2 := getArgs()
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
			os.Exit(1)
		}
		result = num1 / num2
		operator = "÷"
	case "^":
		result = math.Pow(num1, num2)
	default:
		fmt.Printf("Error: Invalid operator '%s'\n", operator)
		fmt.Println("Supported operators: +, -, *, /, ^")
		os.Exit(1)
	}

	fmt.Printf("Result: %.1f %s %.1f = %.2f\n", num1, operator, num2, result)

}

func getArgs() (float64, string, float64) {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <number1> <operator> <number2>")
		fmt.Println("Operators: +, -, *, /, ^")
		os.Exit(1)
	}
	//Get args
	num1Str := os.Args[1]
	operator := os.Args[2]
	num2Str := os.Args[3]

	num1, err := strconv.ParseFloat(num1Str, 64)
	if err != nil {
		fmt.Printf("Error: Invalid number '%s'\n", num1Str)
		os.Exit(1)
	}

	num2, err := strconv.ParseFloat(num2Str, 64)
	if err != nil {
		fmt.Printf("Error: Invalid number '%s'\n", num2Str)
		os.Exit(1)
	}
	return num1, operator, num2
}
