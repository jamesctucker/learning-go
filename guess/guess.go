package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	target := genRandomNumber()
	guess(target)
}

func genRandomNumber() int {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	return target
}

func guess(target int) {
	for guesses := 0; guesses < 10; guesses++ {
		if guesses > 0 {
			remaining := 10 - guesses
			fmt.Println("You have", remaining, "left.")
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops! Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops! Your guess was HIGH.")
		} else {
			fmt.Println("Good job! You guessed it!")
			break
		}
	}
}
