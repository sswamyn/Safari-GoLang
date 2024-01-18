package main

import "fmt"

type Color struct {
	hue, sat, bright int
}

func (c Color) isRed() bool {
	return c.hue == 0
}

type Car struct {
	make string
	//hue int // this would force use of c.Color.hue for disambiguation
	//color Color
	Color // embedded struct... can only have ONE of any given type
}

func (c Car) String() string {
	return "I'm a car, of make " + c.make
}

type Colored interface {
	isRed() bool
}

type Paper struct {
	Thickness int
}

func (p Paper) isRed() bool {
	return false
}

func main() {
	//c1 := Car{color: Color{0, 100, 100}, make: "Frod"}
	c1 := Car{Color: Color{0, 100, 100}, make: "Frod"} // no shorthand syntax for initialization
	//c1 := Car{"Frod", Color{0, 100, 100}}
	fmt.Printf("c1 is %#v, type %T\n", c1, c1)
	//fmt.Println("hue is", c1.Color.hue)
	fmt.Println("hue is", c1.hue) // behaves AS IF hue belongs to Car :)

	fmt.Println("is the car red?", c1.isRed())

	stuff := []Colored{c1, Color{128, 0, 0}, Paper{}}
	for _, v := range stuff {
		fmt.Printf("Is %v red? %t\n", v, v.isRed())
	}

	var st fmt.Stringer = c1
	fmt.Println(st)
}
