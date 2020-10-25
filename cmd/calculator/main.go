package main

import (
	"fmt"
	"os"

	"github.com/mnogu/go-calculator"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s expression\n", os.Args[0])
		return
	}
	val, err := calculator.Calculate(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	fmt.Println(val)
}
