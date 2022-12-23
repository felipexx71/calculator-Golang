package main

import "fmt"

func menuPrincipal() {
	fmt.Println("\nCalculadora --- Operadores\n\n1) Adição\n2) Subtração\n3) Multiplicação\n4) Divisão")
}

func somarNum(x, y int) int {
	return (x + y)
}

func dividirNum(x, y int) int {
	return (x / y)
}

func multiplicarNum(x, y int) int {
	return (x * y)
}

func subtrairNum(x, y int) int {
	return (x - y)
}
