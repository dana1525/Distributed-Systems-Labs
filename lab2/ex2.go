// Scrieți un program Go care să transfere cu success un mesaj
// (de exemplu „salut”) de la goroutina main la alta (diferită de main())
// prin intermediul unui channel.

package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)

	wg := sync.WaitGroup{}

	wg.Add(1)

	// lansarea unei goroutine care va asculta pe canal
	go func() {
		defer wg.Done()
		mesaj := <-ch
		fmt.Println("Mesaj primit in goroutina secundara:", mesaj)
	}()

	ch <- "salut"

	wg.Wait()
}
