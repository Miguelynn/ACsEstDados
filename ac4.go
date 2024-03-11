package main

import (
	"fmt"
)

func torreDeHanoi(n int, origem, destino, auxiliar string) {
	if n == 1 {
		fmt.Printf("Mova o disco 1 de %s para %s\n", origem, destino)
		return
	}
	torreDeHanoi(n-1, origem, auxiliar, destino)
	fmt.Printf("Mova o disco %d de %s para %s\n", n, origem, destino)
	torreDeHanoi(n-1, auxiliar, destino, origem)
}

func encontrarPar(array []int, alvo int) (int, int) {
	n := len(array)
	i, j := 0, n-1

	for i < j {
		soma := array[i] + array[j]
		if soma == alvo {
			return array[i], array[j]
		} else if soma < alvo {
			i++
		} else {
			j--
		}
	}

	return -1, -1
}

func main() {

	fmt.Println("Solução para a Torre de Hanói:")
	torreDeHanoi(3, "A", "C", "B")

	fmt.Println("\nSolução para encontrar um par cuja soma seja igual a 9:")
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	alvo := 9
	num1, num2 := encontrarPar(array, alvo)
	if num1 != -1 && num2 != -1 {
		fmt.Printf("Par encontrado: %d e %d\n", num1, num2)
	} else {
		fmt.Println("Nenhum par encontrado")
	}
}
