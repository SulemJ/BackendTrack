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
	books   map[int]models.Book
	members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	// n = len(books)
	l.books[book.ID] = book
}
func (l *Library) RemoveBook(bookID int) {
	// n = len(books)
	delete(l.books, bookID)
}

func (l *Library) AddMember(member models.Member) {
	// n = len(books)
	l.members[member.ID] = member
}
func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, ok := l.books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	member, there := l.members[memberID]
	if !there {
		return errors.New("member not found")
	}
	if book.Status == "borrowed" {
		return errors.New("book already borrowed")
	}
	book.Status = "borrowed"
	l.books[bookID] = book
	member.BorrowedBooks = append(l.members[memberID].BorrowedBooks, book)
	l.members[memberID] = member
	return nil
}

func removeBook(members map[int]models.Member, bookID int, memberID int) {
	member := members[memberID]
	for i, books := range member.BorrowedBooks {
		if books.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}

	}
	members[memberID] = member
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, ok := l.books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	_, there := l.members[memberID]
	if !there {
		return errors.New("member not found")
	}
	if book.Status == "available" {
		return errors.New("book already returned")
	}
	book.Status = "available"
	l.books[bookID] = book
	removeBook(l.members, bookID, memberID)
	// l.members[memberID] = member
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	avails := []models.Book{}
	for _, book := range l.books {
		if book.Status == "available" {
			avails = append(avails, book)
		}

	}
	return avails
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	borrow := []models.Book{}
	for _, book := range l.books {
		if book.Status == "borrowed" {
			borrow = append(borrow, book)
		}
	}
	return borrow
}
