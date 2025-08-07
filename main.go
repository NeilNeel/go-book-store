package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NeilNeel/go-bookstore/controllers"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	
	r.Get("/books",controllers.GetBooks)

	fmt.Println("Server is running on: http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000",r))
}