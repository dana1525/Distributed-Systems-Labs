// Scrieți un program Go în care funcția main() care primește
// două mesaje pe același canal de comunicare.

package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Primul mesaj\n"
		ch <- "Al doilea mesaj"
	}()

	msg1 := <-ch
	msg2 := <-ch

	fmt.Println(msg1)
	fmt.Println(msg2)
}
