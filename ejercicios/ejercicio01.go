package ejercicios

import (
	"strconv"
)

func DevuelveString(texto string) (int, string) {
	var numero int
	var error error
	var respuesta string
	numero, error = strconv.Atoi(texto)
	if error != nil {
		respuesta = "Hubo un error al transformar a numero el texto enviado: " + error.Error()
		return numero, respuesta
	}
	if numero > 100 {
		respuesta = "Es mayor a 100"
	} else {
		respuesta = "Es menor a 100"
	}
	return numero, respuesta
}
