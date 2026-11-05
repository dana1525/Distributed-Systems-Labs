// Un client cu ceasul decalat cere de un server timpul corect printr-o rețea cu întârzieri simulate
// cu time.Sleep. Pe lângă θ, vom folosi și durata round-trip-ului, fără timpul de procesare al
// serverului:
// δ = (T4 − T1) − (T3 − T2)

// (a) Completați funcția estimate folosind Algoritmul lui Cristian. Simulați apoi în main o
// cerere și un răspuns folosind funcția measure și verificați dacă offset-ul inițial a scăzut.

// (b) Repetați simularea de 10 ori și calculați δ și θ minime. Comparați rezultatele cu cele
// obținute la punctul a).

// --------------------------------------------------------------

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// stocheaza deviatia fata de timpul real
type Clock struct{ offset time.Duration }

func (c *Clock) Now() time.Time { return time.Now().Add(c.offset) }

func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2 // decalajul
	delta = t4.Sub(t1) - t3.Sub(t2)       // durata totala a drumului prin retea (eroare theta <= delta / 2)
	return theta, delta
}

func measure(c_clock, s_clock *Clock) (theta, delta time.Duration) {
	// Request
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)

	t2 := s_clock.Now()
	// Procesare intre 5 si 20 ms
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)

	t3 := s_clock.Now()
	// Response
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)

	t4 := c_clock.Now()
	return estimate(t1, t2, t3, t4)
}

func main() {

	server := &Clock{offset: 0}                      // serverul are timpul corect
	client := &Clock{offset: 500 * time.Millisecond} // clientul e decalat cu +500ms

	// a) o singura masuratoare
	initial := client.offset
	theta, delta := measure(client, server)
	client.offset += theta // actualizare offset client (Cristian)

	fmt.Printf("offset initial: %v\n", initial)
	fmt.Printf("theta estimat: %v, delta: %v\n", theta, delta)
	fmt.Printf("offset ramas: %v\n", client.offset)

	// b) 10 masuratori
	client.offset = initial
	minTheta, minDelta := measure(client, server)

	for i := 1; i < 10; i++ {
		th, de := measure(client, server)
		if de < minDelta {
			minDelta = de
			minTheta = th // theta de la cel mai mic delta
		}
	}

	fmt.Println("theta minim:", minTheta)
	fmt.Println("delta minim:", minDelta)

	// comparatie cu a)
	fmt.Println("delta a):", delta, "| delta minim b):", minDelta)
	fmt.Println("theta a):", theta, "| theta minim b):", minTheta)
}
