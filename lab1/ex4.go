// Scrieți un program care stochează în memorie un slice de float32 și citește de la tastatură un
// index. Programul trebuie să elimine elementul de la poziția index, păstrând ordinea celorlalte
// elemente.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	slice := []float32{3, 4.0, 2.6, 8, -1}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	index, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("eroare conversie")
		return
	}

	if index < 0 || index > len(slice)-1 {
		fmt.Println("index invalid")
		return
	}

	// for i := range len(slice) {
	// 	if i == index {
	// 		for j := i; j < len(slice)-1; j++ {
	// 			slice[j] = slice[j+1]
	// 		}
	// 		break
	// 	}
	// }

	slice = append(slice[:index], slice[(index+1):]...)
	fmt.Print(slice)
}
