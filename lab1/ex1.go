// Scrieți un program Go care citește un număr natural de la tastatură și verifică dacă este prim sau nu.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())
	prim := true

	if n < 2 {
		prim = false
	} else {
		for d := 2; d*d <= n; d++ {
			if n%d == 0 {
				prim = false
				break
			}
		}
	}

	if prim {
		fmt.Printf("%d e prim\n", n)
	} else {
		fmt.Printf("%d nu e prim\n", n)
	}
}
