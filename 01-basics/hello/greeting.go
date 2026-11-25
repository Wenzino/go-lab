package hello

import "fmt"

func Greeting(nome string) string {
	return fmt.Sprintf("Hello %s! Welcome to Go packages.", nome)
}
