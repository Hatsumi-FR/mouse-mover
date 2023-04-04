package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Your mouse will move every 5s")
	for {
		robotgo.MoveMouse(rand.Intn(1000), rand.Intn(1000))
		time.Sleep(5 * time.Second)
	}
}
