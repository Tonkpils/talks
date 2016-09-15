package main

import "fmt"

type Foo struct {
	Exported string
}

func main() {
	foo := Foo{Exported: "Foo"}
	baz := foo
	baz.Exported = "BarBaz"

	fmt.Println(foo.Exported)
	fmt.Println(baz.Exported)
}
