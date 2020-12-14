package main

import (
	"bufio"
	"fmt"
	"main/automato"
	"os"
	"strings"
)

const (
	Green  = "\033[1;32m%s\033[0m"
	Red    = "\033[1;31m%s\033[0m"
	Purple = "\033[1;34m%s\033[0m"
	Pink   = "\033[1;35m%s\033[0m"
	Yellow = "\033[1;33m%s\033[0m"
	Teal   = "\033[1;36m%s\033[0m"
)

var verificados map[string]string

func main() {
	verificados = make(map[string]string)
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
		fmt.Println("Gostaria de testar uma nova entrada? [S]im | [N]ao ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			resposta := scanner.Text()
			if resposta == "S" || resposta == "s" {
				testar = true
			} else if resposta == "N" || resposta == "n" {
				testar = false
			} else {
				testar = true
			}
			break
		}

	}

	aut, result := verificaEquivalencia(automatoDoCara)
	aut = criaAutomatoMinimizado(aut)
	aut = criaFTMinimi(aut, result)
	aut = verificaFinais(aut)
	exibirAutomatoMinimizado(aut)

	testar = true
	for testar {
		fmt.Println("-------------------------------")
		fmt.Printf(Yellow, "Digite uma entrada para teste:\n")
		if recebeEntradaMinimazada(aut) {
			fmt.Printf(Green, "A palavra foi aceita\n")
			fmt.Println("-------------------------------")
		} else {
			fmt.Printf(Red, "A palavra foi recusada\n")
			fmt.Println("-------------------------------")
		}
		fmt.Println("Gostaria de testar uma nova entrada? [S]im | [N]ao ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			resposta := scanner.Text()
			if resposta == "S" || resposta == "s" {
				testar = true
			} else if resposta == "N" || resposta == "n" {
				testar = false
			} else {
				testar = true
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

func recebeEntradaMinimazada(automato automato.Automato) bool {
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
		transicao := atualEstado + simbolo                        //Q0a
		atualEstado = automato.FuncTransicaoMinimizada[transicao] //Q0a -> Q0
	}
	for _, estado := range automato.NovosEstadosFinais {
		fpart := atualEstado[:2]
		spart := atualEstado[2:]
		if estado == atualEstado {
			status = true
		} else if estado == spart+fpart {
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

func resolveTransicao(regraTransicao string) (result1, simbol string) {
	regraTransicaoArray := strings.Split(regraTransicao, "")
	simbol = regraTransicaoArray[len(regraTransicaoArray)/2]
	regraTransicaoArray[len(regraTransicaoArray)/2] = "/"
	result1 = strings.Join(regraTransicaoArray, "")
	return
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

			result1, simbol := resolveTransicao(regraTransicao)

			for _, simbolo := range alfabeto {
				if simbolo == simbol {
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
	fmt.Println(aut.Alfabeto)
	fmt.Printf(Teal, "Estados do Autômato:")
	fmt.Println(aut.Estados)
	fmt.Printf(Teal, "Função de transição:")
	fmt.Println(aut.FuncTransicao)
	fmt.Printf(Teal, "Estado Inicial:")
	fmt.Println(aut.EstadoInicial)
	fmt.Printf(Teal, "Conjunto de estados finais:")
	fmt.Println(aut.EstadosFinais)
}

func exibirAutomatoMinimizado(aut automato.Automato) {
	fmt.Printf(Teal, "-------------------------------\n")
	fmt.Printf(Green, "Automato Minimizado:\n")
	fmt.Printf(Teal, "Alfabeto:")
	fmt.Println(aut.Alfabeto)
	fmt.Printf(Teal, "Estados do Autômato:")
	fmt.Println(aut.EstadosMinimizados)
	fmt.Printf(Teal, "Função de transição:")
	fmt.Println(aut.FuncTransicaoMinimizada)
	fmt.Printf(Teal, "Estado Inicial:")
	fmt.Println(aut.EstadoInicial)
	fmt.Printf(Teal, "Conjunto de estados finais:")
	fmt.Println(aut.NovosEstadosFinais)
}

func verificaFinais(aut automato.Automato) automato.Automato {
	for _, ef := range aut.EstadosFinais {
		for _, em := range aut.EstadosMinimizados {
			if strings.Contains(em, ef) {
				teste := strings.Split(em, ef)
				for i, efe := range aut.EstadosFinais {
					if efe == teste[0] {
						aut.NovosEstadosFinais = append(aut.NovosEstadosFinais, em)
						aut.EstadosFinais[i] = "?"
					}
				}
			}
		}
	}
	return aut
}

func criaFTMinimi(aut automato.Automato, result map[string][]string) automato.Automato {
	aut.FuncTransicaoMinimizada = make(map[string]string)
	for _, novoEstado := range aut.EstadosMinimizados {
		if len(novoEstado) <= 2 {
			for _, simbolo := range aut.Alfabeto {
				estadoAlterar := aut.FuncTransicao[novoEstado+simbolo]
				for _, em := range aut.EstadosMinimizados {
					if strings.Contains(em, estadoAlterar) {
						aut.FuncTransicaoMinimizada[novoEstado+simbolo] = em
					}
				}
			}
		} else if len(novoEstado) > 2 {
			filhos := result[novoEstado]
			for i, simbolo := range aut.Alfabeto {
				aut.FuncTransicaoMinimizada[novoEstado+simbolo] = filhos[i]
			}
		}
	}
	return aut
}

func criaAutomatoMinimizado(aut automato.Automato) automato.Automato {
	var result []bool
	var foi bool
	aut.EstadosMinimizados = append(aut.EstadosMinimizados, aut.NovosEstados...)
	for _, estado := range aut.Estados {
		for _, novoEstado := range aut.NovosEstados {
			if strings.Contains(novoEstado, estado) {
				result = append(result, true)
			}
			result = append(result, false)
		}
		for _, resul := range result {
			if resul {
				foi = true
			}
		}
		if !foi {
			aut.EstadosMinimizados = append(aut.EstadosMinimizados, estado)
		}
	}
	return aut
}

func verificaEquivalencia(aut automato.Automato) (automato.Automato, map[string][]string) {
	pais := verificarEstados(aut) //Q0/Q1
	result := make(map[string][]string)
	//Q0/Q1
	for _, pai := range pais {
		//parts := strings.Split(pai, "/")
		pai1 := pai[:2]
		pai2 := pai[2:]
		result = verificaDoisaDois(pai1, pai2, aut, result)
	}
	for _, pai := range pais {
		var resultFilhos []bool
		filhos := result[pai]
		for _, filho := range filhos {
			_, contidoEmPais := Find(pais, filho)
			if !contidoEmPais {
				ppart := filho[:2]
				spart := filho[2:]
				filhoInvertido := spart + ppart
				_, contidoEmPais = Find(pais, filhoInvertido)
				if contidoEmPais {
					filho = filhoInvertido
				}
			}
			resultFilhos = append(resultFilhos, contidoEmPais)
		}
		count := 0
		for _, result := range resultFilhos {
			if result {
				count++
			}
		}
		//if count == 2*len(aut.Alfabeto) {
		if count == len(aut.Alfabeto) {
			aut.NovosEstados = append(aut.NovosEstados, pai)
		}
	}
	return aut, result
}

func Verificou(estado1, estado2 string) bool {

	if val, ok := verificados[estado2+estado1]; ok && val == estado1 {
		return true
	} else if val, ok = verificados[estado2+estado1]; ok && val == estado2 {
		return true
	} else if val, ok = verificados[estado1+estado2]; ok && val == estado1 {
		return true
	}
	return false
}

func verificarEstados(aut automato.Automato) []string {
	var rejeitados []string
	var pais []string
	//var possiveisEquivalentes []string
	//var Equivalentes []string
	//q0, q2
	for _, estado1 := range aut.Estados {
		//q1, q4
		for _, estado2 := range aut.Estados {
			//for i:=j; i<len(aut.Estados)-1; i++{
			//q4, q2
			if estado1 != estado2 && !Verificou(estado1, estado2) {
				_, found1Final := Find(aut.EstadosFinais, estado1)
				_, found2Final := Find(aut.EstadosFinais, estado2)
				if found1Final && found2Final && estado1 != estado2 {
					//pode adicionar
					//estado1estado2
					pais = append(pais, estado1+estado2)
				} else if !found1Final && found2Final && estado1 != estado2 {
					//rejeita
					rejeitados = append(rejeitados, estado1+estado2)
				} else if found1Final && !found2Final && estado1 != estado2 {
					//rejeita
					rejeitados = append(rejeitados, estado1+estado2)
				} else if !found1Final && !found2Final && estado1 != estado2 {
					//pode adicionar
					//estado1estado2
					pais = append(pais, estado1+estado2)
				}
				verificados[estado1+estado2] = estado2
			}
		}
	}
	fmt.Printf(Red, "Estados não equivalentes: \n")
	fmt.Println(rejeitados)
	return pais
}

func verificaDoisaDois(estado1, estado2 string, aut automato.Automato, result map[string][]string) map[string][]string {
	var filhos []string

	for _, simbolo := range aut.Alfabeto {

		resposta := aut.FuncTransicao[estado1+simbolo]  //Q1
		resposta2 := aut.FuncTransicao[estado2+simbolo] //Q2
		filhos = append(filhos, resposta+resposta2)     //q1q2
		//filhos = append(filhos, resposta2+resposta) //q2q1
	}
	result[estado1+estado2] = filhos
	return result
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
