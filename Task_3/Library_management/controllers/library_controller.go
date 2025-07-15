package controllers

import (
	"Library_management/models"
	"Library_management/services"
	"fmt"
)

func HandleAddBook(lib services.LibraryManager) {
    var id int
    var title, author string

    fmt.Print("Enter Book ID: ")
    fmt.Scanln(&id)
    fmt.Print("Enter Book Title: ")
    fmt.Scanln(&title)
    fmt.Print("Enter Author: ")
    fmt.Scanln(&author)


    book := models.Book{
        ID:     id,
        Title:  title,
        Author: author,
        Status: "Available",
    }

    lib.AddBook(book)
    fmt.Println(" Book added successfully:", title)
}

func HandleBorrowBook (lib services.LibraryManager){
	var bookId,memberId int
	fmt.Print("Enter the book ID you want to borrow: ")
	fmt.Scanln(&bookId)
	fmt.Print("Enter the the memberID ")
	fmt.Scanln(&memberId)


	err := lib.BorrowBook(bookId,memberId)
	if err != nil {
    fmt.Println("Could not borrow book:", err.Error())
    return
}
	fmt.Println("Successfully borrowed")
}
func HandleRemoveBook (lib services.LibraryManager){
	var bookID int
	fmt.Print("Enter the book ID you want to remove: ")
	fmt.Scanln(&bookID)
	err := lib.RemoveBook(bookID)
	if err != nil {
		fmt.Println("Could not remove book:", err.Error())
		return
	}
	fmt.Println("Successfully removed")

}

func HandleReturnBook(lib services.LibraryManager){
	var bookId,memberId int
	fmt.Print("Enter the book ID you want to Return: ")
	fmt.Scanln(&bookId)
	fmt.Print("Enter the the memberID ")
	fmt.Scanln(&memberId)

	err := lib.ReturnBook(bookId,memberId)
	if err != nil {
    fmt.Println("Could not Retur book:", err.Error())
    return
}
	fmt.Println("Successfully Returned the book")
}

func HandleAvailableBooks(lib services.LibraryManager) {
    books := lib.ListAvailableBooks()

    if len(books) == 0 {
        fmt.Println("📭 No books are currently available.")
        return
    }

    fmt.Println("Available Books:")
    for _, book := range books {
        fmt.Printf("• [%d] %s by %s (%d)\n", book.ID, book.Title, book.Author)
    }
}
func HandleBorrowedBooks(lib services.LibraryManager) {
    var memberID int
    fmt.Print("Enter the Member ID: ")
    fmt.Scanln(&memberID)

    books := lib.ListBorrowedBooks(memberID)

    if len(books) == 0 {
        fmt.Println("No borrowed books found for this member.")
        return
    }

    fmt.Printf("📖 Borrowed Books for Member ID %d:\n", memberID)
    for _, book := range books {
        fmt.Printf("• [%d] %s by %s (%d)\n", book.ID, book.Title, book.Author)
    }
}
