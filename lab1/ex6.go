// Fie un map de tipul string-int care stochează perechi de tipul nume-înălțime. Să se afișeze numele
// persoanelor în ordinea descrescătoare a înălțimilor. Dacă două persoane au aceeași înălțime, nu
// contează ordinea de afișare.

// Exemplu de input:
// map[string]int{
// 		"Andrei": 	190,
// 		"Maria": 	165,
// 		"Ion": 		182,
// 		"Elena": 	172,
// 		"George": 	178,
// 		"Ana": 		165,
// 	}

package main

import "fmt"

func main() {
	myMap := map[string]int{
		"Andrei": 190,
		"Maria":  165,
		"Ion":    182,
		"Elena":  172,
		"George": 178,
		"Ana":    165,
	}

	// extragere nume in slice
	names := make([]string, 0, len(myMap))
	for name := range myMap {
		names = append(names, name)
	}

	// sortare descrescatoare dupa inaltime
	for i := 0; i < len(names)-1; i++ {
		for j := i + 1; j < len(names); j++ {
			if myMap[names[i]] < myMap[names[j]] {
				names[i], names[j] = names[j], names[i]
			}
		}
	}

	fmt.Println(names)
}
