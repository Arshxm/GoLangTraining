package main

import (
	"errors"
	"sync"
)

type Comment struct {
	// do not modify or remove these fields
	Score int
	Text  string
	// but you can add anything you want
}

type FlightData struct {
	Name         string
	Passengers   map[string]bool    // passenger name -> exists
	Comments     map[string]Comment // passenger name -> comment
	TotalScore   int
	CommentCount int
}

type Survey struct {
	mu      sync.RWMutex
	flights map[string]*FlightData
	tickets map[string]bool // "passenger:flight" -> exists (to prevent duplicates)
}

func NewSurvey() *Survey {
	return &Survey{
		flights: make(map[string]*FlightData),
		tickets: make(map[string]bool),
	}
}

func (s *Survey) AddFlight(flightName string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if flight already exists
	if _, exists := s.flights[flightName]; exists {
		return errors.New("flight already exists")
	}

	// Add new flight
	s.flights[flightName] = &FlightData{
		Name:         flightName,
		Passengers:   make(map[string]bool),
		Comments:     make(map[string]Comment),
		TotalScore:   0,
		CommentCount: 0,
	}
	return nil
}

func (s *Survey) AddTicket(flightName, passengerName string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if flight exists
	if _, exists := s.flights[flightName]; !exists {
		return errors.New("flight does not exist")
	}

	// Check for duplicate ticket
	ticketKey := passengerName + ":" + flightName
	if s.tickets[ticketKey] {
		return errors.New("duplicate ticket")
	}

	// Add ticket
	s.flights[flightName].Passengers[passengerName] = true
	s.tickets[ticketKey] = true
	return nil
}

func (s *Survey) AddComment(flightName, passengerName string, comment Comment) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Validate score range (1-10)
	if comment.Score < 1 || comment.Score > 10 {
		return errors.New("invalid score")
	}

	// Check if flight exists
	flight, exists := s.flights[flightName]
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
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Check if flight exists
	flight, exists := s.flights[flightName]
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
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make(map[string]float64)

	for flightName, flight := range s.flights {
		if flight.CommentCount > 0 {
			average := float64(flight.TotalScore) / float64(flight.CommentCount)
			result[flightName] = average
		}
	}

	return result
}

func (s *Survey) GetComments(flightName string) ([]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Check if flight exists
	flight, exists := s.flights[flightName]
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
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make(map[string][]string)

	for flightName, flight := range s.flights {
		comments := make([]string, 0, len(flight.Comments))
		for _, comment := range flight.Comments {
			comments = append(comments, comment.Text)
		}
		result[flightName] = comments
	}

	return result
}
