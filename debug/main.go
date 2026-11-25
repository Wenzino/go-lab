package main

import "fmt"

type Item struct {
	Name  string
	Price float64
	Qty   int
}

func total(items []Item) float64 {
	sum := 0.0
	for _, it := range items {
		sum += it.Price * float64(it.Qty)
	}
	return sum
}

func main() {
	items := []Item{
		{"caneta", 10, 2},
		{"caderno", 5, 1},
	}
	fmt.Println("total:", total(items)) // esperado: 25
}
