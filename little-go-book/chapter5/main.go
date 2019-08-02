package main

import "fmt"

func main() {
	goku := &Saiyan{
		"Goku",
		9000,
	}
	goku.Super()
	fmt.Println(goku.Power)
}

func (s *Saiyan) Super() {
	s.Power += 10000
}

type Saiyan struct {
	Name  string
	Power int
}

func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}
