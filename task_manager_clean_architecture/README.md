# User Controller API

======================

This is a RESTful API built with Go and Gin framework to manage users. It provides endpoints for creating, reading, updating and deleting users.

## Features

---

- User registration and login functionality
- JWT token generation for authentication
- Role-based access control (admin and user roles)
- Endpoints for creating, reading, updating and deleting users

## Endpoints

---

### User Endpoints

- **POST /register**: Create a new user
- **GET /api/api/users**: Get all users (admin only)
- **GET /api/users/:id**: Get a user by ID (admin or self)
- **PUT /api/users/:id**: Update a user (admin or self)
- **DELETE /api/users/:id**: Delete a user (admin only)
- **POST /login**: Login a user and generate JWT token

### Task Endpoints

- **POST /api/task**: Create a new task
- **GET /api/tasks**: Get all tasks
- **GET /api/tasks/:id**: Get a task by ID
- **PUT /api/tasks/:id**: Update a task by ID
- **DELETE /api/tasks/:id**: Delete a task by ID

## Requirements

---

- Go 1.17 or later
- Gin framework
- MongoDB database
- JWT library for token generation

Here’s how you can represent your folder structure with relative links in your README file using Markdown:

## Folder Structure

```plaintext
task-manager/
├── [Delivery/](./Delivery/)
│   ├── [main.go](./Delivery/main.go)
│   ├── [controllers/](./Delivery/controllers/)
│   │   └── [controller.go](./Delivery/controllers/controller.go)
│   └── [routers/](./Delivery/routers/)
│       └── [router.go](./Delivery/routers/router.go)
├── [Domain/](./Domain/)
│   └── [domain.go](./Domain/domain.go)
├── [Infrastructure/](./Infrastructure/)
│   ├── [auth_middleWare.go](./Infrastructure/auth_middleWare.go)
│   ├── [jwt_service.go](./Infrastructure/jwt_service.go)
│   └── [password_service.go](./Infrastructure/password_service.go)
├── [Repositories/](./Repositories/)
│   ├── [task_repository.go](./Repositories/task_repository.go)
│   └── [user_repository.go](./Repositories/user_repository.go)
└── [Usecases/](./Usecases/)
    ├── [task_usecases.go](./Usecases/task_usecases.go)
    └── [user_usecases.go](./Usecases/user_usecases.go)
```

## Folder Descriptions

- **[Delivery/](./Delivery/):** Contains files related to the delivery layer, handling incoming requests and responses.

  - **[main.go](./Delivery/main.go):** Sets up the HTTP server, initializes dependencies, and defines the routing configuration.
  - **[controllers/](./Delivery/controllers/):**
    - **[controller.go](./Delivery/controllers/controller.go):** Handles incoming HTTP requests and invokes the appropriate use case methods.
  - **[routers/](./Delivery/routers/):**
    - **[router.go](./Delivery/routers/router.go):** Sets up the routes and initializes the Gin router.

- **[Domain/](./Domain/):** Defines the core business entities and logic.

  - **[domain.go](./Domain/domain.go):** Contains the core business entities such as `Task` and `User` structs.

- **[Infrastructure/](./Infrastructure/):** Implements external dependencies and services.

  - **[auth_middleWare.go](./Infrastructure/auth_middleWare.go):** Middleware to handle authentication and authorization using JWT tokens.
  - **[jwt_service.go](./Infrastructure/jwt_service.go):** Functions to generate and validate JWT tokens.
  - **[password_service.go](./Infrastructure/password_service.go):** Functions for hashing and comparing passwords to ensure secure storage of user credentials.

- **[Repositories/](./Repositories/):** Abstracts the data access logic.

  - **[task_repository.go](./Repositories/task_repository.go):** Interface and implementation for task data access operations.
  - **[user_repository.go](./Repositories/user_repository.go):** Interface and implementation for user data access operations.

- **[Usecases/](./Usecases/):** Contains the application-specific business rules.
  - **[task_usecases.go](./Usecases/task_usecases.go):** Implements the use cases related to tasks, such as creating, updating, retrieving, and deleting tasks.
  - **[user_usecases.go](./Usecases/user_usecases.go):** Implements the use cases related to users, such as registering, and logging in.

## Setup

---

1.  Install Go and Gin framework
2.  Set up a MongoDB database
3.  Clone this repository and run `go build` to build the API
4.  Run the API with `go run main.go`

## API Documentation

---

API documentation is available at [postman documentation](https://documenter.getpostman.com/view/37364622/2sA3rzKt15) or [doc/doc.md](doc/doc.md)

## Contributing

---

Contributions are welcome! Please submit a pull request with your changes and a brief description of what you've added or fixed.

## License

---

This project is licensed under the MIT License. See LICENSE file for details.
