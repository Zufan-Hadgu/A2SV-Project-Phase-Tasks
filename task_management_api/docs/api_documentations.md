
---

## Endpoints Overview

| Method | Endpoint       | Description                |
|--------|----------------|----------------------------|
| GET    | `/tasks`       | Get all tasks              |
| GET    | `/tasks/:id`   | Get task by ID             |
| POST   | `/tasks`       | Add a new task             |
| PUT    | `/tasks/:id`   | Update a task by ID        |
| DELETE | `/tasks/:id`   | Delete a task by ID        |

---

## GET `/tasks`

### Description
Retrieve a list of all tasks.

### Request

- **Method:** GET
- **URL:** `/tasks`
- **Headers:**
  - `Content-Type: application/json`

### Request Body
❌ No request body required.

### Example Response `200 OK`

```json
[
  {
    "id": "1",
    "title": "First_tast",
    "description": "Study Go",
    "due_date": "2025-08-01T10:00:00Z",
    "status": "Pending"
  }
]

# 📋 Task Management API Documentation

This API allows you to manage tasks using in-memory storage. Below are the available endpoints, request formats, and expected responses.

---

## ✅ GET `/tasks/:id`

Fetch a single task by ID.

### URL Parameters

| Name | Type   | Required | Description     |
|------|--------|----------|-----------------|
| id   | string | ✅        | Task ID to find |

### Example Request

```bash
GET /tasks/1
```

### Success Response — 200 OK

```json
{
  "id": "1",
  "title": "First_tast",
  "description": "Study Go",
  "due_date": "2025-08-01T10:00:00Z",
  "status": "Pending"
}
```

### Error Response — 400 Bad Request

```json
{
  "error": "task not found"
}
```

---

## POST `/tasks`

Add a new task.

### 📝 Request Format

**Method**: POST  

### 🔸 Request Body

```json
{
  "title": "New Task",
  "description": "Learn Go routines",
  "due_date": "2025-08-10T14:00:00Z",
  "status": "Pending"
}
```

### Success Response — 201 Created

```json
{
  "message": "Task added successfully"
}
```

### Error Response — 400 Bad Request

```json
{
  "error": "EOR"
}
```

---

## PUT `/tasks/:id`

Update an existing task.

### 🔸 Request Body

```json
{
  "title": "Updated Task Title",
  "description": "Updated description",
  "status": "Done"
}
```

### Success Response — 200 OK

```json
{
  "Message": "Task updated"
}
```

### Error Response — 400 Bad Request

```json
{
  "error": "task not found"
}
```

---

## DELETE `/tasks/:id`

Delete a task by its ID.

### Success Response — 200 OK

```json
{
  "Message": "Task Deleted"
}
```

### Error Response — 400 Bad Request

```json
{
  "error": "task not found"
}
```

---