// Trei goroutine își trimit mesaje între ele prin canale. Ele își trimit mesaje conform topologiei inel
// P1 -> P2 -> P3 -> P1. Fiecare dintre aceste evenimente este înregistrat folosind ceasuri logice
// Lamport. Pornind de la șablonul de mai jos, implementați cele 3 metode ale ceasului și stocați în
// main evenimentele ordonate crescător. Puteți folosi sort.Slice cu funcția less. Între evenimente,
// contoarele ceasurilor sunt incrementate aleator.

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type LamportClock struct{ counter int }

// eveniment intern: Ci = Ci + 1
func (l *LamportClock) Tick() {
	l.counter++
}

// incrementeaza si intoarce ts(m) = Ci (dupa incrementare)
func (l *LamportClock) Send() int {
	l.counter++
	return l.counter
}

// la primire, Cj = max(Cj, ts(m)), apoi Cj = Cj + 1
func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.counter++
}

type Message struct{ sender, ts int }

type Event struct {
	process          int
	internal_counter int
	kind             string // "internal", "send", "recv"
}

// retuneaza true daca a precede b
// la egalitate, ordonam dupa proces
func less(a, b Event) bool {
	if a.internal_counter != b.internal_counter {
		return a.internal_counter < b.internal_counter
	}
	return a.process < b.process
}

const N = 3         // nr de procese
const ROUNDS = 3    // nr rounds (runda = ce face fiecare proces intr-o iteratie a buclei (intern, send, receive))
const NO_EVENTS = 2 // numar maxim de evenimente interne pe runda

func process(id int, inbox <-chan Message, next chan<- Message, out chan<- Event, wg *sync.WaitGroup) {
	defer wg.Done()
	clock := &LamportClock{}

	for range ROUNDS {

		// realizeaza intre 0 si 2 evenimente interne (numarul se alege o data pe runda)
		n := rand.Intn(NO_EVENTS + 1)
		for range n {
			clock.Tick()
			out <- Event{id, clock.counter, "internal"}
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		}

		// send
		ts := clock.Send()
		out <- Event{id, ts, "send"}
		next <- Message{sender: id, ts: ts}

		// receive
		msg := <-inbox
		clock.Receive(msg.ts)
		out <- Event{id, clock.counter, "recv"}
	}
}

func main() {
	var chans [N]chan Message
	for i := range chans {
		chans[i] = make(chan Message, ROUNDS)
	}

	// canal pentru evenimente, destul de mare cat sa nu blocheze procesele
	out := make(chan Event, N*ROUNDS*(NO_EVENTS+2))

	var wg sync.WaitGroup
	for i := 0; i < N; i++ {
		wg.Add(1)
		go process(i+1, chans[i], chans[(i+1)%N], out, &wg)
	}
	wg.Wait()
	close(out)

	// main stocheaza evenimentele
	var events []Event
	for e := range out {
		events = append(events, e)
	}

	sort.Slice(events, func(i, j int) bool { // ordoneaza events conform lui less (primeste 2 indici, nu evenimentele in sine)
		return less(events[i], events[j])
	})

	for _, e := range events {
		fmt.Printf("ts=%2d  P%d  %s\n", e.internal_counter, e.process, e.kind)
	}
}
