package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/NeilNeel/go-bookstore/models"
)

var books = []models.Book{
	{ID:1, Title:"Book One"},
	{ID:2, Title:"Book Two"},
}

func GetBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}