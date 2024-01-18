package main

import "fmt"

func main() {
	x := -100
	for x := 0; x < 3; x++ {
		fmt.Println("x is ", x)
	}

	fmt.Println("x again", x)
	for {
		fmt.Println("infinite")
		break
	}

	a, b := 99, 100
	fmt.Println(a, b)

	sl := []int{10, 1, 2, 3, 4, 5, 6}
	for i, v := range sl {
		if i == 4 {
			//v = 1000 // no change to slice!!!
			sl[i] = 1000
		}
		fmt.Printf("at index %d value is %d\n", i, v)
	}
	for _, v := range sl {
		fmt.Printf("value is %d\n", v)
	}
}
