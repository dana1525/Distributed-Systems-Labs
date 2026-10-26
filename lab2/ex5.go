// Scrieți un program Go în care un goroutine(send) trimite un
// mesaj unui alt goroutine(receive) printr-un canal. Receive va
// returna mesajul primit, iar funcția main() îl va afișa pe ecran.

package main

import "fmt"

// trimite mesajul pe canal
func send(ch chan<- string) {
	ch <- "salut"
}

// primeste mesajul din ch si il returneaza prin result
func receive(ch <-chan string, result chan<- string) {
	mesaj := <-ch
	result <- mesaj
}

func main() {
	ch := make(chan string)
	result := make(chan string)

	go send(ch)
	go receive(ch, result)

	fmt.Println("Main a primit:", <-result)
}
