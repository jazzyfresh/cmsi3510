package main

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// A little utility that simulates performing a task for a random duration.
// For example, calling do(10, "Remy", "is cooking") will compute a random
// number of milliseconds between 5000 and 10000, log "Remy is cooking",
// and sleep the current goroutine for that much time.

func do(seconds int, action ...any) {
	log.Println(action...)
	randomMillis := 500*seconds + rand.Intn(500*seconds)
	time.Sleep(time.Duration(randomMillis) * time.Millisecond)
}

type Order struct {
	id       uint64
	customer string
	// make a reply which is a channel that can take the order (HINT you need a pointer to the order)
	reply chan *Order
	// Also the name of the  cook
}

var nextID atomic.Uint64

// A waiter can only hold 3 orders at once
var Waiter = make(chan *Order, 3)

func Cook(name string) {
	// log that the cook is starting
	log.Println(name, "is starting")
	// loop forever
	for {
		var order *Order = <-Waiter
		log.Println(name, "is cooking order", order.id, "for", order.customer)
		// seed the random number generator with the  current time
		// generate a random integer between  0 and 99
		var r int = rand.Intn(10-5+1) + 5 // random number between 5 and 10
		time.Sleep(time.Duration(r) * time.Second)
		order.reply <- order // send the order back to the customer
	}
	//  Wait for an order from the waiter
	//  cook it
	// put the name of the cook in the order
	// send it back to the reply channel: order.reply <- order
}

func Customer(name string, wg *sync.WaitGroup, counter *int64) {
	for mealsEaten := 0; mealsEaten < 5; mealsEaten++ {
		// increment the counter
		var orderID int64 = atomic.AddInt64(counter, 1)
		// place an order
		order := &Order{id: uint64(orderID), customer: name, reply: make(chan *Order)}
		// select statement so that if the waiter gets it within 7 seconds
		// then you get it from the cook and eat it (mealsEaten++)
		select {
		case Waiter <- order:
			order = <-order.reply
		case <-time.After(7 * time.Second):
			log.Println(name, "is waiting too long, abandoning order", orderID)
		}
		//  if  you dont get leave the restaurant
		// do (5, name, "is waiting too long, abandoning order", order.id)
	}
	wg.Done()
}

func main() {
	var counter int64
	customers := [10]string{
		"Ani", "Bai", "Cat", "Dao", "Eve", "Fay", "Gus", "Hua", "Iza", "Jai",
	}
	// in a loop start each customer as a goroutine
	var wg sync.WaitGroup
	for _, customer := range customers {
		wg.Add(1)
		go Customer(customer, &wg, &counter)

	}

	// Start 3 cooks, Remy , Linguini and Colette
	go Cook("Remy")
	go Cook("Linguini")
	go Cook("Colette")
	//
	// Wait for all customers to finish
	wg.Wait()
	close(Waiter)

	log.Println("Restaurant is closing")
}
