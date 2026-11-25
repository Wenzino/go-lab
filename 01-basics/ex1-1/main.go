package main

import (
	"fmt"

	"flag"
)

// PARTE A
// Este exercicio nos foi pedido para elaborar um FizzBuzz + tabuada:
// Onde para os numeros de 1 a 100, imprimindo:
// 'Fizz' se for multiplo de 3;
// 'Buzz' se for multiplo de 5;
// 'FizzBuzz' se for multiplo de ambos;
// o proprio numero nos outros casos
func exerciseA() {

	// primeiro vamos garantir de imprimir todos valores ate que a condicao seja satisfeita
	for i := 1; i <= 100; i++ {
		// vamos garantir que o caso mais especifico sempre esteja no topo
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// PARTE B
// Uma tabuada de 1-10, onde:
// iremos receber um numero (n) inteiro e imprimir a sua tabuada
func exerciseB() {
	var number int
	flag.IntVar(&number, "number", 5, "numero da tabuada")
	flag.Parse()

	for i := 0; i < 10; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}

}

func main() {
	fmt.Print("======== Parte A ========\n")
	exerciseA()
	fmt.Println("\n========================\n\n====== Parte B ======")
	exerciseB()
}
