// Scrieți un program Go care stochează în memorie un array cu numere naturale. Programul va citi de la
// tastatură un număr s și va afișa în O(n) indicii elementelor care însumate au ca rezultat s.

// Exemplu:
// 6 15
// 2 7 11 8 4 9

// Rezultatul va fi:
// 1 3
// 2 4

package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	comp := make(map[int][]int) // complement suma: indicii la care apare
	for i, val := range arr {
		for _, j := range comp[val] {
			fmt.Println(j, i)
		}
		comp[s-val] = append(comp[s-val], i)
	}
}
