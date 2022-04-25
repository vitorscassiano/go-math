package main

import (
	"fmt"
	"learningmath/pkg/operations"
	"os"
	"strings"
)

func main() {
	inputs := os.Args

	operator := strings.ToLower(inputs[1])

	switch operator {
	case "add":
		resultado := operations.Add(1, 2)
		fmt.Println(resultado)
	case "sub":
		// resultado := operations.Sub(1, 2)
		// fmt.Println(resultado)
		fmt.Println("Operação de subtração")
	case "mut":
		// resultado := operations.Mut(1, 2)
		// fmt.Println(resultado)
		fmt.Println("Operação de multiplicação")
	case "div":
		// resultado := operations.Div(1, 2)
		// fmt.Println(resultado)
		fmt.Println("Operação de divisão")
	default:
		fmt.Println("Operação não reconhecida.")
	}
}
