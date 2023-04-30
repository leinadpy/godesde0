package main

import (
	"fmt"

	"github.com/leinadpy/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1588)

	fmt.Println(estado)
	fmt.Println(texto)
}
