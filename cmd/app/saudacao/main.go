package main

import (
	"fmt"

	"github.com/seu-usuario/meu-projeto-go/internal/saudacao"
)

func main () {
	mensagem := saudacao.Saudacao("Gustavo")
	fmt.Println(mensagem)
}