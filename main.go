package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputs := os.Args

	operator := strings.ToLower(inputs[1])

	switch operator {
	case "add":
		fmt.Println("Operação de adição")
	case "sub":
		fmt.Println("Operação de subtração.")
	case "mut":
		fmt.Println("Operação de multiplicação")
	case "div":
		fmt.Println("Operação de divisão")
	default:
		fmt.Println("Operação não reconhecida.")
	}
}
