package main

import "fmt"

//Faça um algoritmo que
//leia o preço de um produto e mostre o seu valor com desconto de 10%.

func main() {
	var precoProduto float64
	fmt.Print("Qual é o preço do produto? ")
	fmt.Scan(&precoProduto)

	desconto := precoProduto * 0.1
	resultado := precoProduto - desconto

	fmt.Println("O preço com desconto é", resultado)
}
