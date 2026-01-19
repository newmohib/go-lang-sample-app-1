package mathlib

import "fmt"


func Add(a int, b int) int {
	return a + b
}

func Sum(a int, b int) { // parameter
	fmt	.Println("Sum function", Add(a , b))
}