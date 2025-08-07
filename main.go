package main

import (
	"fmt"
	"log"
	"net/http"

	c "github.com/NeilNeel/go-bookstore/controllers"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	
	r.Get("/books",c.GetBooks)
	r.Get("/books/{id}",c.GetBook)
	r.Post("/books",c.CreateBook)
	r.Put("/books/{id}", c.UpdateBook)
	r.Delete("/books/{id}", c.DeleteBook)

	fmt.Println("Server is running on: http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000",r))
}