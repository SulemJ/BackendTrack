package main

import (
	// "fmt"
	"library_management/controllers"
	"library_management/models"
	"library_management/services"
)

func main() {
	var lib = services.NewLibrary()
	// 	test addbook
	lib.AddBook(models.Book{ID: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Status: "available"})
	lib.AddBook(models.Book{ID: 2, Title: "silent patient", Author: "Fitzgerald", Status: "available"})
	lib.AddMember(models.Member{ID: 1, Name: "Alice", BorrowedBooks: []models.Book{}})
	controllers.Start(lib)
}
