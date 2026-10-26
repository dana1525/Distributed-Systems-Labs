// Angajații unei firme sunt stocați într-un dicționar map[string]int cu perechi de forma nume-id_departament.
// Creați un nou map[string]int care să stocheze perechi de tipul id_departament-slice cu numele angajaților din departament.

// Exemplu:
// map[string]int{
//       "Ana Popescu":        1,
//       "Mihai Ionescu":      2,
//       "Bogdan Radu":        1,
//       "Elena Dumitrescu":   3,
//       "Cristian Vasile":    2,
//       "Ioana Enache":       1,
// }
// Rezultatul va fi:
// Departamentul 1: [Ana Popescu, Bogdan Radu, Ioana Enache]
// Departamentul 2: [Mihai Ionescu, Cristian Vasile]
// Departamentul 3: [Elena Dumitrescu]

package main

import (
	"fmt"
	"strings"
)

func main() {
	emp_dep := map[string]int{
		"Ana Popescu":      1,
		"Mihai Ionescu":    2,
		"Bogdan Radu":      1,
		"Elena Dumitrescu": 3,
		"Cristian Vasile":  2,
		"Ioana Enache":     1,
	}

	dep_emp := make(map[int][]string)
	for emp, dept := range emp_dep {
		dep_emp[dept] = append(dep_emp[dept], emp)
	}

	for dep := range dep_emp {
		fmt.Printf("%d: ", dep)
		fmt.Println(strings.Join(dep_emp[dep], ", "))
	}
}
