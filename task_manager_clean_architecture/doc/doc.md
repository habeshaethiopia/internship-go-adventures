# API Documentation for Task Manager

## Base URL
```
localhost:8080
```

## Endpoints

[Here](https://documenter.getpostman.com/view/37364622/2sA3rzKt15) is the published postman doc

### Register User
**Endpoint:** `POST /register`

This endpoint allows you to register a new user.

**Request Body:**
- `email` (string): The email address of the user.
- `password` (string): The password for the user account.
- `role` (string): The role of the user.

**Response:**
The response will include a token that you can use to authenticate and retrieve information later on.

**Testing:**
A test has been included to confirm if a token is returned. Additionally, test scripts have been added to copy the token to the token collection variable, making it easy to reuse the token in other requests within the collection.

**Example Request:**
```json
{
    "email": "adnelkdcn@lkdc.dsks",
    "password": "laskdaskasldas",
    "role": "admin"
}
```

**Example Response:**
```json
{
    "token": "your_jwt_token_here"
}
```

**Example cURL:**
```sh
curl --location 'localhost:8080/register' \
--data-raw '{
    "email": "adnelkdcn@lkdc.dsks",
    "password": "laskdaskasldas",
    "role": "admin"
}'
```

### Login User
**Endpoint:** `POST /login`

This endpoint allows users to log in and obtain a JWT token for authentication.

**Request Body:**
- `email` (string): The email address of the user.
- `password` (string): The password for the user account.

**Response:**
Status: 200
Content-Type: application/json
The response will include a message, a JWT token, and user details including ID, email, password, and role.

**Example Request:**
```json
{
    "email": "adnelkdcn@lkdc.dsks",
    "password": "laskdaskasldas"
}
```

**Example Response:**
```json
{
    "message": "User logged in successfully",
    "token": "your_jwt_token_here",
    "user": {
        "id": "66b2b2449b6a2c3369ea60f3",
        "email": "adnelkdcn@lkdc.dsks",
        "password": "$2a$10$OWQ.MQ.w0x7GVeY0nRbJXO/BhVOvQwFIT.YQoOv99Byej6zFX0w8a",
        "role": "admin"
    }
}
```

**Example cURL:**
```sh
curl --location 'localhost:8080/login' \
--data-raw '{
    "email": "adnelkdcn@lkdc.dsks",
    "password": "laskdaskasldas"
}'
```

### Setting JWT Token in Postman
To set the JWT token obtained from the response to a variable in Postman, you can use the following script in the Tests tab of the request:
```javascript
var jsonData = pm.response.json();
pm.environment.set("jwtToken", jsonData.token);
```

### Get All Users
**Endpoint:** `GET /auth/users`

This request retrieves a list of users' information.

**Request Headers:**
- `Authorization: Bearer <token>`

**Response:**
Status: 200
Content-Type: application/json
The response body includes an array of user objects, each containing the user's ID, email, password, and role.

**Example Response:**
```json
[
    {
        "id": "user_id",
        "email": "user_email",
        "password": "user_password",
        "role": "user_role"
    }
]
```

### Get User by ID
**Endpoint:** `GET /auth/users/:id`

This endpoint retrieves the details of a specific user identified by their unique ID.

**Request Headers:**
- `Authorization: Bearer <token>`

**Response:**
Status: 200
Content-Type: application/json
The response contains the details of the user including their ID, email, password, and role.

**Example Response:**
```json
{
    "id": "user_id",
    "email": "user_email",
    "password": "user_password",
    "role": "user_role"
}
```

### Update User
**Endpoint:** `PUT /auth/user/:id`

This endpoint updates the details of a specific user.

**Request Headers:**
- `Authorization: Bearer <token>`

**Request Body:**
- `email` (string, required): The new email address for the user.

**Response:**
Status: 200 OK
Content-Type: application/json
The updated user details.

**Example Request:**
```json
{
    "email": "updatedemail@main.com"
}
```

**Example Response:**
```json
{
    "id": "user_id",
    "email": "updatedemail@main.com",
    "password": "user_password",
    "role": "user_role"
}
```

### Delete User
**Endpoint:** `DELETE /auth/user/:id`

This endpoint deletes a specific user.

**Request Headers:**
- `Authorization: Bearer <token>`

**Response:**
Status: 200
Content-Type: application/json

**Example Response:**
```json
{
    "message": "User deleted successfully"
}
```

### Create Task
**Endpoint:** `POST /api/tasks/`

This endpoint allows the client to add a new task.

**Request Headers:**
- `Authorization: Bearer <token>`

**Request Body:**
- `title` (string, required): The title of the task.
- `description` (string, required): The description of the task.
- `status` (string, required): The status of the task.

**Response:**
Status: 200
Content-Type: application/json
The created task details.

**Example Request:**
```json
{
    "title": "test task 1",
    "description": "This is a test",
    "status": "progressed"
}
```

**Example Response:**
```json
{
    "id": "task_id",
    "title": "test task 1",
    "description": "This is a test",
    "status": "progressed",
    "due_date": "due_date",
    "user_id": "user_id"
}
```

### Get Tasks
**Endpoint:** `GET /api/tasks/`

This endpoint retrieves a list of tasks from the server.

**Request Headers:**
- `Authorization: Bearer <token>`

**Response:**
Status: 200
Content-Type: application/json
A list of task objects.

**Example Response:**
```json
[
    {
        "id": "task_id",
        "title": "task_title",
        "description": "task_description",
        "status": "task_status",
        "due_date": "task_due_date",
        "user_id": "user_id"
    }
]
```

### Get Task by ID
**Endpoint:** `GET /api/tasks/:id`

This request retrieves the details of a specific task identified by its unique ID.

**Request Headers:**
- `Authorization: Bearer <token>`

**Response:**
Status: 200
Content-Type: application/json
The task details.

**Example Response:**
```json
{
    "id": "task_id",
    "title": "task_title",
    "description": "task_description",
    "status": "task_status",
    "due_date": "task_due_date",
    "user_id": "user_id"
}
```

### Update Task
**Endpoint:** `PUT /api/tasks/:id`

This endpoint updates a specific task identified by its ID.

**Request Headers:**
- `Authorization: Bearer <token>`

**Request Body:**
- `title` (string, optional): The updated title of the task.
- `description` (string, optional): The updated description of the task.

**Response:**
Status: 200
Content-Type: application/json
The updated task details.

**Example Request:**
```json
{
    "title": "test task 1 updated test",
    "description": "This is a test updated test 2"
}
```

**Example Response:**
```json
{
    "id": "task_id",
    "title": "test task 1 updated test",
    "description": "This is a test updated test 2",
    "status": "task_status",
    "due_date": "task_due_date",
    "user_id": "user_id"
}
```

### Delete Task
**Endpoint:** `DELETE /api/tasks/:id`

This endpoint deletes a specific task.

**Request Headers:**
- `Authorization: Bearer <token>`

**Response:**
Status: 200
Content-Type: application/json

**Example Response:**
```json
{
    "message": "Task deleted successfully"
}
```

---

Remember to replace placeholder values like `your_jwt_token_here`, `user_id`, `task_id`, `task_title`, etc., with actual values as needed.