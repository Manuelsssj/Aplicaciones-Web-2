package main

import "fmt"

func main() {
	a := 10
	b := 3

	fmt.Println("===================================")
	fmt.Println("   CALCULADORA CIENTÍFICA v1.0")
	fmt.Println("===================================")
	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b)
	fmt.Printf("a / b = %.2f  ", float64(a)/float64(b))
}
