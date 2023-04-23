package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(2)
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Println("My favorite number is", rand.Intn(50))

}
