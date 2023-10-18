package main

import (
	"fmt"
	"sync"
)

const NUM_PHILOSOPHERS int = 5
const NUM_CHOPSTICKS int = 5
const EAT_TIMES int = 3
const NUM_EATING_PHILOSOPHER int = 2

type Host struct {
	// channel for allowing a philosopher to eat
	requestChannel chan *Philosopher

	// channel that quits the program; tells the host
	// to stop hosting
	quitChannel chan int

	// keep track/bookkeeping of currently eating philosophers
	eatingPhilosophers map[int]bool

	// we need to lock the bookkeeping variable
	mu sync.Mutex
}

func (h *Host) manage() {
	for {
		if len(h.requestChannel) == NUM_EATING_PHILOSOPHER {
			finished := <-h.requestChannel // Pops a Philosopher object off the channel
			currentlyEating := make([]int, 0, NUM_PHILOSOPHERS)
			for index, eating := range h.eatingPhilosophers {
				if eating {
					currentlyEating = append(currentlyEating, index)
				}
			}
			fmt.Printf("%v have been eating, clearing plates %d\n", currentlyEating, finished.ID)

			h.eatingPhilosophers[finished.ID] = false
		}

		// similar to a switch stmt
		select {
		case <-h.quitChannel:
			// when the channel receives a signal
			// end the host manage function
			fmt.Println("party is over")
			return
		default:
		}
	}

}

func main() {
	var wg sync.WaitGroup
	requestChannel := make(chan *Philosopher, NUM_EATING_PHILOSOPHER)
	quitChannel := make(chan int, 1)
	host := Host{
		requestChannel: requestChannel,
		quitChannel: quitChannel,
		eatingPhilosophers: make(map[int]bool)
	}

	// make chopsticks
	// TODO
	for i := range(

	// make philos9oophers
	philos := make([]*Philosopher, NUM_PHILOSOPHERS)

	// for i := 0; i < NUM_PHILOSOPHERS; i++ {
	for i := range(NUM_PHILOSOPHERS) {
		philos[i] = &Philosopher{
			ID: i + 1,
			Name: "",
			LeftChopStick: chopsticks[i],
			RightChopStick: chopsticks[(i+1)%5],
			Host: &host
		}
	}

	go host.manage()


	for i in philos {
		go philosopher[i].eat(&wg)
	} 


	wg.Wait()
	host.quitChannel <- 1

	<-host.quitChannel



}

