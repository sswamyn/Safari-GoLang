package main

import "fmt"

// func add(a int, b int) int {
// func add(a, b int, c string) int {
// args passed BY VALUE (copy the entire thing!) might not be bad avoid GC in most cases
// if you pass out a pointer, "escape analysis" moves that thing to the heap and GC takes over the lifetime control
func add(a, b int) int {
	fmt.Printf("a is at %p\n", &a)
	return a + b
}

func main() {
	x := 99
	y := 100
	fmt.Printf("x is at %p\n", &x)
	fmt.Println(add(x, y))
}
