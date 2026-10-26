// Scrieți un program Go care stochează în memorie un șir de numere și care afișează un slice cu
// numerele pare din acel șir.

package main

import "fmt"

func main() {
	arr := [5]int32{4, 7, 2, 1, 5}

	slice := []int32{}
	for i := range len(arr) {
		if arr[i]%2 == 0 {
			slice = append(slice, arr[i])
		}
	}

	fmt.Print(slice)
}
