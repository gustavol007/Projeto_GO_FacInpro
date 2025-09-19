package anamnese

type Paciente struct {
	Nome   string
	Idade  int
	Peso   float64
	Altura float64
}

func CalcularIMC(peso, altura float64) float64 {
	return peso / (altura * altura)
}
