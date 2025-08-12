package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var season string
	coats := make([]string, 0)
	shirts := make([]string, 0)
	pants := make([]string, 0)
	caps := make([]string, 0)
	jackets := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	linesRead := 0

	for scanner.Scan() && linesRead < 6 {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}
		linesRead++

		if strings.Contains(line, ":") {
			parts := strings.SplitN(line, ":", 2)
			category := strings.TrimSpace(parts[0])
			colors := strings.Fields(parts[1])
			switch category {
			case "COAT":
				coats = append(coats, colors...)
			case "SHIRT":
				shirts = append(shirts, colors...)
			case "PANTS":
				pants = append(pants, colors...)
			case "CAP":
				caps = append(caps, colors...)
			case "JACKET":
				jackets = append(jackets, colors...)
			}
		} else {
			season = strings.ToUpper(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
		return
	}
	generateCombinations(season, coats, shirts, pants, caps, jackets)
}

func generateCombinations(season string, coats, shirts, pants, caps, jackets []string) {
	for _, shirt := range shirts {
		for _, pant := range pants {
			combination := map[string]string{
				"SHIRT": shirt,
				"PANTS": pant,
			}

			switch season {
			case "SUMMER":
				for _, cap := range caps {
					finalCombination := copyMap(combination)
					finalCombination["CAP"] = cap
					printCombination(finalCombination)
				}

			case "SPRING", "FALL":
				capOptions := append([]string{""}, caps...)
				coatOptions := append([]string{""}, getValidCoats(coats, season)...)

				for _, cap := range capOptions {
					for _, coat := range coatOptions {
						finalCombination := copyMap(combination)
						if cap != "" {
							finalCombination["CAP"] = cap
						}
						if coat != "" {
							finalCombination["COAT"] = coat
						}
						printCombination(finalCombination)
					}
				}

			case "WINTER":
				validCoats := getValidCoats(coats, season)

				for _, coat := range validCoats {
					finalCombination := copyMap(combination)
					finalCombination["COAT"] = coat
					printCombination(finalCombination)
				}

				for _, jacket := range jackets {
					finalCombination := copyMap(combination)
					finalCombination["JACKET"] = jacket
					printCombination(finalCombination)
				}
			}
		}
	}
}

func getValidCoats(coats []string, season string) []string {
	if season == "FALL" {
		validCoats := make([]string, 0)
		for _, coat := range coats {
			if coat != "yellow" && coat != "orange" {
				validCoats = append(validCoats, coat)
			}
		}
		return validCoats
	}
	return coats
}

func copyMap(original map[string]string) map[string]string {
	copy := make(map[string]string)
	for k, v := range original {
		copy[k] = v
	}
	return copy
}

func printCombination(combination map[string]string) {
	output := []string{}

	if coat, exists := combination["COAT"]; exists {
		output = append(output, "COAT: "+coat)
	}
	if shirt, exists := combination["SHIRT"]; exists {
		output = append(output, "SHIRT: "+shirt)
	}
	if pant, exists := combination["PANTS"]; exists {
		output = append(output, "PANTS: "+pant)
	}
	if cap, exists := combination["CAP"]; exists {
		output = append(output, "CAP: "+cap)
	}
	if jacket, exists := combination["JACKET"]; exists {
		output = append(output, "JACKET: "+jacket)
	}

	fmt.Println(strings.Join(output, " "))
}
