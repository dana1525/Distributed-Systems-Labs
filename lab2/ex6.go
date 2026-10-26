// Fie un sistem distribuit format din doi roboți identici. Ei sunt porniți de către un observer,
// reprezentat de funcția main(). Cei doi roboți pornesc în același timp și se întrec în linie dreaptă
// până la un finish. Scrieți un program Go în care simulați acest sistem, unde funcția main() va
// afișa care robot a ajuns primul la finish. Hint: folosiți select.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const steps = 5

func robot(name string, start <-chan struct{}, finish chan<- string) {
	<-start // asteapta startul de la observer (daca nu exista nicio valoare de citit si canalul nu e inchis, se blocheaza goroutina)

	for i := range steps {
		fmt.Println(name, "- pasul", i)
		ms := 50 + rand.Intn(100) // numar intre 50 si 149
		pauza := time.Duration(ms) * time.Millisecond
		time.Sleep(pauza)
	}

	finish <- name
}

func main() {
	start := make(chan struct{})
	finish1 := make(chan string, 1) // daca canalul ar fi unbuffered, robotul care pierde ar ramane blocat la trimitere
	finish2 := make(chan string, 1)

	go robot("Robot 1", start, finish1)
	go robot("Robot 2", start, finish2)

	fmt.Println("Observer: start!")
	close(start) // schimbare stare deschis -> inchis (robotii blocati pe <-start se deblocheaza)
	// o citire de pe un canal inchis nu asteapta niciodata, ci returneaza valoarea zero a tipului
	// (struct{} nu are campuri, deci avem struct{}{}, care nu contine informatie)

	// select asteapta simultan pe finish1 si finish2 si executa ramura primului canal care primeste o val
	select {
	case name := <-finish1:
		fmt.Println("Observer:", name, "a ajuns primul la finish")
	case name := <-finish2:
		fmt.Println("Observer:", name, "a ajuns primul la finish")
	}
}
