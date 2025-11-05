package controllers

import (
	"fmt"
	"library_management/concurrency"
	"library_management/models"
	"library_management/services"
)

func RunLibraryConsole() {
	library := services.NewLibrary()

	// Preload members
	library.Members[1] = &models.Member{ID: 1, Name: "Suleman"}
	library.Members[2] = &models.Member{ID: 2, Name: "Hana"}

	// Concurrency handler
	concurrentLib := concurrency.NewConcurrentLibrary(library.Books, library.Members)

	for {
		fmt.Println("\n===== Library Management System =====")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("7. Reserve Book (Concurrent)")
		fmt.Println("8. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id int
			var title, author string
			fmt.Print("Enter Book ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Title: ")
			fmt.Scan(&title)
			fmt.Print("Enter Author: ")
			fmt.Scan(&author)
			library.AddBook(models.Book{ID: id, Title: title, Author: author})
			fmt.Println(" Book added successfully!")

		case 2:
			var id int
			fmt.Print("Enter Book ID to remove: ")
			fmt.Scan(&id)
			library.RemoveBook(id)
			fmt.Println(" Book removed successfully!")

		case 3:
			var bookID, memberID int
			fmt.Print("Enter Book ID: ")
			fmt.Scan(&bookID)
			fmt.Print("Enter Member ID: ")
			fmt.Scan(&memberID)
			if err := library.BorrowBook(bookID, memberID); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(" Book borrowed successfully!")
			}

		case 4:
			var bookID, memberID int
			fmt.Print("Enter Book ID: ")
			fmt.Scan(&bookID)
			fmt.Print("Enter Member ID: ")
			fmt.Scan(&memberID)
			if err := library.ReturnBook(bookID, memberID); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(" Book returned successfully!")
			}

		case 5:
			fmt.Println("\n Available Books:")
			for _, b := range library.ListAvailableBooks() {
				fmt.Printf("ID: %d | %s by %s | Status: %s\n", b.ID, b.Title, b.Author, b.Status)
			}

		case 6:
			var memberID int
			fmt.Print("Enter Member ID: ")
			fmt.Scan(&memberID)
			fmt.Println("\n Borrowed Books:")
			for _, b := range library.ListBorrowedBooks(memberID) {
				fmt.Printf("ID: %d | %s by %s\n", b.ID, b.Title, b.Author)
			}

		case 7:
			var bookID, memberID int
			fmt.Print("Enter Book ID: ")
			fmt.Scan(&bookID)
			fmt.Print("Enter Member ID: ")
			fmt.Scan(&memberID)
			err := concurrentLib.ReserveBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(" Reservation initiated! (Auto-cancels after 5 seconds if not borrowed.)")
			}

		case 8:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice!")
		}
	}
}
