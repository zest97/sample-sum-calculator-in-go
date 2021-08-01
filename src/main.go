package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseStringToInt(x string) float64 {
	result, err := strconv.ParseFloat(strings.TrimSpace(x), 64)
	if err != nil {
		panic("Input type must be a number.")
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Value 1: ")
	input1, _ := reader.ReadString('\n')
	float1 := parseStringToInt(input1)

	fmt.Print("Value 2: ")
	input2, _ := reader.ReadString('\n')
	float2 := parseStringToInt(input2)

	sum := float1 + float2
	sum = math.Round((sum * 100)) / 100

	fmt.Printf("The sum of %v and %v is %v\n", float1, float2, sum)
}
