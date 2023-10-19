package main

/*
 * Crea una función que encuentre todas las combinaciones de los números
 * de una lista que suman el valor objetivo.
 * - La función recibirá una lista de números enteros positivos
 *   y un valor objetivo.
 * - Para obtener las combinaciones sólo se puede usar
 *   una vez cada elemento de la lista (pero pueden existir
 *   elementos repetidos en ella).
 * - Ejemplo: Lista = [1, 5, 3, 2],  Objetivo = 6
 *   Soluciones: [1, 5] y [1, 3, 2] (ambas combinaciones suman 6)
 *   (Si no existen combinaciones, retornar una lista vacía)
 */

import (
	"fmt"
)

type DataSum struct {
	Numbers   []int
	Objective int
}

func sumNumbers(n *DataSum) [][]int {

	if len(n.Numbers) == 0 || n.Objective < 0 {
		return [][]int{}
	}

	result := [][]int{}
	seen := make(map[string]bool)

	for i := 0; i < len(n.Numbers); i++ {
		newResults := sumNumbers(&DataSum{Numbers: n.Numbers[i+1:], Objective: n.Objective - n.Numbers[i]})

		if n.Objective == n.Numbers[i] {
			result = append(result, []int{n.Numbers[i]})
			seen[fmt.Sprint(n.Numbers[i])] = true
		}

		for _, r := range newResults {
			combination := append([]int{n.Numbers[i]}, r...)
			key := fmt.Sprint(combination)
			if !seen[key] {
				result = append(result, combination)
				seen[key] = true
			}
		}
	}

	return result
}
func main() {
	data := DataSum{Numbers: []int{1, 1, 1, 2, 2, 1, 1}, Objective: 6}
	combinations := sumNumbers(&data)
	fmt.Printf("Numbers is: %d\n", data.Numbers)
	fmt.Printf("Objective is: %d\n", data.Objective)
	fmt.Printf("combinations: %d\n", combinations)

}
