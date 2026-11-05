// Implementați algoritmul de Flooding pentru un inel unidirecțional.

package main

import (
	"fmt"
	"math"
	"sync"
)

type Msg struct {
	Value int
	Done  bool // liderul anunta ca a fost ales
}

const NULL = math.MinInt

// realizeaza operatiile interne nodului
func node(id int, in <-chan Msg, out chan<- Msg, isLeader *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	recvID := (<-in).Value // 2. recv(recv_id, stanga)

	sendID := NULL
	if recvID > id { // 3.
		sendID = recvID
	} else if recvID == id { // 4.
		*isLeader = true
		sendID = id
	} else { // 5. recvID < id
		sendID = NULL
	}

	out <- Msg{Value: sendID} // 1. send(send_id, dreapta), pentru runda urmatoare
}

// returneaza id-ul liderului
func RingFlooding(ids []int) int {
	n := len(ids)
	if n == 0 {
		return -1
	}

	chans := make([]chan Msg, n)
	for i := range chans {
		chans[i] = make(chan Msg, n)
	}

	leaders := make([]bool, n)

	// runda 0: fiecare nod isi trimite id-ul spre dreapta
	for i := 0; i < n; i++ {
		chans[(i+1)%n] <- Msg{Value: ids[i]}
	}

	// n runde, cu bariera (wg.Wait) la finalul fiecareia
	for round := 0; round < n; round++ {
		var wg sync.WaitGroup
		wg.Add(n)
		for i := 0; i < n; i++ {
			go node(ids[i], chans[i], chans[(i+1)%n], &leaders[i], &wg)
		}
		wg.Wait()
	}

	for i := 0; i < len(leaders); i++ {
		if leaders[i] == true {
			return ids[i]
		}
	}
	return -1
}

func main() {
	fmt.Println(RingFlooding([]int{-1, 8, 2, 11, 5}))
}
