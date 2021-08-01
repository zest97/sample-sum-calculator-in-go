package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	value1 := getInputValue("Value 1: ")
	value2 := getInputValue("Value 2: ")

	operator := getOperator("Select an operation (+ - * /): ")

	calc := Calculator{value1, value2, operator}
	result := calc.operate()

	fmt.Printf("The result of calculation is %v\n", result)
}

func getInputValue(prompt string) float64 {
	input, _ := askInput(prompt)
	return parseStringToInt(input)
}

func getOperator(prompt string) string {
	input, _ := askInput(prompt)
	return strings.TrimSpace(input)
}

func askInput(prompt string) (string, error) {
	fmt.Printf("%v: ", prompt)
	return reader.ReadString('\n')
}

func parseStringToInt(x string) float64 {
	result, err := strconv.ParseFloat(strings.TrimSpace(x), 64)
	if err != nil {
		panic("Input type must be a number.")
	}
	return result
}

type Calculator struct {
	first     float64
	second    float64
	operation string
}

func (c Calculator) operate() float64 {
	var result float64
	switch c.operation {
	case "+":
		result = sum(c.first, c.second)
	case "-":
		result = substract(c.first, c.second)
	case "*":
		result = multiply(c.first, c.second)
	case "/":
		result = divide(c.first, c.second)
	default:
		panic("Invalid operator syntax.")
	}

	return math.Round(result*100) / 100
}

func sum(x float64, y float64) float64 {
	return x + y
}

func substract(x float64, y float64) float64 {
	return x - y
}

func multiply(x float64, y float64) float64 {
	return x * y
}

func divide(x float64, y float64) float64 {
	return x / y
}
