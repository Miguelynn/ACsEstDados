package main

import (
	"bufio"
	"fmt"
	"os"
)

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) AlterarEmail(novoEmail string) {
	c.Email = novoEmail
}

func main() {
	var contatos [5]Contato
	var nome, email, opcao string

	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Lista de contatos!")
	for {
		fmt.Print("Digite (1) para adicionar, (2) para remover ou qualquer outra coisa para sair: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Informe o seu nome: ")
			nome, _ = leitor.ReadString('\n')

			fmt.Print("Informe o seu email: ")
			fmt.Scanln(&email)

			contatos = adicionaContato(Contato{Nome: nome, Email: email}, contatos)
		case "2":
			contatos = excluiContato(contatos)
		default:
			return
		}

		fmt.Println(contatos)
	}

}

func adicionaContato(c Contato, a [5]Contato) [5]Contato {
	for ind, contato := range a {
		if (contato == Contato{}) {
			a[ind] = c
			break
		}
	}

	return a
}

func excluiContato(a [5]Contato) [5]Contato {
	if (a[0] == Contato{}) {
		fmt.Println("Lista de contatos vazia!")
		return a
	}

	for ind, contato := range a {
		if (contato == Contato{}) {
			a[ind-1] = Contato{}
		}
	}

	return a
}
