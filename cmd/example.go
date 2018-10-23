package main

import (
	"fmt"
	c "golang-sample-project-with-makefile-and-test/pkg/calc"
)

func main() {
	calc := &c.Calc{}

	calc.Sum(1)
	calc.Sub(1)
	calc.Mult(1)
	calc.Div(1)

	fmt.Println(calc.GetResult())
	fmt.Println(calc.GetHistory())
}
