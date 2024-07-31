package services

import (
	"errors"
	"fmt"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
	AddMembers(member models.Member)
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

// var mylib Library{
// 	Book:  make(map[int]models.Book),
// 	Member: make(map[int]models.Member),
// }


func (mylib *Library) AddBook(book models.Book) {
	book.Status = "Available"
	mylib.Books[book.ID] = book
	fmt.Printf("%s is added Successfully\n", book.Title)
}
func (mylib *Library) AddMembers(member models.Member) {
	mylib.Members[member.ID] = member
	fmt.Printf("%s is added Successfully\n", member.Name)

}
func (mylib *Library) RemoveBook(bookID int) {
	if value, ok := mylib.Books[bookID]; ok {
		delete(mylib.Books, bookID)
		fmt.Printf("%s is deleted Successfully", value.Title)

	} else {
		fmt.Println("The book in not found")

	}
}
func (mylib *Library) BorrowBook(bookID int, memberID int) error {
	err := errors.New("the book is not available")
	if value, ok := mylib.Books[bookID]; ok {
		if value.Status == "Available" {
			value.Status = "Borrowed"
			mylib.Books[bookID] = value
			member := mylib.Members[memberID]
			member.BorrowedBooks = append(member.BorrowedBooks, value)
			mylib.Members[memberID] = member
			err = nil
			fmt.Println("The book is borrowed successfully")
		}

	}
	return err
}
func (mylib *Library) ReturnBook(bookID int, memberID int) error {
	err := errors.New("the book is not borrowed")
	if value, ok := mylib.Books[bookID]; ok {
		if value.Status == "Borrowed" {
			member := mylib.Members[memberID]
			var borr []models.Book
			for _, val := range member.BorrowedBooks {
				if val != value {

					borr = append(borr, val)
				}
			}
			member.BorrowedBooks = borr
			value.Status = "Available"
			mylib.Books[bookID] = value
			err = nil
			fmt.Println("The book is returned successfully")

		}
	}
	return err
}
func (mylib *Library) ListAvailableBooks() []models.Book {
	var B []models.Book
	for _, val := range mylib.Books {
		if val.Status == "Available" {
			B = append(B, val)
		}
	}
	return B
}
func (mylib *Library) ListBorrowedBooks(memberID int) []models.Book {
	return mylib.Members[memberID].BorrowedBooks
}
