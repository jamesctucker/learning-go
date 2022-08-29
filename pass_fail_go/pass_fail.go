// pass_fail reports whether a grade is passing or failing
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	match, _ := regexp.MatchString("^[0-9]*$", input)
	if !match {
		fmt.Println("Sorry, you can only enter numeric digits. Please try again.")
		// TODO: figure out why the initial input is still stored even after recursively running the function again and setting input to a new value
		reader.Reset(reader)
		main()		
	}

	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}