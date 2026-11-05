// Implementați algoritmul de Flooding pentru grafuri tare conexe.

package main

import (
	"fmt"
	"sync"
)

type Msg struct {
	Value int
}

type Status int

const (
	NonLeader Status = iota
	Leader
)

// Mi din slide
type State struct {
	id     int
	maxID  int
	status Status
	rounds int
	diam   int
}

// functia de transformare pentru nodul i, o singura runda
func node(s *State, ins []<-chan Msg, outs []chan<- Msg, wg *sync.WaitGroup) {
	defer wg.Done()

	s.rounds++ // 1. t := t+1

	for _, in := range ins { // 2. U = ID-urile primite de la vecinii de intrare
		if m := <-in; m.Value > s.maxID { // 3. max_id := max({max_id} U U)
			s.maxID = m.Value
		}
	}

	if s.rounds == s.diam { // 4.
		if s.maxID == s.id {
			s.status = Leader
		} else {
			s.status = NonLeader
		}
	} else { // 5.
		for _, out := range outs {
			out <- Msg{Value: s.maxID}
		}
	}
}

// adj[u] = vecinii de iesire ai lui u; diam = estimarea diametrului
func GraphFlooding(ids []int, adj [][]int, diam int) int {
	n := len(ids)
	if n == 0 {
		return -1
	}

	// cate un canal pentru fiecare muchie orientata u -> v
	ins := make([][]<-chan Msg, n)
	outs := make([][]chan<- Msg, n)
	for u := 0; u < n; u++ {
		for _, v := range adj[u] {
			c := make(chan Msg, 1)
			outs[u] = append(outs[u], c)
			ins[v] = append(ins[v], c)
		}
	}

	states := make([]State, n)
	for i := range states {
		states[i] = State{id: ids[i], maxID: ids[i], diam: diam}
	}

	// trimiterea initiala: fiecare nod isi trimite id-ul vecinilor de iesire
	for i := 0; i < n; i++ {
		for _, out := range outs[i] {
			out <- Msg{Value: ids[i]}
		}
	}

	for round := 0; round < diam; round++ {
		var wg sync.WaitGroup
		wg.Add(n)
		for i := 0; i < n; i++ {
			go node(&states[i], ins[i], outs[i], &wg)
		}
		wg.Wait()
	}

	for i := range states {
		if states[i].status == Leader {
			return states[i].id
		}
	}
	return -1
}

func main() {
	ids := []int{-1, 8, 2, 11, 5}
	adj := [][]int{{1}, {2, 3}, {0}, {4}, {0, 2}}
	fmt.Println(GraphFlooding(ids, adj, len(ids)-1))
}
