package main

import (
	"fmt"
)

// Calcular o máximo divisor comum (MDC) de dois números
func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Simplificar uma fração
func simplificar(numerator, denominator int) (int, int) {
	divisorComum := mdc(numerator, denominator)
	return numerator / divisorComum, denominator / divisorComum
}

// Soma de duas frações
func soma(n1, d1, n2, d2 int) (int, int) {
	numerator := n1*d2 + n2*d1
	denominator := d1 * d2
	return simplificar(numerator, denominator)
}

// Subtração de duas frações
func subtracao(n1, d1, n2, d2 int) (int, int) {
	numerator := n1*d2 - n2*d1
	denominator := d1 * d2
	return simplificar(numerator, denominator)
}

// Multiplicação de duas frações
func multiplicacao(n1, d1, n2, d2 int) (int, int) {
	numerator := n1 * n2
	denominator := d1 * d2
	return simplificar(numerator, denominator)
}

// Divisão de duas frações
func divisao(n1, d1, n2, d2 int) (int, int) {
	numerator := n1 * d2
	denominator := n2 * d1
	return simplificar(numerator, denominator)
}

func main() {
	var casos int
	fmt.Scanln(&casos)

	for i := 0; i < casos; i++ {
		var n1, d1, n2, d2 int
		var operacao string
		fmt.Scanln(&n1, "/", &d1, &operacao, &n2, "/", &d2)

		var resultadoN, resultadoD int
		switch operacao {
		case "+":
			resultadoN, resultadoD = soma(n1, d1, n2, d2)
		case "-":
			resultadoN, resultadoD = subtracao(n1, d1, n2, d2)
		case "*":
			resultadoN, resultadoD = multiplicacao(n1, d1, n2, d2)
		case "/":
			resultadoN, resultadoD = divisao(n1, d1, n2, d2)
		}

		fmt.Printf("%d/%d = %d/%d\n", n1, d1, resultadoN, resultadoD)
	}
}
