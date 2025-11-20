package main

import (
	"fmt"
	"library_management/concurrency"
	"library_management/models"
	"library_management/services"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var lib = services.NewLibrary()
	// 	test addbook
	lib.AddBook(models.Book{ID: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Status: "available"})
	lib.AddBook(models.Book{ID: 2, Title: "silent patient", Author: "Fitzgerald", Status: "available"})
	lib.AddMember(models.Member{ID: 1, Name: "Alice", BorrowedBooks: []models.Book{}})
	lib.AddMember(models.Member{ID: 2, Name: "Bob", BorrowedBooks: []models.Book{}})
	lib.AddMember(models.Member{ID: 3, Name: "jack", BorrowedBooks: []models.Book{}})
	// controllers.Start(lib)
	concurrency.StartReservationWorker(lib)
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(MemberID int) {
			defer wg.Done()
			err := lib.ReserveBook(1, MemberID)
			if err != nil {
				fmt.Println("Member", MemberID, "failed to reserve:", err)
			} else {
				fmt.Println("Member", MemberID, "successfully reserved the book")
			}
		}(i)
	}
	wg.Wait()
}
