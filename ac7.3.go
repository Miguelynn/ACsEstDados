package main

import (
	"fmt"
)

func main() {
	var codigoPeca1, codigoPeca2, numeroPecas1, numeroPecas2 int
	var valorUnitarioPeca1, valorUnitarioPeca2 float64

	// Leitura dos dados da peça 1
	fmt.Scan(&codigoPeca1, &numeroPecas1, &valorUnitarioPeca1)

	// Leitura dos dados da peça 2
	fmt.Scan(&codigoPeca2, &numeroPecas2, &valorUnitarioPeca2)

	// Cálculo do valor 
	valorTotal := float64(numeroPecas1)*valorUnitarioPeca1 + float64(numeroPecas2)*valorUnitarioPeca2

	// Valor a ser pago
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", valorTotal)
}
