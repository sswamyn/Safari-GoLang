package main

import (
	"fmt"
	"math/rand"
)

// func getTwo(a, b, c int) (x int, y int) {
// func getTwo(a, b, c int) (x, y int) {
func getTwo(a, b, c int) (int, int) {
	return a, c // this might be more readable :)
	//x = a
	//y = c
	//return // returns whatever values x, y currently contain
}

// ALL allocated storage in Go gets a zero-like value
func sum(vals ...int) (res int) { // arrives as a slice
	fmt.Printf("received %#v, type %T\n", vals, vals)
	for _, v := range vals {
		res += v
	}
	return
}

func mightFail() (value int, error string) {
	success := rand.Intn(2) > 0

	if success {
		return 10, ""
	} else {
		return -1, "That broke"
	}
}

func filter(crit func(int) bool, vals ...int) []int {
	res := make([]int, 0, len(vals))
	// v is overwritten each time round the loop! take care if you catch a closure on it!!!
	for _, v := range vals {
		if crit(v) {
			res = append(res, v)
		}
	}
	return res
}

// functions in Go can create "closures" as you would expect (they are mutable values)
// TAKE CARE WITH CLOSURES, external variables have the JavaScript "var" effect!!!
// solution in Go is like the old solution in JS; copy the external variable to a NEW one INSIDE the loop
func main() {
	a, c := getTwo(1, 2, 3)
	fmt.Println(a, c)

	fmt.Println("sum of 1..10 is", sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	//oneToTen := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // not spreadable
	oneToTen := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // slice is spreadable
	fmt.Println("again sum of 1..10 is", sum(oneToTen...))

	adder := sum // "value" of the function, rather than result

	fmt.Println(adder(1, 2, 3, 4))

	isEven := func(x int) bool { // local func must NOT have a name :)
		return x%2 == 0
	}
	evens := filter(isEven, oneToTen...)
	fmt.Println(evens)

	val, err := mightFail()
	if err == "" { // have to test result every time
		fmt.Println(val) // happy path
	} else {
		// it broke
	}
}
