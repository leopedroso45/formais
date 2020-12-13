package automato

type Automato struct {
 Alfabeto []string
 Estados []string
 FuncTransicao map[string]string
 EstadoInicial string
 EstadosFinais []string
}

type Transicao struct {
 Estados []string
 FuncTransicao map[string]string
}


