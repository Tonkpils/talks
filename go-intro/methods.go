package main

import "fmt"

type Foo struct {
	Exported string
}

func (f *Foo) SetExported(val string) {
	f.Exported = val
}
func (f *Foo) PrintExported() {
	fmt.Println(f.Exported)
}

func main() {
	foo := Foo{Exported: "Foo"}
	foo.PrintExported()
	foo.SetExported("Set Foo")
	foo.PrintExported()
}
