package main

import (
	"fmt"

	"github.com/seu-usuario/meu-projeto-go/internal/anamnese"
)

func main() {
	var p anamnese.Paciente

	fmt.Print("Digite o nome: ")
	fmt.Scanln(&p.Nome)

	fmt.Print("Digite a idade: ")
	fmt.Scanln(&p.Idade)

	fmt.Print("Digite o peso (kg): ")
	fmt.Scanln(&p.Peso)

	fmt.Print("Digite a altura (m): ")
	fmt.Scanln(&p.Altura)

	fmt.Printf("Paciente: %+v\n", p)

	imc := anamnese.CalcularIMC(p.Peso, p.Altura)
	fmt.Printf("IMC: %.2f\n", imc)

	switch {
	case imc < 18.5:
		fmt.Println("Abaixo do peso")
	case imc < 25:
		fmt.Println("Peso normal")
	case imc < 30:
		fmt.Println("Sobrepeso")
	case imc < 35:
		fmt.Println("Obesidade grau I")
	case imc < 40:
		fmt.Println("Obesidade grau II")
	default:
		fmt.Println("Obesidade grau III")
	}

}
