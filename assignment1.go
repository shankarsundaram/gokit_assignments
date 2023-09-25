package main

import (
	"fmt"
)

func doCalculate(y float64, z string) float64 {
	var x float64
	switch z {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		return x / y
	default:
		return 0
	}

}

func main() {
	var result float64
	result = doCalculate(100, "+")
	fmt.Println(result)
	result = doCalculate(50, "-")
	fmt.Println(result)
	result = doCalculate(20, "/")
	fmt.Println(result)
	result = doCalculate(10, "*")
	fmt.Println(result)
}
