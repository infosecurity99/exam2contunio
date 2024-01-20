package main

import (
	"fmt"
	"math/rand"
	"time"
)

//function1
func randomsonbers(ch chan int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {

		randomsonber := rand.Intn(21) + 10

		ch <- randomsonber

		time.Sleep(time.Millisecond * 500)

	}

	close(ch)
}


func numbersquer(enter, resualtChanells chan int) {
	for son := range enter {

		kv := son * son

		resualtChanells <- kv

		time.Sleep(time.Millisecond * 500)
	}
	close(resualtChanells)
}

func main() {

	channel1 := make(chan int)

	channel2 := make(chan int)

	go randomsonbers(channel1)

	go numbersquer(channel1, channel2)

	for kv := range channel2 {
		fmt.Printf("kvadrati soning----%d\n", kv)
	}
}
