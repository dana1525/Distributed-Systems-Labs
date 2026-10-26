// Scrieți un program Go în care un goroutine să notifice
// funcția main() atunci când procesarea sa se sfârșește,
// fără a folosi un wait group.

package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	go func() {
		defer close(done) // semnaleaza sfarsitul procesarii

		fmt.Println("Start processing...")
		time.Sleep(time.Second)
		fmt.Println("Stop processing")
	}()

	<-done // main se blocheaza pana channel-ul e inchis
	fmt.Println("Main: Received notification")
}
