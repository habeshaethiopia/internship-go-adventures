package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
	"strings"

	"bufio"
	"os"
	"regexp"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func clear() {
	print("\033[H\033[2J")

}
func helper(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z\s,&]+$`)
	return re.MatchString(s)
}
func Input() {
	numofbooks:=1
	clear()
	lib := &services.Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
	boldwelcome := color.New(color.FgCyan, color.Bold)
	input := color.New(color.FgCyan)
	color.New(color.FgGreen, color.Bold).Println("Welcome to the Library Management System!")
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Println("Choose an option:")
		color.Cyan("1. Add a new book")
		color.Cyan("2. Remove an existing book")
		color.Cyan("3. Borrow a book")
		color.Cyan("4. Return a book")
		color.Cyan("5. Add a new member")
		color.Cyan("6. List all available books")
		color.Cyan("7. List all borrowed books by a member")
		color.Cyan("8. List all the member")
		color.Cyan("9. Exit")

		fmt.Print("Enter the number of your choice: ")
		var in int
		_, err := fmt.Scanln(&in)
		if err != nil {
			clear()
			color.Red("Error in input")
			continue
		}

		switch in {
		case 1:
			clear()
			newBook := models.Book{}
			boldwelcome.Println("Add a new book")
			input.Print("Enter the book Title: ")
			// fmt.Scanln(&newBook.Title)
			newBook.Title, _ = reader.ReadString('\n')
			newBook.Title = strings.TrimSpace(newBook.Title)
			if newBook.Title == "" {
				color.Red("Error in input title must be a alphabet")
				continue
			}
			input.Print("Enter the book Author: ")
			// fmt.Scanln(&newBook.Author)
			text, _ := reader.ReadString('\n')
			if !helper(text) {
				color.Red("Error in input author must be a alphabet and separated by comma or &")
				continue

			}
			newBook.Author = text

			newBook.Status = "Available"
			newBook.ID = numofbooks
			numofbooks++

			lib.AddBook(newBook)

			// Call the function to add a new book
		case 2:
			clear()
			boldwelcome.Println("Remove the book")
			input.Print("Enter the book ID: ")
			var bookID int
			_, err = fmt.Scanln(&bookID)
			if err != nil {
				color.Red("Error in input id must be a number")
				continue
			}
			lib.RemoveBook(bookID)

			// Call the function to remove an existing book
		case 3:

			clear()
			boldwelcome.Println("Borrow a book")
			input.Print("Enter the book ID: ")
			var bookID int
			_, err = fmt.Scanln(&bookID)
			if err != nil {
				color.Red("Error in input id must be a number")
				continue

			}
			input.Print("Enter the member ID: ")
			var memberID int
			_, err = fmt.Scanln(&memberID)
			if err != nil {
				color.Red("Error in input id must be a number")
				continue
			}
			if lib.BorrowBook(bookID, memberID) != nil {
				color.Red("fail to borrow")
			}
			// Call the function to borrow a book
		case 4:
			clear()
			boldwelcome.Println("Return Borrowed Book")
			input.Print("Enter the book ID: ")
			var bookID int
			_, err = fmt.Scanln(&bookID)
			if err != nil {
				color.Red("Error in input id must be a number")
				continue

			}
			input.Print("Enter the member ID: ")
			var memberID int
			_, err = fmt.Scanln(&memberID)
			if err != nil {
				color.Red("Error in input id must be a number")
				continue
			}
			if lib.ReturnBook(bookID, memberID) != nil {
				color.Red("fail to return ")

			}
			// Call the function to return a book
		case 5:
			clear()
			boldwelcome.Println("Add a new member")
			newMember := models.Member{}
			input.Print("Enter the member Name: ")
			// fmt.Scanln(&newMember.Name)
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			if !helper(text) {
				color.Red("Error in input author must be a alphabet and separated by comma or &")
				continue
			}
			newMember.Name = text
			newMember.ID = len(lib.Members) + 1
			lib.AddMembers(newMember)

		case 6:
			clear()
			boldwelcome.Println("List of all available books")
			books := lib.ListAvailableBooks()
			if len(books) == 0 {
				color.Red("No books available")
				continue
			}
			tab := table.New("ID", "Title", "Author", "Status")
			for _, book := range books {
				tab.AddRow(book.ID, book.Title, book.Author, book.Status)
			}
			tab.Print()
			// Call the function to list all available books
		case 7:
			clear()
			boldwelcome.Print("Borrowed Books")
			input.Print("Enter the member ID: ")
			var memberID int
			_, err = fmt.Scanln(&memberID)
			if err != nil {
				color.Red("Error in input id must be a number")
				continue
			}
			books := lib.ListBorrowedBooks(memberID)
			if len(books) == 0 {

				color.Red("No books available")
				continue
			}
			tab := table.New("ID", "Title", "Author", "Status")
			for _, book := range lib.ListBorrowedBooks(memberID) {
				tab.AddRow(book.ID, book.Title, book.Author, book.Status)
			}
			tab.Print()

			// Call the function to list all borrowed books by a member
		case 8:
			tab := table.New("ID", "Name")
			for _, mem := range lib.Members {
				tab.AddRow(mem.ID, mem.Name)
			}
			tab.Print()
		case 9:
			clear()
			return
		default:
			color.Red("Invalid option. Please enter a number between 1 and 9.")
		}
	}

}
