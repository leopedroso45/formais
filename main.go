package main

import (
	"bufio"
	"fmt"
	"main/automato"
	"os"
	"strings"
)

const (

	Green = "\033[1;32m%s\033[0m"
	Red   = "\033[1;31m%s\033[0m"
	Purple = "\033[1;34m%s\033[0m"
	Pink = "\033[1;35m%s\033[0m"
	Yellow  = "\033[1;33m%s\033[0m"
	Teal   = "\033[1;36m%s\033[0m"
)

func main() {

	fmt.Printf(Pink, "Bem vindo ao criador de AFD mais top das galaxia!\n")

	testar := true
	automatoDoCara := criarAutomato()
	exibirAutomato(automatoDoCara)

	for testar {
		fmt.Println("-------------------------------")
		fmt.Printf(Yellow, "Digite uma entrada para teste:\n")
		if recebeEntrada(automatoDoCara) {
			fmt.Printf(Green, "A palavra foi aceita\n")
			fmt.Println("-------------------------------")
		} else {
			fmt.Printf(Red, "A palavra foi recusada\n")
			fmt.Println("-------------------------------")
		}
		fmt.Println("Gostaria de testar uma nova entrada? digite 'S' para Sim e 'N' para nao")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			resposta := scanner.Text()
			if resposta == "S" {
				testar = true
			} else {
				testar = false
			}
			break
		}

	}

}

func recebeEntrada(automato automato.Automato) bool {
	var entrada string
	atualEstado := automato.EstadoInicial
	status := false

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		entrada = scanner.Text()
		break
	}
	//aaa
	//['a', 'a', 'a']
	entradaArray := strings.Split(entrada, "")

	for _, simbolo := range entradaArray {
		transicao := atualEstado + simbolo              //Q0a
		atualEstado = automato.FuncTransicao[transicao] //Q0a -> Q0
	}
	for _, estado := range automato.EstadosFinais {
		if estado == atualEstado {
			status = true
		}
	}
	if status {
		return true
	} else {
		return false
	}

	return status
}

func addFuncTransicao(alfabeto []string) map[string]string {
	funcTrans := make(map[string]string)
	//Q0a -> Q0
	//Q0b -> Q1
	scanner := bufio.NewScanner(os.Stdin)
	temMais := true

	for temMais {
		fmt.Println("Digite uma regra de transição: ")
		fmt.Println("Exemplo: Q0aQ0 ou Q0bQ1 onde Q0 é o estado atual, a/b é a entrada e Q0/Q1 é o estado de destino.")
		for scanner.Scan() {
			//Q00Q0
			regraTransicao := scanner.Text()
			//Q01Q0 EXISTE simbolo
			//simbolos 0 || 1

			regraTransicaoArray := strings.Split(regraTransicao, "")
			simbol := regraTransicaoArray[len(regraTransicaoArray)/2]
			regraTransicaoArray[len(regraTransicaoArray)/2] = "/"
			result1 := strings.Join(regraTransicaoArray, "")

			for _, simbolo := range alfabeto {
				if simbolo == simbol{
					parts := strings.Split(result1, "/")
					estadoAtual := parts[0]
					estadoDestino := parts[1]

					funcTrans[estadoAtual+simbolo] = estadoDestino //Q0a -> Q0
					break
				}
			}
			break
		}
		fmt.Printf(Purple, "Tem mais? [S]im | [N]ão")
		fmt.Println("")

		var result string
		fmt.Scan(&result)

		switch result {
		case "S":
			temMais = true
			break
		case "N":
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
		fmt.Printf(Purple, "Tem mais? [S]im | [N]ão")
		fmt.Println("")

		var result string
		fmt.Scan(&result)

		switch result {
		case "S":
			temMais = true
			break
		case "N":
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
		fmt.Printf(Purple, "Tem mais? [S]im | [N]ão")
		fmt.Println("")

		var result string
		fmt.Scan(&result)

		switch result {
		case "S":
			temMais = true
			break
		case "N":
			temMais = false
			break
		}
	}
	return
}

func addEstadoInicial(estados []string) (estadoInicial string, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite o estado inicial do automato:")
	for scanner.Scan() {
		estadoInicial = scanner.Text()

		for _, estado := range estados {
			if strings.Contains(estadoInicial, estado) {
				return estadoInicial, nil
			}
			//else {
			//	err1 := errors.New("estado desconhecido: o estado informado é desconhecido pela aplicação")
			//	return estadoInicial, err1
			//}
		}
		break
	}
	return estadoInicial, nil
}

func addEstadosFinais(estados []string) (estadosFinais []string) {
	scanner := bufio.NewScanner(os.Stdin)
	temMais := true
	//foi := true

	for temMais {
		fmt.Println("Digite um estado final: ")
		for scanner.Scan() {
			estadoFinal := scanner.Text()
			//q0, q1, q2
			for _, estado := range estados {
				if strings.Contains(estadoFinal, estado) {
					estadosFinais = append(estadosFinais, estadoFinal)
					//foi = true
					break
				}
			}
			break
		}
		fmt.Printf(Purple, "Tem mais? [S]im | [N]ão")
		fmt.Println("")

		var result string
		fmt.Scan(&result)

		switch result {
		case "S":
			temMais = true
			break
		case "N":
			temMais = false
			break
		}
	}
	return estadosFinais
}

func criarAutomato() (aut automato.Automato) {
	aut.Alfabeto = addAlfabeto()
	aut.Estados = addEstados()
	aut.FuncTransicao = addFuncTransicao(aut.Alfabeto)
	aut.EstadoInicial, _ = addEstadoInicial(aut.Estados)
	aut.EstadosFinais = addEstadosFinais(aut.Estados)
	return
}

func exibirAutomato(aut automato.Automato) {
	fmt.Printf(Teal, "-------------------------------\n")
	fmt.Printf(Teal, "Alfabeto:")
	fmt.Println("Alfabeto:", aut.Alfabeto)
	fmt.Printf(Teal, "Estados do Autômato:")
	fmt.Println("Estados do Autômato:", aut.Estados)
	fmt.Printf(Teal, "Função de transição:")
	fmt.Println("Função de transição:", aut.FuncTransicao)
	fmt.Printf(Teal, "Estado Inicial:")
	fmt.Println("Estado Inicial:", aut.EstadoInicial)
	fmt.Printf(Teal, "Conjunto de estados finais:")
	fmt.Println("Conjunto de estados finais:", aut.EstadosFinais)
}

//Remover estados que não recebem
//func verificarEstadosNaoAlcancados(aut automato.Automato) (aut2 automato.Automato) {


	//for _, estado1 := range aut.Estados {
	//
	//	for _, estado2 := range aut.Estados {
//
	//	}
	//}

//}
