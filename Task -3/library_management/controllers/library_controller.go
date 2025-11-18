package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
)

func Start(lib *services.Library) {
	// accpting input from user
	for {

		fmt.Println("\n===== Library Management System =====")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Register member")
		fmt.Println("4. Borrow Book")
		fmt.Println("5. Return Book")
		fmt.Println("6. List Available Books")
		fmt.Println("7. List Borrowed Books by Member")
		fmt.Println("8. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var (
				id     int
				title  string
				author string
			)
			fmt.Println("Enter book ID, Title, Author:")
			fmt.Scanln(&id)
			fmt.Scanln(&title)
			fmt.Scanln(&author)
			lib.AddBook(models.Book{ID: id, Title: title, Author: author, Status: "available"})
			fmt.Println("Book added successfully!")
		case 2:

			// remove book from the id user passes
			var ridId int
			fmt.Println("Enter book ID to remove:")
			fmt.Scanln(&ridId)
			lib.RemoveBook(ridId)
			fmt.Println("Book removed successfully!")
		case 3:
			// register member from user input
			var (
				memId int
				name  string
			)
			fmt.Println("Enter your Id and name:")
			fmt.Scanln(&memId)
			fmt.Scanln(&name)
			lib.AddMember(models.Member{ID: memId, Name: name, BorrowedBooks: []models.Book{}})
			fmt.Println("member registered successfully!")
		case 4:
			// borrow book from the id user passes
			var (
				bookId   int
				memberId int
			)
			fmt.Println("Enter book ID and member Id to borrow:")
			fmt.Scanln(&bookId)
			fmt.Scanln(&memberId)
			chec := lib.BorrowBook(bookId, memberId)
			if chec != nil {
				fmt.Println(chec)
				break
			}
			fmt.Println("Book borrowed successfully!")
			// fmt.Println(lib.members)
		case 5:
			// return book from the id user passes
			var (
				retId    int
				memretId int
			)
			fmt.Println("Enter book ID and member Id to return:")
			fmt.Scanln(&retId)
			fmt.Scanln(&memretId)
			ret := lib.ReturnBook(retId, memretId)
			if ret != nil {
				fmt.Println(ret)
				break
			}
			fmt.Println("Book returned successfully!")
			// fmt.Println(lib.members)
		case 6:
			// test list available books
			fmt.Println(lib.ListAvailableBooks())
		case 7:
			// test list borrowed books
			fmt.Println(lib.ListBorrowedBooks(1))
		case 8:
			return
		}

	}

}
