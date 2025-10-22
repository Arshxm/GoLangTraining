package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Flight struct {
	Name        string
	Passengers  map[string]bool   // passenger -> exists
	Reviews     map[string]Review // passenger -> review
	TotalScore  int
	ReviewCount int
}

type Review struct {
	Passenger string
	Flight    string
	Score     int
	Comment   string
}

type Booking struct {
	Passenger string
	Flight    string
}

// stripBOM removes the UTF-8 BOM if present
func stripBOM(s string) string {
	const bom = "\ufeff"
	if strings.HasPrefix(s, bom) {
		return s[len(bom):]
	}
	if r, size := utf8.DecodeRuneInString(s); r == '\ufeff' {
		return s[size:]
	}
	return s
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read number of flights
	scanner.Scan()
	text := stripBOM(strings.TrimSpace(scanner.Text()))
	n, _ := strconv.Atoi(text)

	// Store valid flights
	flights := make(map[string]*Flight)

	for i := 0; i < n; i++ {
		scanner.Scan()
		flightName := strings.TrimSpace(scanner.Text())
		flights[flightName] = &Flight{
			Name:       flightName,
			Passengers: make(map[string]bool),
			Reviews:    make(map[string]Review),
		}
	}

	// Read number of bookings
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	// Track valid bookings to avoid duplicates
	validBookings := make(map[string]bool) // "passenger:flight" -> exists

	for i := 0; i < m; i++ {
		scanner.Scan()
		booking := strings.TrimSpace(scanner.Text())
		bookingParts := strings.Split(booking, " ")
		passenger := bookingParts[0]
		flight := bookingParts[1]

		bookingKey := passenger + ":" + flight

		// Check if flight exists
		if _, exists := flights[flight]; !exists {
			fmt.Printf("Invalid flight %s\n", flight)
			continue
		}

		// Check for duplicate booking
		if validBookings[bookingKey] {
			fmt.Printf("Duplicate ticket for %s %s\n", flight, passenger)
			continue
		}

		// Valid booking
		flights[flight].Passengers[passenger] = true
		validBookings[bookingKey] = true
	}

	// Read number of reviews
	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < q; i++ {
		scanner.Scan()
		review := strings.TrimSpace(scanner.Text())
		reviewParts := strings.SplitN(review, " ", 4) // Split into max 4 parts
		passenger := reviewParts[0]
		flight := reviewParts[1]
		score, _ := strconv.Atoi(reviewParts[2])
		comment := ""
		if len(reviewParts) > 3 {
			comment = reviewParts[3]
		}

		// Check if flight exists
		if _, exists := flights[flight]; !exists {
			fmt.Printf("Invalid flight %s\n", flight)
			continue
		}

		// Check if passenger was on this flight
		if !flights[flight].Passengers[passenger] {
			fmt.Printf("Invalid passenger for %s %s\n", flight, passenger)
			continue
		}

		// Check for duplicate review
		if _, exists := flights[flight].Reviews[passenger]; exists {
			fmt.Printf("Duplicate comment for %s by %s\n", flight, passenger)
			continue
		}

		// Valid review
		flights[flight].Reviews[passenger] = Review{
			Passenger: passenger,
			Flight:    flight,
			Score:     score,
			Comment:   comment,
		}
		flights[flight].TotalScore += score
		flights[flight].ReviewCount++

		fmt.Printf("Accepted comment for %s by %s\n", flight, passenger)
	}

	// Calculate and print average scores in sorted (lexicographic) order
	// Get flight names that have reviews and sort them
	flightNames := make([]string, 0)
	for name, flight := range flights {
		if flight.ReviewCount > 0 {
			flightNames = append(flightNames, name)
		}
	}
	sort.Strings(flightNames)

	// Print averages in sorted order
	for _, flightName := range flightNames {
		flight := flights[flightName]
		average := float64(flight.TotalScore) / float64(flight.ReviewCount)
		fmt.Printf("Average score for %s is %.2f\n", flight.Name, average)
	}
}
