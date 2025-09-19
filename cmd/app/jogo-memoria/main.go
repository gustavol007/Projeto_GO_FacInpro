package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/seu-usuario/meu-projeto-go/internal/jogosequencia"

)


func main() {
	rand.Seed(time.Now().UnixNano())

	const tamanho = 5
	const digitoMax = 9

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Jogo da Memória (Go) ===")
	fmt.Println("Objetivo: memorize a sequência e digite na mesma ordem.")
	fmt.Println()

	seq := jogosequencia.GerarSequencia(tamanho, digitoMax)

	// Mostra a sequência para memorizar
	fmt.Println("Sequência (memorize):", jogosequencia.Juntar(seq))
	fmt.Println("Pressione ENTER quando estiver pronto para responder...")
	_, _ = reader.ReadString('\n')

	// "Limpa" a tela (ou empurra para cima)
	jogosequencia.LimparTela()

	// Pede o palpite
	palpite := jogosequencia.LerPalpite(reader, tamanho)

	// Calcula pontuação
	acertos := jogosequencia.Pontuar(seq, palpite)

	// Mostra resultado
	fmt.Println()
	fmt.Println("Sequência correta:", jogosequencia.Juntar(seq))
	fmt.Println("Sua resposta     :", jogosequencia.Juntar(palpite))
	fmt.Printf("Você fez %d/%d ponto(s). ", acertos, tamanho)

	switch acertos {
	case 5:
		fmt.Println("Perfeito! Mandou muito bem!")
	case 3, 4:
		fmt.Println("Quase lá! Bora de novo?")
	case 1, 2:
		fmt.Println("Tá no caminho. Treina mais um pouco!")
	default:
		fmt.Println("Agora vai! Tenta outra vez!")
	}
}