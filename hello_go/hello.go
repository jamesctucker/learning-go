package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	day, month, year := now.Day(), now.Month(), now.Year()

	fmt.Println("Hello, world! Today is", month, day, year)
}
