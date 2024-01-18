package main

import (
	"fmt"
	"time"
)

func ticker(ch chan int) {
	for x := 0; x < 10; x++ {
		time.Sleep(time.Duration(1000) * time.Millisecond)
		fmt.Println("tick!", x)
	}
	fmt.Println("ending")
	ch <- 0
}

func main() {
	comms := make(chan int)
	go ticker(comms)

	v := <-comms
	fmt.Println("finished, v received", v)
}
