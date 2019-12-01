package main

import "fmt"

type hotdog int

// struct
type person struct {
	// lowercase inside the package
	fname string
	lname string
	age   int
}

type secretAgent struct {
	person
	licenceToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says "Shaken not stirred!"`)
}

// if the type has the metnod it implements the interface
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	// wacky type
	var t hotdog
	t = 7
	fmt.Println(t)

	// slice
	v := []int{1, 2, 3, 4, 5}
	fmt.Println(v)

	// map
	m := map[string]int{
		"John":    42,
		"Lyndsey": 41,
	}

	fmt.Println(m)

	p := person{
		"John",
		"Spurgin",
		42,
	}
	fmt.Println(p)
	p.speak()

	sa := secretAgent{
		person{
			"some",
			"person",
			45,
		},
		false,
	}

	sa.speak()
	sa.person.speak()

	// polymorphism
	saySomething(p)
	saySomething(sa)

}
