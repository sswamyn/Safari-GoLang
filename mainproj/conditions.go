package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// conditions, NO ? :
	// if/else & switch

	x := 99
	if x := rand.Intn(200); x > 100 { // curlies required!
		fmt.Println("It's big", x)
	} else {
		fmt.Println("it's not so big", x)
	}
	fmt.Println("x is", x)

	// if cases are constants, they must be distinct
	switch x := rand.Intn(20) + 80; x {
	case 98:
		fmt.Println("it's 98")
	case 99:
		fmt.Println("99")
		//fallthrough // :)
	case 100:
		fmt.Println("100")
	default:
		fmt.Println("something else")
	}
	fmt.Println("x is", x)

	switch 3 {
	case getVal():
		fmt.Println("it matched!")
	case getVal():
		fmt.Println("match 2")
	case getVal():
		fmt.Println("match again")
	default:
		fmt.Println("no match")
	}

	switch x := getVal(); /* blank? value is true */ {
	case x == 0:
		fmt.Println("x is zero")
	case x == 1:
		fmt.Println("x is 1")
	case x == 2:
		fmt.Println("x is 2")
	case x == 3:
		fmt.Println("x is 3")
	}
}

func getVal() int {
	return rand.Intn(4)
}
