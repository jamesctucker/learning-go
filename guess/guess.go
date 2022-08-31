package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	res := genRandomNumber()
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println(res)
}

func genRandomNumber() int {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1

	return target
}
