package main

import "fmt"

func main() {
	calculatePaintNeeded(4.2, 3.0)
	calculatePaintNeeded(5.2, 3.5)
}

func calculatePaintNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)
}