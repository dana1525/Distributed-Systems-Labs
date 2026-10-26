// Scrieți un program Go care stochează un array în memorie. Programul trebuie să afișeze tabloul
// sortat crescător folosind o metodă de sortare la alegere.

package main

import (
	"fmt"
)

func main() {
	arr := [5]int32{4, 7, 2, 1, 5}

	for i := range len(arr) {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	fmt.Print(arr)
}
