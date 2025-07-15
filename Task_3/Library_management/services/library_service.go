package services

import (
	"Library_management/models"
	"errors"
)

type Library struct{
	Books map[int] models.Book
	Members map[int]*models.Member
	

}

type LibraryManager interface{

	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book


}
func (L *Library) AddBook(book models.Book){
	L.Books[book.ID] = book
}

func (L *Library) RemoveBook(bookID int) error {
    _, exists := L.Books[bookID]
    if !exists {
        return errors.New("book not found")
    }

    delete(L.Books, bookID)
    return nil
}

func(L *Library) BorrowBook(bookID int, memberID int) error{
	book,exists := L.Books[bookID]
	if !exists{
		return errors.New("book not found")
	}
	member,found := L.Members[memberID]
	if !found{
		return errors.New("user not found")
	} 
	if book.Status == "Borrowed" {
    return errors.New("book already borrowed")
} 
	
	book.Status = "Borrowed"
	L.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks,book)

	return nil	

}

func(L *Library) ReturnBook(bookID int, memberID int) error{
	book,exists := L.Books[bookID]
	if !exists{
		return errors.New("book not found")
	}
	member,found := L.Members[memberID]
	if !found{
		return errors.New("user not found")
	} 
	length := len(member.BorrowedBooks)
	for i := 0; i < length; i++ {
		if member.BorrowedBooks[i].ID == bookID{
			book.Status = "Available"
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			L.Books[bookID] = book
			
			return nil
		}

	}
	return errors.New("book was not borrowed by this member")

}

func (L *Library) ListAvailableBooks() []models.Book {
    listOfBooks := []models.Book{}
    for _, book := range L.Books {
        if book.Status == "Available" {
            listOfBooks = append(listOfBooks, book)
        }
    }
    return listOfBooks
}
func (L *Library) ListBorrowedBooks(memberID int) []models.Book {
    member, found := L.Members[memberID]
    if !found {
        return []models.Book{}   
    }
    return member.BorrowedBooks
}