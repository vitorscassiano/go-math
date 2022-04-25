package main

import (
	"fmt"
	"os"
)

func main() {
	inputs := os.Args
	fmt.Println(inputs[1:])
}
