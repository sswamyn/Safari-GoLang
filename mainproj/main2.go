package main

import (
	"fmt"
	"strconv"
)

// constants have a "kind" i.e. integer, floating, etc.
// const has enormous range and precision--calculations are in compiler
// variables have a type
const (
	DaysInJanuary = 31 // iota is zero here
	One           = iota
	Two           = iota
	DaysInWeek    = 7
	MonthsInYear  = 12
	What          = iota
	Big           = 1 << iota
)

func main() {
	fmt.Println("Hello again")

	var small int8 = DaysInJanuary
	fmt.Println(small)
	//var medium int32 = small // NOPE, cannot assign dissimilar types (generally)
	var medium int32 = DaysInJanuary
	fmt.Println(medium)
	medium = int32(small) // type as a function..

	var num = string(1000)
	fmt.Println(num) // converted utf-8

	//var numText string = strconv.Itoa(1000)
	//var numText = strconv.Itoa(1000)
	numText := strconv.Itoa(1000) // for local var, much more common
	fmt.Println(numText)

	str1 := "He"
	str1 += "llo"
	str2 := "Hello"
	fmt.Println(str1 == str2) // equality tests "bitpattern" for most types "comparable" and things made from them

	fmt.Printf("str1 is %s or %v type is %T at %p\n", str1, str1, str1, &str1)
	fmt.Printf("str2 is %s or %v type is %T at %p\n", str2, str2, str2, &str2)

	var strptr *string = &str1
	fmt.Printf("strprt is %v type %T address: %p\n", strptr, strptr, strptr)
	//strptr = strptr + 1 // NOPE, no "arithmetic" on pointers, phew!!!

	// operators: mostly as expeted < > <= == != + - * /
	// DO NOT HAVE ? :
	// assignment does NOT have value

	//a := 99
	//var b int // b is definitely zero

	//b = (a = 100) // NOPE
	//b = a++ // ALSO NOPE

	// arrays have a dimension!!!
	ia := [4]int{} // arrays have fixed size!
	//ia := [...]int{1, 2, 3, 4}
	//ia := []int{1, 2, 3, 4} // Slice--NOT the same as an array can change "size" of slice
	fmt.Printf("ia is %#v, type is %T\n", ia, ia)
	ia[0] = 1000
	fmt.Printf("ia is %#v, type is %T\n", ia, ia)
	var ia2 [4]int
	fmt.Printf("ia2 is %#v, type is %T\n", ia2, ia2)
	ia2 = ia // assignment for arrays is OK PROVIDED, type and size is the same
	fmt.Printf("ia2 is %#v, type is %T\n", ia2, ia2)

	fmt.Printf("ia is %#v, at %p, type is %T\n", ia, &ia, ia)
	ia3 := ia // assignment (and function argument passing!) COPIES the whole value
	fmt.Printf("ia  is %#v, at %p, type is %T\n", ia, &ia, ia)
	fmt.Printf("ia3 is %#v, at %p, type is %T\n", ia3, &ia3, ia3)

}
