package main

import (
	"fmt"
	"html"
	"strconv"
	"sync"
	"time"
)

const (
	FOOD   = "0x0001F35C"
	FINISH = "0x0001F44C"
)

type Philosopher struct {
	ID             int
	Name           string
	LeftChopStick  *ChopStick
	RightChopStick *ChopStick
	Host           *Host
}

func (p *Philosopher) Eat(wg *sync.WaitGroup) {
	wg.Add(1)

	p.LeftChopStick.Lock()
	fmt.Printf("%d locked Left chopstick: %d\n", p.ID, p.LeftChopStick.ID)
	p.RightChopStick.Lock()
	fmt.Printf("%d locked Right chopstick: %d\n", p.ID, p.RightChopStick.ID)

	p.Host.requestChannel <- p

	fmt.Printf("%d is eating %s\n", p.ID, GetEmoticon(FOOD))
	time.Sleep(time.Millisecond)
	fmt.Printf("%d is done eating %s\n", p.ID, GetEmoticon(FINISH))

	p.LeftChopStick.Unlock()
	fmt.Printf("%d unlocked Left chopstick: %d\n", p.ID, p.LeftChopStick.ID)
	p.RightChopStick.Unlock()
	fmt.Printf("%d unlocked Right chopstick: %d\n", p.ID, p.RightChopStick.ID)

	wg.Done()
}

func GetEmoticon(value string) string {
	i, _ := strconv.ParseInt(value, 0, 64)
	foodEmoticon := html.UnescapeString(string(i))
	return foodEmoticon
}
