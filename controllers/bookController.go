package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	m "github.com/NeilNeel/go-bookstore/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

var books = []m.Book{
	{ID:"1", Title:"Book One"},
	{ID:"2", Title:"Book Two"},
}

func GetBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(books)

	if err != nil{
		http.Error(w,"Error while fetching the books.",http.StatusNotFound)
		return
	}
}

func GetBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	bookID := strings.TrimSpace(chi.URLParam(r, "id"))

	for _, book := range books{
		if strings.EqualFold(book.ID, bookID){
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var newBook m.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)

	for _, book := range books{
		if strings.EqualFold(book.Title, newBook.Title){
			http.Error(w, "Book already exists!", http.StatusBadRequest)
			return
		}
	}

	if err != nil{
		http.Error(w, "Error while creating a new Book!",http.StatusBadRequest)
		return
	}

	if newBook.Title == ""{
		http.Error(w,"Book title missing.",http.StatusBadRequest)
		return
	}

	newBook.ID = uuid.New().String()
	newBook.Title = strings.TrimSpace(newBook.Title)
	books = append(books, newBook)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	bookID := strings.TrimSpace(chi.URLParam(r,"id"))
	if bookID == ""{
		http.Error(w, "Book ID missing.", http.StatusBadRequest)
		return
	}

	var updatedBook m.Book
	for index, book := range books{
		if strings.EqualFold(book.ID,bookID){

			err := json.NewDecoder(r.Body).Decode(&updatedBook)
			if err != nil{
				http.Error(w, "Error while updating the book.", http.StatusBadRequest)
				return
			}

			books[index].Title = updatedBook.Title

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(books[index])
			return
		}
	}


}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application.json")
	bookID := strings.TrimSpace(chi.URLParam(r,"id"))

	if bookID == ""{
		http.Error(w, "Book ID missing.", http.StatusBadRequest)
		return 
	}

	for index, book := range books{
		if strings.EqualFold(book.ID,bookID){

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)

			books = append(books[:index], books[index+1:]...)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}