package main

import "fmt"

// OMIT
type Stringer interface {
	String() string
}

func Print(s Stringer) { fmt.Println(s.String()) }

type Foo struct{}

func (f *Foo) String() string { return "I am Foo" }

type Bar struct{}

func (b *Bar) String() string { return "I am Bar" }

func main() {
	foo := new(Foo)
	bar := &Bar{}

	Print(foo)
	Print(bar)
}

// END OMIT
