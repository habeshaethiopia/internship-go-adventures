# Library Management System

This project is a Library Management System implemented in Go. It provides a simple console-based interface for managing books and members in a library.

## Project Structure

The project is organized into several packages:

- `main.go`: This is the entry point of the application.
- `models/`: This package contains the data structures used in the application, including `Book` and `Member`.
- `controllers/`: This package contains the `library_controller.go` file, which handles user input and controls the flow of the application.
- `services/`: This package contains the `library_service.go` file, which provides the core functionality of the library management system, such as adding and removing books and members.

## Features

The Library Management System provides the following options:

1. Add a new book
2. Remove an existing book
3. Borrow a book
4. Return a book
5. Add a new member
6. List all available books
7. List all borrowed books by a member
8. List all the members
9. Exit

## Running the Project

To run the project, navigate to the root directory of the project and use the `go run` command:

```sh
go run main.go
```

## Testing

The project includes unit tests. To run the tests, use the `go test` command:

```sh
go test ./...
```

## Documentation

For more detailed information about the project, see the `documentation.md` file in the `docs/` directory.

Please note that this is a basic overview of the project. For more detailed information, please refer to the individual files and their comments.