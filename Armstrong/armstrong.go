package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func armstrong(s string) {
	sum := 0
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(s, -1)
	
	if len(s) == 0 {
		fmt.Printf("YES")
		return
	}
	if len(matches) == 0 {
		fmt.Printf("NO")
		return
	}

	for _, match := range matches {
		number, err := strconv.Atoi(match)
		if err != nil {
			fmt.Printf("NO")
			return
		}
		sum += number
	}

	temp := sum
	powerSum := 0
	powerTemp := 0

	for temp > 0 {
		digit := temp % 10
		powerTemp = digit
		for i := 1; i < len(strconv.Itoa(sum)); i++ {
			powerTemp *= digit
		}
		temp /= 10
		powerSum += powerTemp
	}

	if sum == powerSum {
		fmt.Printf("YES")
		return
	}
	fmt.Printf("NO")
}
func main() {
	var s string
	fmt.Scan(&s)
	armstrong(s)
}