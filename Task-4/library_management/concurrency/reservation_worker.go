package concurrency

import (
	"fmt"
	"library_management/services"
	"time"
)

// func (l *Library) ReserveBook(bookID, memberID int) error{
// 	var ReservationQueue chan ReservationRequests

// }

func StartReservationWorker(lib *services.Library) {
	go func() {
		for req := range lib.ReservationQueue {
			lib.Mu.Lock()
			book, ok := lib.Books[req.BookID]
			if !ok {
				lib.Mu.Unlock()
				req.Response <- fmt.Errorf("book not found")
				continue
			}

			if book.Status != "available" {
				lib.Mu.Unlock()
				req.Response <- fmt.Errorf("book already reserved or borrowed")
				continue
			}
			// fmt.Println("Received reservation request:", req.BookID, req.MemberID)
			book.Status = "reserved"
			lib.Books[req.BookID] = book
			lib.Mu.Unlock()
			req.Response <- nil
			go func(bookID int) {
				time.Sleep(7 * time.Second)

				lib.Mu.Lock()
				defer lib.Mu.Unlock()

				currentBook := lib.Books[bookID]
				if currentBook.Status == "reserved" {
					currentBook.Status = "available"
					lib.Books[bookID] = currentBook
					fmt.Println("Reservation expired for book:", bookID)
				}
			}(req.BookID)
		}
	}()
}
