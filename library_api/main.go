package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Server struct {
	port string
}

type Book struct {
	Author     string `json:"author"`
	Title      string `json:"title"`
	isBorrowed bool
}

func (book Book) GetId() string {
	return strings.ToLower(book.Author + " " + book.Title)
}

var books map[string]*Book

func NewServer(port string) *Server {
	books = make(map[string]*Book)
	return &Server{
		port: port,
	}
}

func (s *Server) Start() {
	http.HandleFunc("/book", booksHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%s", s.port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{

		}
	case "POST":
		{
			var newBook Book
			err := json.NewDecoder(r.Body).Decode(&newBook)
			if err != nil {
				http.Error(w, "Failed to parse the request body", http.StatusBadRequest)
				return
			} else {
				err = addNewBook(&newBook)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}
		}
	case "PUT":
		{
			var newBook Book
			err := json.NewDecoder(r.Body).Decode(&newBook)
			if err != nil {
				http.Error(w, "Failed to parse the request body", http.StatusBadRequest)
				return
			} else {
				err = addNewBook(&newBook)
				if err != nil {
					http.Error(w, "Failed to parse the request body", http.StatusBadRequest)
					return
				}
			}
		}
	case "DELETE":
		{

		}
	default:
		{
			http.Error(w, "Only GET,POST,PUT,DELETE requests are allowed", http.StatusMethodNotAllowed)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(""))
}

func addNewBook(book *Book) error {
	if book.Title == "" || book.Author == "" {
		return errors.New("title or author cannot be empty")
	}
	book.isBorrowed = false
	if _, ok := books[book.GetId()]; ok {
		return errors.New("this book is already in the library")
	} else {
		books[book.GetId()] = book
	}
	return nil
}

func updateBook(book *Book) error {
	if book.Title == "" || book.Author == "" {
		return errors.New("title or author cannot be empty")
	}

	if _, ok := books[book.GetId()]; ok {
		return errors.New("this book is already in the library")
	} else {
		books[book.GetId()] = book
	}
	return nil
}

func main() {
	server := NewServer("4001")
	go server.Start()
	for {

	}
}
