package services

import (
	"errors"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]*models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]*models.Member),
	}
}

func (lib *Library) AddBook(book models.Book) {
	book.Status = "Available"
	lib.Books[book.ID] = book
}

func (lib *Library) RemoveBook(bookID int) {
	delete(lib.Books, bookID)
}

func (lib *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := lib.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status != "Available" {
		return errors.New("book is not available")
	}

	member, exists := lib.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	lib.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	return nil
}

func (lib *Library) ReturnBook(bookID int, memberID int) error {
	member, exists := lib.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			book := lib.Books[bookID]
			book.Status = "Available"
			lib.Books[bookID] = book
			return nil
		}
	}
	return errors.New("book not borrowed by this member")
}

func (lib *Library) ListAvailableBooks() []models.Book {
	available := []models.Book{}
	for _, book := range lib.Books {
		if book.Status == "Available" {
			available = append(available, book)
		}
	}
	return available
}

func (lib *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := lib.Members[memberID]
	if !exists {
		return []models.Book{}
	}
	return member.BorrowedBooks
}
