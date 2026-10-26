// Scrieți un program Go care ia elementele unui slice de string-uri și îl modifică astfel încât
// elementele să apară în ordine inversă. Nu folosiți un alt slice pentru această operațiune.

// Exemplu:
// slice := []string"Laboratorul", "de", "sisteme", "distribuite".
// Rezultatul va fi: ["distribuite" "sisteme" "de" "Laboratorul"]

package main

import "fmt"

func main() {
	slice := []string{"Laboratorul", "de", "sisteme", "distribuite", "ele"}

	for i := range len(slice) / 2 {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}

	fmt.Println(slice)
}
