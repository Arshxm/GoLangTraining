package main

import (
	"fmt"
	"math"
)

func main() {
	scanInput()
}

func scanInput() {
	var p int
	fmt.Scanf("%d", &p)

	var sumTax int = 120
	p -= 1000
	tax := int(math.Floor(float64(p) * 0.2))
	sumTax += tax
	fmt.Println(sumTax)
}
