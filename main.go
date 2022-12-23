package main

import "fmt"

var n1 int
var n2 int
var result int

func main() {
	fmt.Println("Digite o primeiro número para calcular: ")
	fmt.Scan(&n1)
	fmt.Println("Digite o segundo número para calcular: ")
	fmt.Scan(&n2)

	menuPrincipal()

	fmt.Scan(&result)

	switch result {
	case 1:
		fmt.Println(somarNum(n1, n2))
	case 2:
		fmt.Println(subtrairNum(n1, n2))
	case 3:
		fmt.Println(multiplicarNum(n1, n2))
	case 4:
		fmt.Println(dividirNum(n1, n2))
	default:
		fmt.Println("Você não entrou com um número acima. Tente novamente!")
	}
}
