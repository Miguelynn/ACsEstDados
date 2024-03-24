package main

import (
	"fmt"
)

func main() {
	var N int
	var L, Q float64

	// Leitura do número de pessoas na roda, volume da garrafa térmica e quantidade de água na cuia
	fmt.Scanln(&N, &L, &Q)

	participants := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&participants[i])
	}

	// Cálculo da quantidade total de água 
	totalWater := Q * float64(N)

	// Encontrando o participante que será o "rico do mate"
	richestIndex := int(L / Q) % N

	// Imprimindo o nome do participante e a quantidade de água
	fmt.Printf("%s %.1f\n", participants[richestIndex], totalWater)
}
