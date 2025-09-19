package main

import (
	"fmt"

	"github.com/seu-usuario/meu-projeto-go/internal/boletim"
)

func main (){

	var n1, n2 float64

	fmt.Print("Digite a nota 1: ")
	fmt.Scanln(&n1)
	
	fmt.Print("Digite a nota 2: ")
	fmt.Scanln(&n2)

	resultado := boletim.CalculaMediaAlunos(n1, n2)
	fmt.Printf("A média das notas é: %.2f\n", resultado)

	resultado2 := boletim.CalcMedia(n1,n2)
	fmt.Printf("A média das notas é (Função Anônima): %.2f\n", resultado2)

	if resultado >= 7 {
		fmt.Println("Aluno aprovado!")
	} else {
		fmt.Println("Aluno reprovado!")
	}

}