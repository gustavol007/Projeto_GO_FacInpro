package boletim

// Funções anônimas em Go são funções sem nome, geralmente usadas para criar funções rápidas e locais. Elas podem ser atribuídas a variáveis, passadas como argumento ou executadas diretamente.

// Função padrão
func CalculaMediaAlunos (n1, n2 float64) float64{
	return (n1 + n2)/2
}

// Função anônima
var CalcMedia = func(n1, n2 float64) float64 {
	return (n1 + n2) / 2
}