package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanInput()
}

func scanInput() {
	var n int
	fmt.Scanf("%d", &n)

	var arithmeticSequences []int
	var names []string

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			// Handle case where there are fewer lines than expected
			fmt.Printf("Warning: Expected %d lines but only got %d\n", n, i)
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			i-- // Don't count empty lines, retry
			continue
		}

		fields := strings.Fields(line)

		if len(fields) < 2 {
			fmt.Printf("Warning: Line %d has insufficient data: %s\n", i+1, line)
			i-- // Retry this iteration
			continue
		}

		// First field is the name
		name := fields[0]
		names = append(names, name)

		// Remaining fields are fuel values
		var fuels []int
		for j := 1; j < len(fields); j++ {
			fuel, err := strconv.Atoi(fields[j])
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			fuels = append(fuels, fuel)
		}

		arithmeticSequences = append(arithmeticSequences, arithmeticSequence(fuels))
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i], arithmeticSequences[i])
	}
}

func arithmeticSequence(fuels []int) int {
	if len(fuels) < 3 {
		return 0
	}

	count := 0
	n := len(fuels)
	i := 0

	for i < n-2 {
		j := i + 1
		diff := fuels[j] - fuels[i]

		for j < n-1 && fuels[j+1]-fuels[j] == diff {
			j++
		}
		length := j - i + 1
		if length >= 3 {
			count += (length - 2) * (length - 1) / 2
		}

		i = j
	}

	return count
}
