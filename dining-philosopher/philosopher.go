package main

import (
	"fmt"
	"html"
	"strconv"
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

func (p Philosopher) Eat() {
	p.LeftChopStick.Lock()
	p.RightChopStick.Lock()

	fmt.Println(p.Name + " is eating " + GetEmoticon(FOOD))
}

func GetEmoticon(value string) string {
	i, _ := strconv.ParseInt(value, 0, 64)
	foodEmoticon := html.UnescapeString(string(i))
	return foodEmoticon
}
