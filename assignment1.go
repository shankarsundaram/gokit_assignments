package main

import (
	"fmt"
	"log"
)

var acc float64

func doCalculate(input float64, operat string) float64 {
	switch operat {
	case "+":
		acc = acc + input
	case "-":
		acc = acc - input
	case "*":
		acc = acc * input
	case "/":
		acc = acc / input
	default:
		log.Println("Invalid operator")
	}
	return acc

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
	result = doCalculate(10, "@")
	fmt.Println(result)
}
