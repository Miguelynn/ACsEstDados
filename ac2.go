package main

import (
    "fmt"
)

func main() {
    // Vetor de inteiros
    vetor := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println("Elementos do vetor:")
    for _, num := range vetor {
        fmt.Println(num)
    }

    // Função para inverter uma string
    fmt.Println("\nDigite uma string:")
    var input string
    fmt.Scanln(&input)
    fmt.Println("String invertida:", inverterString(input))

    // Tipo Ponto e método DistanciaOrigem()
    ponto := Ponto{3, 4}
    fmt.Printf("\nDistância do ponto (%d, %d) até a origem: %.2f\n", ponto.X, ponto.Y, ponto.DistanciaOrigem())

    // Uso do pacote geometria para calcular área e perímetro de um retângulo
    var base, altura float64
    fmt.Println("\nDigite a base do retângulo:")
    fmt.Scanln(&base)
    fmt.Println("Digite a altura do retângulo:")
    fmt.Scanln(&altura)

    area := CalcularAreaRetangulo(base, altura)
    perimetro := CalcularPerimetroRetangulo(base, altura)

    fmt.Printf("Área do retângulo: %.2f\n", area)
    fmt.Printf("Perímetro do retângulo: %.2f\n", perimetro)
}

// Função para inverter uma string
func inverterString(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

// Definição do tipo Ponto
type Ponto struct {
    X, Y float64
}

// Método para calcular a distância do ponto até a origem
func (p Ponto) DistanciaOrigem() float64 {
    return (p.X*p.X + p.Y*p.Y) * 0.5
}

// Funções para calcular a área e o perímetro de um retângulo
func CalcularAreaRetangulo(base, altura float64) float64 {
    return base * altura
}

func CalcularPerimetroRetangulo(base, altura float64) float64 {
    return 2 * (base + altura)
}
