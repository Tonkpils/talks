package main

import "fmt"

// OMIT
type Person struct {
	Age int
}

type Ninja struct {
	Languages []string
}

type Leo struct {
	Person
	Developer Ninja
}

func main() {
	leo := Leo{
		Person:    Person{Age: 28},
		Developer: Ninja{Languages: []string{"ruby", "go"}},
	}

	fmt.Printf("%+v\n", leo)
}

// END OMIT
