package concurrency

import (
	"errors"
	"fmt"
	"library_management/models"
	"sync"
	"time"
)

type Reservation struct {
	BookID   int
	MemberID int
}

type ConcurrentLibrary struct {
	Books   map[int]models.Book
	Members map[int]*models.Member
	mu      sync.Mutex
}

func NewConcurrentLibrary(books map[int]models.Book, members map[int]*models.Member) *ConcurrentLibrary {
	return &ConcurrentLibrary{
		Books:   books,
		Members: members,
	}
}

// ReserveBook reserves a book concurrently
func (lib *ConcurrentLibrary) ReserveBook(bookID int, memberID int) error {
	lib.mu.Lock()
	defer lib.mu.Unlock()

	book, exists := lib.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status != "Available" {
		return errors.New("book already borrowed or reserved")
	}

	book.Status = "Reserved"
	lib.Books[bookID] = book
	fmt.Printf("\n[INFO] Book '%s' reserved by Member %d\n", book.Title, memberID)

	// Auto-cancel after 5 seconds if not borrowed
	go func(id int) {
		time.Sleep(5 * time.Second)
		lib.mu.Lock()
		defer lib.mu.Unlock()
		if lib.Books[id].Status == "Reserved" {
			book := lib.Books[id]
			book.Status = "Available"
			lib.Books[id] = book
			fmt.Printf("\n[INFO] Reservation for '%s' expired and is now available again.\n", book.Title)
		}
	}(bookID)

	return nil
}
