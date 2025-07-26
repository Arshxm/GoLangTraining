package main

import (
	"fmt"
	"strings"
)

func main() {
	scanInput()
}

func scanInput() string {

	var p, q int
	fmt.Scanf("%d %d", &p, &q)

	for i := 1; i <= q; i++ {
		if i%p == 0 {
			fmt.Println(strings.Repeat("Hope ", i/p))
		} else {
			fmt.Println(i)
		}
	}
	return "Hope"
}
