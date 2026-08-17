package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter first number: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	num1, err := strconv.ParseFloat(input1, 64)
	if err != nil {
		fmt.Println("invalid input", input1)
		return
	}
	fmt.Println("input operation (*, /, -, +): ")
	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)
	fmt.Println("input second number: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	num2, err := strconv.ParseFloat(input2, 64)
	if err != nil {
		fmt.Println("invalid input", input2)
		return
	}
	var result float64
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("invalid division by zero")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("unknown operation", op)
		return
	}
	fmt.Printf("Result: %.2f\n", result)

}
