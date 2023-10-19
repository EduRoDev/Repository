package main

/*
 * Escribe un programa que sea capaz de generar contraseñas de forma aleatoria.
 * Podrás configurar generar contraseñas con los siguientes parámetros:
 * - Longitud: Entre 8 y 16.
 * - Con o sin letras mayúsculas.
 * - Con o sin números.
 * - Con o sin símbolos.
 * (Pudiendo combinar todos estos parámetros entre ellos)
 */

import (
	"fmt"
	"math/rand"
	"strconv"
)

var input string

func generar_contraseña(t int) {
	// Variables
	letras := "qwertyuiopasdfghjklñzxcvbnm"
	letras_mayusculas := "QWERTYUIOPASDFGHJKLÑZXCVBNM"
	numeros := "1234567890"
	caracteres := "!#$%&/()=.-_,+*;:"
	contraseña := ""

	// Bucle
	for i := 1; i <= t && len(contraseña) < 16; i++ {
		q := rand.Intn(4)
		switch q {
		case 0:
			n := rand.Intn(len(letras))
			contraseña += string(letras[n])
		case 1:
			n := rand.Intn(len(letras_mayusculas))
			contraseña += string(letras_mayusculas[n])
		case 2:
			n := rand.Intn(len(numeros))
			contraseña += string(numeros[n])
		case 3:
			n := rand.Intn(len(caracteres))
			contraseña += string(caracteres[n])
		default:
			return
		}
	}

	if len(contraseña) >= 16 {
		contraseña = contraseña[:16]
		fmt.Println("El limite para caracteres es 16")
	}

	fmt.Printf("%s %s", "su contraseña es: ", contraseña)
}

func main() {
	fmt.Println("Tamaño de la contraseña:")
	fmt.Scanln(&input)
	tamaño, _ := strconv.Atoi(input)
	generar_contraseña(tamaño)
}
