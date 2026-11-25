package main

import (
	"flag"
	"fmt"
	"go-lab/01-basics/hello"
)

func main() {
	var nome string
	flag.StringVar(&nome, "nome", "mundo", "passe um nome")
	flag.Parse()

	fmt.Println(hello.Greeting(nome))
}
