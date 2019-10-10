// Go's structs are typed collections of fields. They're useful
// for grouping data together to form records.
package main

import "fmt"

type person struct {
	name string
	age  int
}

func NewPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Printf("%v\n", person{"Bob", 20})
	fmt.Printf("%v\n", person{name: "Bob", age: 20})
	fmt.Printf("%v\n", person{name: "Fred"})
	fmt.Printf("%v\n", &person{name: "Sharon", age: 45})
	fmt.Printf("%v\n", NewPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
