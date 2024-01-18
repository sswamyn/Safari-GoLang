package main

import (
	"fmt"
	"math/rand"
)

func MightFail() {
	if rand.Intn(4) > 2 {
		fmt.Println("Failing!!!!")
		panic("Blew up")
	}
	fmt.Println("Succeeded")
}

func getTxt() string {
	fmt.Println("in getTxt")
	return "message!"
}

func callIt() {
	fmt.Println("starting")
	defer func(x string) {
		fmt.Println("executing deferred function")
		thing := recover()
		if thing == nil {
			fmt.Println("normal exit from function")
		} else {
			fmt.Println("recovered from a panic", thing)
		}
	}(getTxt())
	fmt.Println("deferral configured")
	MightFail()
	fmt.Println("returned...")

}

func main() {
	callIt()
}
