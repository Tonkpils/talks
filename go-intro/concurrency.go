package main

import (
	"fmt"
	"math/rand"
	"time"
)

// OMIT
func TryMe(ch chan time.Duration) {
	for {
		i := time.Duration(rand.Intn(1000)) * time.Millisecond
		time.Sleep(i)
		ch <- i
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	ch := make(chan time.Duration)
	go TryMe(ch)

	for {
		fmt.Printf("I waited %s\n", <-ch)
	}
}

// END OMIT
