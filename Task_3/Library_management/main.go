package main

import (
	"Library_management/controllers"
	"Library_management/models"
	"Library_management/services"
	"fmt"
)

func main() {
   lib := &services.Library{
    Books:   make(map[int]models.Book),
    Members: make(map[int]*models.Member),
}

   
    lib.AddBook(models.Book{ID: 1, Title: "Go in Action", Author: "William Kennedy",Status: "Available"})
    lib.Members[101] = &models.Member{ID: 101, Name: "Zufan"}
    for {
        fmt.Println("\nLibrary Management System")
        fmt.Println("1. Add Book")
        fmt.Println("2. Borrow Book")
        fmt.Println("3. Return Book")
        fmt.Println("4. Remove Book")
        fmt.Println("5. List Available Books")
        fmt.Println("6. List Borrowed Books by Member")
        fmt.Println("7. Exit")

        var choice int
        fmt.Print("Choose an option: ")
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            controllers.HandleAddBook(lib)
        case 2:
            controllers.HandleBorrowBook(lib)
        case 3:
            controllers.HandleReturnBook(lib)
        case 4:
            controllers.HandleRemoveBook(lib)
        case 5:
            controllers.HandleAvailableBooks(lib)
        case 6:
            controllers.HandleBorrowedBooks(lib)
        case 7:
            fmt.Println("Goodbye!")
            return
        default:
            fmt.Println("Invalid option. Try again.")
        }
    }
}