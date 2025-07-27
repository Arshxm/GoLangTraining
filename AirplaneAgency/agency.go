package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	output := []string{}

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	countryCodes := make(map[string]string)

	for i := 0; i < n; i++ {
		scanner.Scan()
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Split(line, " ")
		countryCodes[parts[1]] = parts[0]
	}

	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())

	for j := 0; j < q; j++ {
		scanner.Scan()
		line := strings.TrimSpace(scanner.Text())
		countryCode := line[:3]

		if _, exists := countryCodes[countryCode]; exists {
			output = append(output, countryCodes[countryCode])
		} else {
			output = append(output, "Invalid Number")
		}
	}

	for _, v := range output {
		fmt.Println(v)
	}
}
