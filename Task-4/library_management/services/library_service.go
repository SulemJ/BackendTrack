package services

import (
	"errors"
	// "library_management/concurrency"
	"library_management/models"
	"sync"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
	ReserveBook(bookID int, memberID int) error
}
type Library struct {
	Books            map[int]models.Book
	Members          map[int]models.Member
	Mu               sync.Mutex
	ReservationQueue chan ReservationRequests
}
type ReservationRequests struct {
	BookID   int
	MemberID int
	Response chan error
}

func NewLibrary() *Library {
	return &Library{
		Books:            make(map[int]models.Book),
		Members:          make(map[int]models.Member),
		ReservationQueue: make(chan ReservationRequests, 10),
	}
}
func (l *Library) ReserveBook(bookID, memberID int) error {
	response := make(chan error)
	req := ReservationRequests{
		BookID:   bookID,
		MemberID: memberID,
		Response: response,
	}
	l.ReservationQueue <- req
	err := <-response
	return err
}

// func (l *Library) ReserveBook(bookID int, memberID int) error{

// 	if
// }

func (l *Library) AddBook(book models.Book) {
	// n = len(Books)
	l.Books[book.ID] = book
}
func (l *Library) RemoveBook(bookID int) {
	// n = len(Books)
	l.Mu.Lock()
	defer l.Mu.Unlock()
	delete(l.Books, bookID)
}

func (l *Library) AddMember(member models.Member) {
	// n = len(Books)
	l.Mu.Lock()
	defer l.Mu.Unlock()
	l.Members[member.ID] = member
}
func (l *Library) BorrowBook(bookID int, memberID int) error {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	book, ok := l.Books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	member, there := l.Members[memberID]
	if !there {
		return errors.New("member not found")
	}
	if book.Status == "borrowed" {
		return errors.New("book already borrowed")
	}
	book.Status = "borrowed"
	l.Books[bookID] = book
	member.BorrowedBooks = append(l.Members[memberID].BorrowedBooks, book)
	l.Members[memberID] = member
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
	l.Mu.Lock()
	defer l.Mu.Unlock()
	book, ok := l.Books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	_, there := l.Members[memberID]
	if !there {
		return errors.New("member not found")
	}
	if book.Status == "available" {
		return errors.New("book already returned")
	}
	book.Status = "available"
	l.Books[bookID] = book
	removeBook(l.Members, bookID, memberID)
	// l.members[memberID] = member
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	avails := []models.Book{}
	for _, book := range l.Books {
		if book.Status == "available" {
			avails = append(avails, book)
		}

	}
	return avails
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	borrow := []models.Book{}
	for _, book := range l.Books {
		if book.Status == "borrowed" {
			borrow = append(borrow, book)
		}
	}
	return borrow
}
