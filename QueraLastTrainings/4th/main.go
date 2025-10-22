package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Comment struct
type Comment struct {
	Score int
	Text  string
}

// Flight data structure
type Flight struct {
	Name         string
	Passengers   map[string]bool    // passenger name -> exists
	Comments     map[string]Comment // passenger name -> comment
	TotalScore   int
	CommentCount int
}

// Survey struct
type Survey struct {
	Flights map[string]*Flight
	Tickets map[string]bool // "passenger:flight" -> exists (to prevent duplicates)
}

// NewSurvey creates a new survey instance
func NewSurvey() *Survey {
	return &Survey{
		Flights: make(map[string]*Flight),
		Tickets: make(map[string]bool),
	}
}

func (s *Survey) AddFlight(flightName string) error {
	if _, exists := s.Flights[flightName]; exists {
		return errors.New("flight already exists")
	}

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
	if _, exists := s.Flights[flightName]; !exists {
		return errors.New("flight does not exist")
	}

	ticketKey := passengerName + ":" + flightName
	if s.Tickets[ticketKey] {
		return errors.New("duplicate ticket")
	}

	s.Flights[flightName].Passengers[passengerName] = true
	s.Tickets[ticketKey] = true
	return nil
}

func (s *Survey) AddComment(flightName, passengerName string, comment Comment) error {
	if comment.Score < 1 || comment.Score > 10 {
		return errors.New("invalid score")
	}

	flight, exists := s.Flights[flightName]
	if !exists {
		return errors.New("flight does not exist")
	}

	if !flight.Passengers[passengerName] {
		return errors.New("passenger does not have a ticket")
	}

	if _, hasCommented := flight.Comments[passengerName]; hasCommented {
		return errors.New("duplicate comment")
	}

	flight.Comments[passengerName] = comment
	flight.TotalScore += comment.Score
	flight.CommentCount++

	return nil
}

func (s *Survey) GetCommentsAverage(flightName string) (float64, error) {
	flight, exists := s.Flights[flightName]
	if !exists {
		return 0, errors.New("flight does not exist")
	}

	if flight.CommentCount == 0 {
		return 0, errors.New("no comments")
	}

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
	flight, exists := s.Flights[flightName]
	if !exists {
		return nil, errors.New("flight does not exist")
	}

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

// Server struct
type Server struct {
	portNumber int
	survey     *Survey
}

func NewServer(port int) *Server {
	return &Server{
		portNumber: port,
		survey:     NewSurvey(),
	}
}

// Health check handler
func (server *Server) handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"Message": "OK"})
}

// Add flight handler
func (server *Server) handleAddFlight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req struct {
		Name string `json:"Name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"Message": "Invalid request"})
		return
	}

	if err := server.survey.AddFlight(req.Name); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"Message": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"Message": "OK"})
}

// Add ticket handler
func (server *Server) handleAddTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req struct {
		FlightName    string `json:"FlightName"`
		PassengerName string `json:"PassengerName"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"Message": "Invalid request"})
		return
	}

	if err := server.survey.AddTicket(req.FlightName, req.PassengerName); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"Message": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"Message": "OK"})
}

// Add comment handler
func (server *Server) handleAddComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req struct {
		FlightName    string `json:"FlightName"`
		PassengerName string `json:"PassengerName"`
		Score         int    `json:"Score"`
		Text          string `json:"Text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"Message": "Invalid request"})
		return
	}

	comment := Comment{
		Score: req.Score,
		Text:  req.Text,
	}

	if err := server.survey.AddComment(req.FlightName, req.PassengerName, comment); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"Message": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"Message": "OK"})
}

// Get comments handler
func (server *Server) handleGetComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get average query parameter (default is false if not specified)
	averageParam := r.URL.Query().Get("average")
	isAverage := averageParam == "true"

	// Check if we have a flight name in the path
	// Path format: /comments/f1 or /comments
	path := r.URL.Path
	flightName := ""

	if strings.HasPrefix(path, "/comments/") {
		flightName = strings.TrimPrefix(path, "/comments/")
		flightName = strings.TrimSpace(flightName)
	}

	if flightName != "" {
		// Single flight query
		if isAverage {
			avg, err := server.survey.GetCommentsAverage(flightName)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"Message": err.Error()})
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"Message": "OK",
				"Average": avg,
			})
		} else {
			comments, err := server.survey.GetComments(flightName)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"Message": err.Error()})
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"Message": "OK",
				"Texts":   comments,
			})
		}
	} else {
		// All flights query
		if isAverage {
			averages := server.survey.GetAllCommentsAverage()
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"Message":  "OK",
				"Averages": averages,
			})
		} else {
			allComments := server.survey.GetAllComments()
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"Message": "OK",
				"Texts":   allComments,
			})
		}
	}
}

// Start the server (blocking)
func (server *Server) Start() {
	listenAddress := fmt.Sprintf(":%d", server.portNumber)

	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" && r.Method == "GET" {
			server.handleRoot(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	mux.HandleFunc("/flights", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			server.handleAddFlight(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	mux.HandleFunc("/tickets", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			server.handleAddTicket(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	mux.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			server.handleAddComment(w, r)
		} else if r.Method == "GET" {
			server.handleGetComments(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	mux.HandleFunc("/comments/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			server.handleGetComments(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	http.ListenAndServe(listenAddress, mux)
}
