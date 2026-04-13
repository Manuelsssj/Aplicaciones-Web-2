package main

import "fmt"

func main() {
	var a, b float64
	var operacion string
	var opcion string

	historial := ""
	contador := 0

	for {
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
			resultado := a + b
			fmt.Printf("Resultado: %.2f + %.2f = %.2f\n", a, b, resultado)
			historial += fmt.Sprintf("%d) %.2f + %.2f = %.2f\n", contador+1, a, b, resultado)

		case "-":
			resultado := a - b
			fmt.Printf("Resultado: %.2f - %.2f = %.2f\n", a, b, resultado)
			historial += fmt.Sprintf("%d) %.2f - %.2f = %.2f\n", contador+1, a, b, resultado)

		case "*":
			resultado := a * b
			fmt.Printf("Resultado: %.2f * %.2f = %.2f\n", a, b, resultado)
			historial += fmt.Sprintf("%d) %.2f * %.2f = %.2f\n", contador+1, a, b, resultado)

		case "/":
			if b == 0 {
				fmt.Println("Error: no se puede dividir entre cero")
				fmt.Println("-----------------------------------")
				continue
			}
			resultado := a / b
			fmt.Printf("Resultado: %.2f / %.2f = %.2f\n", a, b, resultado)
			historial += fmt.Sprintf("%d) %.2f / %.2f = %.2f\n", contador+1, a, b, resultado)

		case "^":
			resultado := 1.0
			for i := 0; i < int(b); i++ {
				resultado *= a
			}
			fmt.Printf("Resultado: %.2f ^ %.2f = %.2f\n", a, b, resultado)
			historial += fmt.Sprintf("%d) %.2f ^ %.2f = %.2f\n", contador+1, a, b, resultado)

		case "!":
			n := int(a)
			if n == 0 {
				fmt.Println("Resultado: 0! = 1")
				historial += fmt.Sprintf("%d) %d! = 1\n", contador+1, n)
			} else {
				resultado := 1
				for i := 1; i <= n; i++ {
					resultado *= i
				}
				fmt.Printf("Resultado: %d! = %d\n", n, resultado)
				historial += fmt.Sprintf("%d) %d! = %d\n", contador+1, n, resultado)
			}

		default:
			fmt.Println("Error: operación no reconocida")
			fmt.Println("-----------------------------------")
			continue
		}

		contador++

		fmt.Println("-----------------------------------")
		fmt.Print("¿Otra operación? (s/n): ")
		fmt.Scan(&opcion)
		fmt.Println()

		if opcion == "n" {
			break
		}
	}

	fmt.Println("===================================")
	fmt.Println("     HISTORIAL DE OPERACIONES")
	fmt.Println("===================================")
	fmt.Print(historial)
	fmt.Printf("Total de operaciones realizadas: %d\n", contador)
	fmt.Println("¡Hasta luego!")
}
