package ejercicios

import (
	"strconv"
)

func DevuelveString(texto string) (int, string) {
	var numero int
	var error error
	// numero, _ = strconv.Atoi(texto) // se pasa con _ para argumentos que no voy a utilizar
	numero, error = strconv.Atoi(texto)
	if error != nil {
		return 0, "Hubo un error al transformar a numero el texto enviado: " + error.Error()
	}
	if numero > 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "Es menor a 100"
	}
}
