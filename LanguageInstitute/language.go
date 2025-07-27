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

	var names []string
	var averages []string

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		scanner.Scan()
		name := scanner.Text()
		names = append(names, name)

		scanner.Scan()
		scoreText := scanner.Text()
		strScores := strings.Fields(scoreText)

		var scores []int
		for _, strScore := range strScores {
			score, err := strconv.Atoi(strScore)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			scores = append(scores, score)
		}

		if len(scores) > 0 {
			avg := average(scores)
			if avg >= 80 {
				averages = append(averages, "Excellent")
			} else if avg >= 60 {
				averages = append(averages, "Very Good")
			} else if avg >= 40 {
				averages = append(averages, "Good")
			} else {
				averages = append(averages, "Fair")
			}
		} else {
			averages = append(averages, "No scores")
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(names[i], averages[i])
	}
}

func average(scores []int) int {
	sum := 0
	for _, score := range scores {
		sum += score
	}
	return sum / len(scores)
}
