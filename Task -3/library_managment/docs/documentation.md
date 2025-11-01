# This is a simple library management system written in go, to demonstrate use of go struct, struct methods, interfaces.

## Folder structure

    library_management/
        ├── main.go
        ├── controllers/
        │   └── library_controller.go
        ├── models/
        │   └── book.go
        │   └── member.go
        ├── services/
        │   └── library_service.go
        ├── docs/
        │   └── documentation.md
        └── go.mod

- main.go: Entry point of the application.
- controllers/library_controller.go: Handles console input and invokes the appropriate service methods.
- models/book.go: Defines the Book struct.
- models/member.go: Defines the Member struct.
- services/library_service.go: Contains business logic and data manipulation functions.
- docs/documentation.md: Contains system documentation and other related information.
- go.mod: Defines the module and its dependencies.

## How to run

In your terminal - go run main.go
If everything’s in place, you’ll see:

===== Library Management System =====

1. Add Book
2. Remove Book
3. Borrow Book
4. Return Book
5. List Available Books
6. List Borrowed Books by Member
7. Exit
   Enter your choice:
