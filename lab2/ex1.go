// Scrieți un program Go care să transfere cu succes un mesaj
// (de exemplu „salut”) de la o goroutină (diferită de main())
// la goroutina main() prin intermediul unui channel.

package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	// pornire goroutina in paralel cu functia main
	go func() {
		defer close(ch)
		ch <- "salut" // trimitere sir in canal
	}()

	mesaj := <-ch // citire sir din canal
	fmt.Println(mesaj)
}
