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

type GetResponse struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
type BorrowResponse struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}
type ResponseRes struct {
	Result string `json:"result"`
	Error  string `json:"error"`
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
	// Initialize the books map if it's nil
	if books == nil {
		books = make(map[string]Book)
	}
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
	title := strings.ToLower(strings.TrimSpace(query.Get("title")))
	author := strings.ToLower(strings.TrimSpace(query.Get("author")))

	if title == "" || author == "" {
		resp := ResponseRes{
			Result: "",
			Error:  "title or author cannot be empty",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Look for the specific book
	found := false
	for _, book := range books {
		if strings.ToLower(book.Title) == title && strings.ToLower(book.Author) == author {
			found = true
			if book.IsBorrowed {
				resp := ResponseRes{
					Result: "",
					Error:  "this book is borrowed",
				}
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(resp)
				return
			}
			resp := GetResponse{
				Title:  book.Title,
				Author: book.Author,
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
			return
		}
	}

	// Book not found
	if !found {
		resp := ResponseRes{
			Result: "",
			Error:  "this book does not exist",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}
}

func (s *Server) addBook(w http.ResponseWriter, r *http.Request) {
	book := Book{}

	// Check content type to handle both JSON and form data
	contentType := r.Header.Get("Content-Type")

	if strings.Contains(contentType, "application/json") {
		// Handle JSON data
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		// Handle form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		book.Title = r.FormValue("title")
		book.Author = r.FormValue("author")
	}

	// Validate input before converting to lowercase
	if strings.TrimSpace(book.Title) == "" || strings.TrimSpace(book.Author) == "" {
		resp := ResponseRes{
			Result: "",
			Error:  "title or author cannot be empty",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	book.Title = strings.ToLower(strings.TrimSpace(book.Title))
	book.Author = strings.ToLower(strings.TrimSpace(book.Author))

	// Check for duplicate books
	for _, b := range books {
		if strings.ToLower(b.Title) == book.Title && strings.ToLower(b.Author) == book.Author {
			resp := ResponseRes{
				Result: "this book is already in the library",
				Error:  "",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
			return
		}
	}

	resp := ResponseRes{
		Result: "added book " + book.Title + " by " + book.Author,
		Error:  "",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	book.IsBorrowed = false
	books[book.Title] = book
	return
}

func (s *Server) deleteBook(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	title := strings.ToLower(strings.TrimSpace(query.Get("title")))
	author := strings.ToLower(strings.TrimSpace(query.Get("author")))

	if title == "" || author == "" {
		resp := ResponseRes{
			Result: "",
			Error:  "title or author cannot be empty",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	for _, book := range books {
		if strings.ToLower(book.Title) == title && strings.ToLower(book.Author) == author {
			delete(books, book.Title)
			resp := ResponseRes{
				Result: "successfully deleted",
				Error:  "",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
			return
		}
	}
	resp := ResponseRes{
		Result: "",
		Error:  "this book does not exist",
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(resp)
}

func (s *Server) borrowBook(w http.ResponseWriter, r *http.Request) {
	var req BorrowRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := ResponseRes{
			Result: "",
			Error:  "borrow value cannot be empty",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	query := r.URL.Query()
	title := strings.ToLower(strings.TrimSpace(query.Get("title")))
	author := strings.ToLower(strings.TrimSpace(query.Get("author")))

	if title == "" || author == "" {
		resp := ResponseRes{
			Result: "",
			Error:  "title or author cannot be empty",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	for bookKey, book := range books {
		if strings.ToLower(book.Title) == title && strings.ToLower(book.Author) == author {
			if !book.IsBorrowed && req.Borrow {
				book.IsBorrowed = true
				books[bookKey] = book // Important: Update the book in the map
				resp := BorrowResponse{
					Result: "you have borrowed this book successfully",
					Error:  "",
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(resp)
				return
			}
			if book.IsBorrowed && req.Borrow {
				resp := ResponseRes{
					Result: "",
					Error:  "this book is already borrowed",
				}
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(resp)
				return
			}
			if !book.IsBorrowed && !req.Borrow {
				resp := ResponseRes{
					Result: "",
					Error:  "this book is already in the library",
				}
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(resp)
				return
			}
			if book.IsBorrowed && !req.Borrow {
				book.IsBorrowed = false
				books[bookKey] = book // Important: Update the book in the map
				resp := BorrowResponse{
					Result: "thank you for returning this book",
					Error:  "",
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(resp)
				return
			}
		}
	}

	// Book not found
	resp := ResponseRes{
		Result: "",
		Error:  "this book does not exist",
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(resp)
}
