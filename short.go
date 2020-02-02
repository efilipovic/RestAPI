package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type book struct {
	ID                    string `json:"ID"`
	Title       		  string `json:"Title"`
	ISBN      			  string `json:"ISBN"`	
	YearOfPublishing      string `json:"YearOfPublishing"`	
}

type allBooks []book

var books allBooks

func initBooks() {
   books = allBooks{
	{
		ID:          "1",
		Title:       "Harry Potter and The Philosopher's Stone",
		ISBN:        "9788700631625",
		YearOfPublishing: "1997.",
	},
	{
		ID:          "2",
		Title:       "Harry Potter and The Chamber of Secrets",
		ISBN:        "9788700631625",
		YearOfPublishing: "1998.",
	},
	{
		ID:          "3",
		Title:       "Harry Potter and The Prisoner of Azkaban",
		ISBN:        "9788700631625",
		YearOfPublishing: "1999.",
	},
	{
		ID:          "4",
		Title:       "Harry Potter and The Goblet of Fire",
		ISBN:        "9788700631625",
		YearOfPublishing: "2000.",
	},
	{
		ID:          "5",
		Title:       "Harry Potter and The Order of the Phoenix",
		ISBN:        "9788700631625",
		YearOfPublishing: "2003.",
	},
	{
		ID:          "6",
		Title:       "Harry Potter and The Half-Blood Prince",
		ISBN:         "9788700631625",
		YearOfPublishing: "2005.",
	},
	{
		ID:          "7",
		Title:       "Harry Potter and The Deathly Hallows",
		ISBN: 		  "9788700631625",
		YearOfPublishing: "2007.",
	},
}
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var newBook book
	reqBody,_ := ioutil.ReadAll(r.Body)	
	
	json.Unmarshal(reqBody, &newBook)
	books = append(books, newBook)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}


func updateBook(w http.ResponseWriter, r *http.Request) {
	bookID := mux.Vars(r)["id"]
	var updatedBook book

	reqBody, _:= ioutil.ReadAll(r.Body)
	
	json.Unmarshal(reqBody, &updatedBook)

	for i, singleBook := range books {
		if singleBook.ID == bookID {
			singleBook.Title = updatedBook.Title
			singleBook.ISBN = updatedBook.ISBN
			singleBook.YearOfPublishing = updatedBook.YearOfPublishing
			books = append(books[:i], singleBook)
			json.NewEncoder(w).Encode(singleBook)			
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	bookID := mux.Vars(r)["id"]

	for i, singleBook := range books {
		if singleBook.ID == bookID {
			books = append(books[:i], books[i+1:]...)			
		}
	}
}

func main() {
	initBooks()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/book", createBook).Methods("POST")
	router.HandleFunc("/books", getAllBooks).Methods("GET")	
	router.HandleFunc("/books/{id}", updateBook).Methods("PATCH")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}