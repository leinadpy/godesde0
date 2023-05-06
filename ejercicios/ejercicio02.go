package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)
	var numero int
	var err error
	var texto string

	fmt.Println("Ingrese el número para la tabla de multiplicación: ")
	for {
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("El dato ingresado es incorrecto " + err.Error())
				fmt.Println("Vuelva a ingresar el número para la tabla de multiplicación")
				continue
			}
			break
		}
	}

	texto = fmt.Sprintln("La tabla de multiplicar del", numero)
	for i := 1; i < 11; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, numero*i)
	}

	return texto
}
