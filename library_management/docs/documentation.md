# Library Management System (GoLang)

## Task 3: Basic System

- Demonstrates structs, interfaces, methods, slices, and maps.
- Features: Add, Remove, Borrow, Return, List Books.

## Task 4: Concurrent Reservations

- Introduced Goroutines, Channels, and Mutexes.
- Supports concurrent reservation requests.
- Auto-cancels reservations not borrowed within 5 seconds.

### Concurrency Design:

1. Reservation requests call `ReserveBook()` in a goroutine.
2. Mutex ensures safe concurrent access to the shared `Books` map.
3. A background timer goroutine reverts reservations after 5 seconds.

### Example Scenario:

- Member 1 reserves Book 101.
- Member 2 tries to reserve the same book → receives an error.
- If Member 1 doesn’t borrow within 5 seconds, reservation is auto-cancelled.
