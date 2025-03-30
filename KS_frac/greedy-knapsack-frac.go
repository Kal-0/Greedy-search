package main

import (
	"fmt"
	"sort"
)

// Item representa um objeto com valor e peso
type Item struct {
	nome  string
	valor int
	peso  int
}

// Função gulosa para o problema da mochila fracionária
func mochilaFracionaria(itens []Item, capacidade int) float64 {
	// Passo 1: Ordenar os itens por maior valor/peso
	sort.Slice(itens, func(i, j int) bool {
		viPi1 := float64(itens[i].valor) / float64(itens[i].peso)
		viPi2 := float64(itens[j].valor) / float64(itens[j].peso)
		return viPi1 > viPi2 // Ordena em ordem decrescente
	})

	// Exibe os itens ordenados
	fmt.Printf("Itens ordenados por rendimento: %v\n", itens)

	// Passo 2: Iterar sobre os itens e adicionar à mochila
	valorTotal := 0.0
	pesoAtual := 0

	for _, item := range itens {
		if pesoAtual+item.peso <= capacidade {
			// Adiciona o item inteiro

			valorTotal += float64(item.valor)
			pesoAtual += item.peso

			fmt.Printf("Adicionando item: %v\n", item)
		} else {
			// Adiciona apenas a fração do item que cabe

			restante := capacidade - pesoAtual
			pesoAtual += restante // Atualiza o peso atual da mochila
			// Calcula o valor proporcional do item que cabe na mochila
			valorTotal += (float64(item.valor) / float64(item.peso)) * float64(restante)

			fmt.Printf("Adicionando fração(%d/%d) do item: %v\n", restante, item.peso, item)
			break // A mochila está cheia, sai do loop
		}
	}

	return valorTotal
}

func main() {
	// Definição dos itens (valor, peso)
	itens := []Item{
		{nome: "Diamante", valor: 100, peso: 50},
		{nome: "Ferro", valor: 25, peso: 25},
		{nome: "Carvao", valor: 5, peso: 20},
		{nome: "Ouro", valor: 50, peso: 100},
		{nome: "Carvao", valor: 5, peso: 20},
		{nome: "Ouro", valor: 50, peso: 100},
		{nome: "Diamante", valor: 100, peso: 50},
		{nome: "Ferro", valor: 25, peso: 25},
	}

	// Exibe os itens
	fmt.Printf("Itens disponíveis: %v\n", itens)

	// Capacidade da mochila
	capacidade := 220

	// Executa o algoritmo guloso
	resultado := mochilaFracionaria(itens, capacidade)

	// Exibe o resultado
	fmt.Printf("Valor máximo da mochila: %.2f\n", resultado)
}
