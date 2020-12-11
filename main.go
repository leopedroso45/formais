package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//aut := automato.Automato{}
	//var receitaAutomato string
	alfabeto := addAlfabeto()
	fmt.Println(alfabeto)
	fmt.Println(addEstados())
	fmt.Println(addFuncTransicao(alfabeto))

}

func addFuncTransicao(alfabeto []string) map[string]string{
	funcTrans := make(map[string]string)
	//Q0a -> Q0
	//Q0b -> Q1
	scanner := bufio.NewScanner(os.Stdin)
	temMais := true

	for temMais {
		fmt.Println("Digite uma regra de transição: ")
		fmt.Println("Exemplo: Q0aQ0 ou Q0bQ1 onde Q0 é o estado atual, a/b é a entrada e Q0/Q1 é o estado de destino.")
		for scanner.Scan() {
			regraTransicao := scanner.Text()

			for _, simbolo := range alfabeto {
				if strings.Contains(regraTransicao, simbolo) {
					parts := strings.Split(regraTransicao, simbolo)
					estadoAtual := parts[0]
					estadoDestino := parts[1]

					funcTrans[estadoAtual+simbolo] = estadoDestino //Q0a -> Q0
					break
				} else {
					fmt.Println("O simbolo informado está incoerente com o alfabeto.")
					temMais = true
					continue
				}
			}
			break
		}
		fmt.Println("Tem mais? 1 para sim | 2 para não")
		var result int
		fmt.Scan(&result)

		switch result {
		case 1:
			temMais = true
			break
		case 2:
			temMais = false
			break
		}
	}
	return funcTrans
}

func addEstados() (estados []string) {
	scanner := bufio.NewScanner(os.Stdin)
	temMais := true

	for temMais {
		fmt.Println("Digite um estado: ")
		for scanner.Scan() {
			estado := scanner.Text()
			estados = append(estados, estado)
			break
		}
		fmt.Println("Tem mais? 1 para sim | 2 para não")
		var result int
		fmt.Scan(&result)

		switch result {
		case 1:
			temMais = true
			break
		case 2:
			temMais = false
			break
		}
	}
	return
}

func addAlfabeto() (alfabeto []string) {
	scanner := bufio.NewScanner(os.Stdin)
	temMais := true

	for temMais {
		fmt.Println("Digite um simbolo do alfabeto: ")
		for scanner.Scan() {
			simbolo := scanner.Text()
			alfabeto = append(alfabeto, simbolo)
			break
		}
		fmt.Println("Tem mais? 1 para sim | 2 para não")
		var result int
		fmt.Scan(&result)

		switch result {
		case 1:
			temMais = true
			break
		case 2:
			temMais = false
			break
		}
	}
	return
}
