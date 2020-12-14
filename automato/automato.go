package automato

type Automato struct {
 Alfabeto []string // a, b, c, ...
 Estados []string //Q1
 NovosEstados []string //Q1Q3
 EstadosMinimizados []string
 FuncTransicaoMinimizada map[string]string //Q0aQ0
 FuncTransicao map[string]string //Q0aQ0
 EstadoInicial string
 EstadosFinais []string
 NovosEstadosFinais []string
}



