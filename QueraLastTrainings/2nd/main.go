package main

import (
	"errors"
)

type Comment struct {
	// do not modify or remove these fields
	Score int
	Text  string
	// but you can add anything you want
}

type Flight struct {
	Name         string
	Passengers   map[string]bool    // passenger name -> exists
	Comments     map[string]Comment // passenger name -> comment
	TotalScore   int
	CommentCount int
}

type Survey struct {
	Flights map[string]*Flight
	Tickets map[string]bool // "passenger:flight" -> exists (to prevent duplicates)
}

func NewSurvey() *Survey {
	return &Survey{
		Flights: make(map[string]*Flight),
		Tickets: make(map[string]bool),
	}
}

func (s *Survey) AddFlight(flightName string) error {
	// Check if flight already exists
	if _, exists := s.Flights[flightName]; exists {
		return errors.New("flight already exists")
	}

	// Add new flight
	s.Flights[flightName] = &Flight{
		Name:         flightName,
		Passengers:   make(map[string]bool),
		Comments:     make(map[string]Comment),
		TotalScore:   0,
		CommentCount: 0,
	}
	return nil
}

func (s *Survey) AddTicket(flightName, passengerName string) error {
	// Check if flight exists
	if _, exists := s.Flights[flightName]; !exists {
		return errors.New("flight does not exist")
	}

	// Check for duplicate ticket
	ticketKey := passengerName + ":" + flightName
	if s.Tickets[ticketKey] {
		return errors.New("duplicate ticket")
	}

	// Add ticket
	s.Flights[flightName].Passengers[passengerName] = true
	s.Tickets[ticketKey] = true
	return nil
}

func (s *Survey) AddComment(flightName, passengerName string, comment Comment) error {
	// Validate score range (1-10)
	if comment.Score < 1 || comment.Score > 10 {
		return errors.New("invalid score")
	}

	// Check if flight exists
	flight, exists := s.Flights[flightName]
	if !exists {
		return errors.New("flight does not exist")
	}

	// Check if passenger has a ticket for this flight
	if !flight.Passengers[passengerName] {
		return errors.New("passenger does not have a ticket")
	}

	// Check if passenger already commented on this flight
	if _, hasCommented := flight.Comments[passengerName]; hasCommented {
		return errors.New("duplicate comment")
	}

	// Add comment
	flight.Comments[passengerName] = comment
	flight.TotalScore += comment.Score
	flight.CommentCount++

	return nil
}

func (s *Survey) GetCommentsAverage(flightName string) (float64, error) {
	// Check if flight exists
	flight, exists := s.Flights[flightName]
	if !exists {
		return 0, errors.New("flight does not exist")
	}

	// Return error if no comments
	if flight.CommentCount == 0 {
		return 0, errors.New("no comments")
	}

	// Calculate average
	average := float64(flight.TotalScore) / float64(flight.CommentCount)
	return average, nil
}

func (s *Survey) GetAllCommentsAverage() map[string]float64 {
	result := make(map[string]float64)

	for flightName, flight := range s.Flights {
		if flight.CommentCount > 0 {
			average := float64(flight.TotalScore) / float64(flight.CommentCount)
			result[flightName] = average
		}
	}

	return result
}

func (s *Survey) GetComments(flightName string) ([]string, error) {
	// Check if flight exists
	flight, exists := s.Flights[flightName]
	if !exists {
		return nil, errors.New("flight does not exist")
	}

	// Collect all comment texts
	comments := make([]string, 0, len(flight.Comments))
	for _, comment := range flight.Comments {
		comments = append(comments, comment.Text)
	}

	return comments, nil
}

func (s *Survey) GetAllComments() map[string][]string {
	result := make(map[string][]string)

	for flightName, flight := range s.Flights {
		comments := make([]string, 0, len(flight.Comments))
		for _, comment := range flight.Comments {
			comments = append(comments, comment.Text)
		}
		result[flightName] = comments
	}

	return result
}
