// Fie un sistem distribuit format din doi roboți identici. Ei sunt porniți de către un observer,
// reprezentat de funcția main(). Fiecare robot va parcurge un traseu și va trimite periodic printr-
// un canal coordonatele sale 2D către celălalt robot, în timp ce se deplasează. Când un robot
// își termină de parcurs traseul, îl va notifica pe celălalt că s-a oprit. Acesta trebuie să citească
// toate pozițiile primite până la oprire, iar apoi ambii vor raporta observer-ului că și-au terminat
// sarcina. Simulați, folosind un program Go, comportamentul roboților.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Point struct {
	X, Y int
}

const steps = 3

// robot: trimite pozitiile pe out (buffer), apoi citeste pozitiile celuilalt de pe in
func robot(name string, start <-chan struct{}, out chan<- Point, in <-chan Point, done chan<- string) {
	<-start // asteapta startul de la Observer

	// parcurge traseul si trimite periodic pozitia (bufferul face ca transmiterea sa nu se blocheze)
	x, y := 0, 0
	for i := range steps {
		// robotul inainteaza pe diagonala: la fiecare pas, x si y cresc cu 1..3
		x += rand.Intn(3) + 1
		y += rand.Intn(3) + 1
		p := Point{x, y}

		fmt.Println(name, "- pasul", i, "- pozitia", p)
		out <- p

		ms := 50 + rand.Intn(100) // numar intre 50 si 149
		pauza := time.Duration(ms) * time.Millisecond
		time.Sleep(pauza)
	}
	close(out) // notifica celalalt robot ca s-a oprit

	// citeste toate pozitiile celuilalt robot, pana cand acesta se opreste
	for p := range in {
		fmt.Println(name, "a primit pozitia", p)
	}

	done <- name
}

func main() {
	start := make(chan struct{})
	done := make(chan string)

	ch12 := make(chan Point, steps)
	ch21 := make(chan Point, steps)

	go robot("Robot 1", start, ch12, ch21, done)
	go robot("Robot 2", start, ch21, ch12, done)

	fmt.Println("Observer: start!")
	close(start)

	for range 2 {
		fmt.Println("Observer:", <-done, "a terminat")
	}
	fmt.Println("Observer: ambii roboti au terminat")
}
