

# API Documentation

This document describes the API endpoints, request/response formats for our Task Management Service.

## Base URL

All URLs referenced in the documentation have the following base:

`http://localhost:8080/api/v1`

The APIs are defined below.

## Endpoints

### GET /tasks

Fetches all tasks.

**Response**

A JSON array of tasks.

```json
[
    {
        "ID": "1",
        "Title": "Task 1",
        "Description": "First task",
        "DueDate": "2022-03-01T14:30:00Z",
        "Status": "Pending"
    },
    // ...
]
```

### POST /tasks

Creates a new task.

**Request**

A JSON object representing the task.

```json
{
    "Title": "Task 1",
    "Description": "First task",
    "DueDate": "2022-03-01T14:30:00Z",
    "Status": "Pending"
}
```

**Response**

A JSON object representing the created task.

```json
{
    "ID": "1",
    "Title": "Task 1",
    "Description": "First task",
    "DueDate": "2022-03-01T14:30:00Z",
    "Status": "Pending"
}
```

### GET /tasks/{id}

Fetches a task by ID.

**Response**

A JSON object representing the task.

```json
{
    "ID": "1",
    "Title": "Task 1",
    "Description": "First task",
    "DueDate": "2022-03-01T14:30:00Z",
    "Status": "Pending"
}
```

### PUT /tasks/{id}

Updates a task.

**Request**

A JSON object representing the updated task.

```json
{
    "Title": "Updated Task 1",
    "Description": "Updated first task",
    "DueDate": "2022-03-01T14:30:00Z",
    "Status": "In Progress"
}
```

**Response**

A JSON object representing the updated task.

```json
{
    "ID": "1",
    "Title": "Updated Task 1",
    "Description": "Updated first task",
    "DueDate": "2022-03-01T14:30:00Z",
    "Status": "In Progress"
}
```

### DELETE /tasks/{id}

Deletes a task.

**Response**

A JSON object representing the deleted task.

```json
{
    "ID": "1",
    "Title": "Task 1",
    "Description": "First task",
    "DueDate": "2022-03-01T14:30:00Z",
    "Status": "Pending"
}
```

## Postman Documentation

You can also view the Postman API documentation [here](https://documenter.getpostman.com/view/37364622/2sA3kdBdij).

## Data Storage

This API uses MongoDB as its data storage solution. When a new task is created via the API, it is inserted into a MongoDB collection. The task data is stored as a document in the collection, with each field in the task represented as a field in the document.

Here's an example of how a task is stored in MongoDB:

```json
{
    "id": "1",
    "Title": "Task 1",
    "Description": "First task",
    "DueDate": "2022-03-01T14:30:00Z",
    "Status": "Pending"
}
```

When tasks are fetched via the API, the API queries the MongoDB collection and returns the matching documents as tasks.
