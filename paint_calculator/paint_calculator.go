package main

import (
	"fmt"
	"log"
)

func main() {
	var total float64
	amount, err := calculatePaintNeeded(4.2, 3.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount

	amount, err = calculatePaintNeeded(5.2, 3.5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount
	fmt.Printf("Total: %.2f liters needed\n", total)
}

func calculatePaintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}

	area := width * height
	return area / 10.0, nil
}