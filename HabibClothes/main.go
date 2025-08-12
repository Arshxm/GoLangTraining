package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// creating a new scanner class
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	// reading a line with multiple values and adding them to a list, then printing it
	scanner.Scan()
	coatInput := scanner.Text()
	coatList := strings.Fields(coatInput)[1:]

	scanner.Scan()
	shirtInput := scanner.Text()
	shirtList := strings.Fields(shirtInput)[1:]

	scanner.Scan()
	pantsInput := scanner.Text()
	pantsList := strings.Fields(pantsInput)[1:]

	scanner.Scan()
	capInput := scanner.Text()
	capsList := strings.Fields(capInput)[1:]

	scanner.Scan()
	jacketInput := scanner.Text()
	jacketList := strings.Fields(jacketInput)[1:]

	scanner.Scan()
	season := scanner.Text()
	// converting string to int (can also use .Itoa() for int to string)

	if season == "SPRING" {
		for _, coat := range coatList {
			for _, shirt := range shirtList {
				for _, pants := range pantsList {
					fmt.Printf("COAT: %s SHIRT: %s PANTS: %s\n", coat, shirt, pants)
				}
			}
		}
		for _, shirt := range shirtList {
			for _, pants := range pantsList {
				for _, caps := range capsList {
					fmt.Printf("SHIRT: %s PANTS: %s CAP: %s\n", shirt, pants, caps)
				}
			}
		}
		for _, coat := range coatList {
			for _, shirt := range shirtList {
				for _, pants := range pantsList {
					for _, caps := range capsList {
						fmt.Printf("COAT: %s SHIRT: %s PANTS: %s CAP: %s\n", coat, shirt, pants, caps)
					}
				}
			}
		}
		for _, shirt := range shirtList {
			for _, pants := range pantsList {
				fmt.Printf("SHIRT: %s PANTS: %s\n", shirt, pants)
			}
		}
	}
	if season == "SUMMER" {
		for _, shirt := range shirtList {
			for _, pants := range pantsList {
				for _, caps := range capsList {
					fmt.Printf("SHIRT: %s PANTS: %s CAP: %s\n", shirt, pants, caps)
				}
			}
		}
	}

	if season == "FALL" {
		for _, coat := range coatList {
			if coat != "orange" && coat != "yellow" {
				for _, shirt := range shirtList {
					for _, pants := range pantsList {
						fmt.Printf("COAT: %s SHIRT: %s PANTS: %s\n", coat, shirt, pants)
					}
				}
			}
		}
		for _, shirt := range shirtList {
			for _, pants := range pantsList {
				for _, caps := range capsList {
					fmt.Printf("SHIRT: %s PANTS: %s CAP: %s\n", shirt, pants, caps)
				}
			}
		}
		for _, coat := range coatList {
			if coat != "orange" && coat != "yellow" {
				for _, shirt := range shirtList {
					for _, pants := range pantsList {
						for _, caps := range capsList {
							fmt.Printf("COAT: %s SHIRT: %s PANTS: %s CAP: %s\n", coat, shirt, pants, caps)
						}
					}
				}
			}
		}
		for _, shirt := range shirtList {
			for _, pants := range pantsList {
				fmt.Printf("SHIRT: %s PANTS: %s\n", shirt, pants)
			}
		}
	}

	if season == "WINTER" {
		for _, shirt := range shirtList {
			for _, pants := range pantsList {
				for _, jacket := range jacketList {
					fmt.Printf("SHIRT: %s PANTS: %s JACKET: %s\n", shirt, pants, jacket)
				}
			}
		}
		for _, coat := range coatList {
			for _, shirt := range shirtList {
				for _, pants := range pantsList {
					fmt.Printf("COAT: %s SHIRT: %s PANTS: %s\n", coat, shirt, pants)
				}
			}
		}
	}
}