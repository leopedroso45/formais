package automato

type Automato struct {
 Alfabeto []string
 Estados []string
 FuncTransicao map[string]string
 EstadoInicial int
 EstadosFinais []int
}
//a, b, c
//Q0, Q1, Q3
//Q0aQ0,


