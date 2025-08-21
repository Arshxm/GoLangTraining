package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Server struct {
	port string
}

type Book struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	IsBorrowed bool
}

type Response struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BorrowRequest struct {
	Borrow bool `json:"borrow"`
}

var books = make(map[string]Book)

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) Start() {
	http.HandleFunc("/book", s.bookHandler)
	http.ListenAndServe(":"+s.port, nil)
}

func (s *Server) bookHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodGet:
		s.getBooks(w, r)
	case http.MethodPost:
		s.addBook(w, r)
	case http.MethodDelete:
		s.deleteBook(w, r)
	case http.MethodPut:
		s.borrowBook(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) getBooks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	title := strings.ToLower(query.Get("title"))
	author := strings.ToLower(query.Get("author"))

	for _, book := range books {
		if title != "" && author != "" {
			if strings.ToLower(book.Title) == title && strings.ToLower(book.Author) == author && !book.IsBorrowed {
				resp := Response{
					Title:  book.Title,
					Author: book.Author,
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(resp)
				return
			}
		}
	}
}

func (s *Server) addBook(w http.ResponseWriter, r *http.Request) {
	book := Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	book.Title = strings.ToLower(book.Title)
	book.Author = strings.ToLower(book.Author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, b := range books {
		if strings.ToLower(b.Title) == book.Title && strings.ToLower(b.Author) == book.Author {
			http.Error(w, "Book already exists", http.StatusBadRequest)
			return
		}

	}
	book.IsBorrowed = false
	books[book.Title] = book
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (s *Server) deleteBook(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	title := strings.ToLower(query.Get("title"))
	author := strings.ToLower(query.Get("author"))
	for _, book := range books {
		if strings.ToLower(book.Title) == title && strings.ToLower(book.Author) == author {
			delete(books, book.Title)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func (s *Server) borrowBook(w http.ResponseWriter, r *http.Request) {
	var req BorrowRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := r.URL.Query()
	title := strings.ToLower(query.Get("title"))
	author := strings.ToLower(query.Get("author"))
	for _, book := range books {
		if strings.ToLower(book.Title) == title && strings.ToLower(book.Author) == author {
			if !book.IsBorrowed && req.Borrow {
				book.IsBorrowed = true
				return
			}
			if book.IsBorrowed && !req.Borrow {
				book.IsBorrowed = false
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
			break
		}
	}
}
