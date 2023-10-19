/*
 * Escribe un programa que, dado un número, compruebe y muestre si es primo, fibonacci y par.
 * Ejemplos:
 * - Con el número 2, nos dirá: "2 es primo, fibonacci y es par"
 * - Con el número 7, nos dirá: "7 es primo, no es fibonacci y es impar"
 */

package main

import (
	"fmt"
	"math"
	"strconv"
)

func isFibonacciNumber(number int) string {
	if isPerfectSquare(5*number*number+4) || isPerfectSquare(5*number*number-4) {
		return "es un numero fibonacci"
	} else {
		return "no es numero fibonacci"
	}
}

func isPrimeNumber(number int) string {
	if number < 2 {
		return "No es primo y "
	}
	for n := 2; n < number; n++ {
		if number%n == 0 {
			return "No es primo y "
		}
	}
	return "Es primo "
}

func isEvenNumber(number int) string {
	if number%2 == 0 {
		return " es par ,"
	} else {
		return " es impar , "
	}
}

func isPerfectSquare(number int) bool {
	sqrtX := int(math.Sqrt(float64(number)))
	return sqrtX*sqrtX == number
}

func printNumber(number int) string {
	var (
		isEven      string = isEvenNumber(number)
		isPrime     string = isPrimeNumber(number)
		isFibonacci string = isFibonacciNumber(number)
	)
	return "El numero " + strconv.Itoa(number) + " " + isEven + isPrime + isFibonacci
}

func main() {
	fmt.Println(printNumber(5))
	fmt.Println(printNumber(9))
	fmt.Println(printNumber(8))
}
