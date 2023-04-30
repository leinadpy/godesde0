package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() {
	scanner := bufio.NewScanner(os.Stdin)
	var numero int
	var err error

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

	fmt.Println("La tabla de multiplicar del", numero)
	for i := 1; i < 11; i++ {
		fmt.Println(numero * i)
	}
}
