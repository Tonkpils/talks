package main

import "fmt"

func main() {
	// REGULAR OMIT
	var arr []int

	for i := 0; i < 5; i++ {
		arr = append(arr, i)
	}

	fmt.Println(arr)

	var b bool
	for b == false {
		fmt.Println("b is false")
		b = true
	}

	fmt.Println("For each")
	for index, value := range arr {
		fmt.Printf("%d: %d ", index, value)
	}

	// REGULAR END OMIT
}
