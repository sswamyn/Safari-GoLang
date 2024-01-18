package main

import "fmt"

type DayOfWeek int8

type Person struct {
	//Name, Nddress string
	Name    string
	Address string
}

func (p Person) Greet(handle string) string {
	return "Dear " + handle + " " + p.Name
}

// mutating functions (almost certainly) require pointers, not copies of the object!
// func (p Person) MoveHouse(newaddr string) { // not so good...
func (p *Person) MoveHouse(newaddr string) {
	fmt.Println("Before p is", p)
	p.Address = newaddr
	fmt.Println("After p is", p)
}

type Human struct {
	Name, Address string
}

func main() {
	var monday DayOfWeek = 1
	fmt.Println(monday)
	var small int8 = 3
	//monday = small // NO, different types
	monday = DayOfWeek(small) // type conversion

	//p1 := Person{}
	//var p1 Person
	//p1 := Person{"Fred", "Over the rainbow"}
	p1 := Person{Address: "Over the rainbow", Name: "Fred"}
	fmt.Printf("p1 is %#v, type %T\n", p1, p1)

	var h1 Human
	//h1 = p1 // NOPE, different types
	h1 = Human(p1) // names, types, and sequence of struct fields are identica? Coertion is permitted
	fmt.Println(h1)

	// equality tests ok if contents are comparable
	p2 := Person{Address: "Over the rainbow", Name: "Fred"}
	fmt.Println(p1 == p2)

	// can take a pointer from a literal struct
	pp := &Person{Address: "Over the rainbow", Name: "Fred"}
	fmt.Printf("pp is %T, %p, points at %#v\n", pp, pp, *pp)

	fmt.Println("Greeting for fred is:")
	fmt.Println(p1.Greet("Mr"))

	//(&p1).MoveHouse("Down the lane") // below is equivalent, we don't have to extract the pointer
	p1.MoveHouse("Down the lane")
	fmt.Println("After function call", p1)
}
