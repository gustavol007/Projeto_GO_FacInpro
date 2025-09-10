// Arquivo principal do programa (entrypoint)
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários 
import (
    "fmt"
    "github.com/seu-usuario/meu-projeto-go/internal/hello"
    "github.com/seu-usuario/meu-projeto-go/internal/fibonacci"
)

// Função principal do programa
func main() {
    fmt.Println("🚀 Meu primeiro projeto em Go com estrutura de mercado!")
    hello.SayHello()
    
    var pos int
	fmt.Print("Digite a posição da sequência de Fibonacci: ")
	fmt.Scan(&pos)

	resultado := fibonacci.Fibonacci(pos)
	fmt.Printf("O número de Fibonacci na posição %d é %d\n", pos, resultado)
    
}
