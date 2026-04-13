package main

import "fmt"

func main() {
	var a, b float64
	var operacion string

	fmt.Println("===================================")
	fmt.Println("   CALCULADORA CIENTÍFICA v1.0")
	fmt.Println("===================================")

	fmt.Print("Ingrese el primer número: ")
	fmt.Scan(&a)

	fmt.Print("Ingrese el segundo número: ")
	fmt.Scan(&b)

	fmt.Print("Ingrese la operación (+, -, *, /, ^, !): ")
	fmt.Scan(&operacion)

	fmt.Println("-----------------------------------")

	switch operacion {
	case "+":
		fmt.Printf("Resultado: %.2f + %.2f = %.2f\n", a, b, a+b)

	case "-":
		fmt.Printf("Resultado: %.2f - %.2f = %.2f\n", a, b, a-b)

	case "*":
		fmt.Printf("Resultado: %.2f * %.2f = %.2f\n", a, b, a*b)

	case "/":
		if b == 0 {
			fmt.Println("Error: no se puede dividir entre cero")
			return
		}
		fmt.Printf("Resultado: %.2f / %.2f = %.2f\n", a, b, a/b)

	case "^":
		resultado := 1.0
		for i := 0; i < int(b); i++ {
			resultado *= a
		}
		fmt.Printf("Resultado: %.2f ^ %.2f = %.2f\n", a, b, resultado)

	case "!":
		n := int(a)
		if n == 0 {
			fmt.Println("Resultado: 0! = 1")
		} else {
			resultado := 1
			for i := 1; i <= n; i++ {
				resultado *= i
			}
			fmt.Printf("Resultado: %d! = %d\n", n, resultado)
		}

	default:
		fmt.Println("Error: operación no reconocida")
	}

	fmt.Println("-----------------------------------")
	fmt.Println("Programa finalizado")
}
