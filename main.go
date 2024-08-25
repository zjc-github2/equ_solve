package main

import (
	"fmt"

	"none.com/equ_solve/lib"
)

func main() {
	equ := "3 + 5 a - b * h"
	fmt.Println(lib.EuqToArr(equ))
}
