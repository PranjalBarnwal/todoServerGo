# 📘 Go ToDo API Documentation

This RESTful API allows you to manage a simple to-do list with full CRUD operations.

---

## 🔗 Frontend Live URL - 'https://keploygotodo.netlify.app/'

---

## 🔗 Base URL

- Local: `http://localhost:8080`
- Production (example): `https://todoservergo.onrender.com`

---

## 📄 Endpoints

### 📍 GET `/todos`
Retrieve all to-do items.

**Method:** `GET`  
**Response:**

```
[
  {
    "id": 1,
    "title": "Buy groceries",
    "completed": false
  },
  {
    "id": 2,
    "title": "Walk the dog",
    "completed": true
  }
]
```

---

### 🆕 POST `/todos`
Create a new to-do.

**Method:** `POST`  
**Request Body:**

```
{
  "title": "Read a book",
  "completed": false
}
```

**Response:**

```
{
  "id": 3,
  "title": "Read a book",
  "completed": false
}
```

---

### ✏️ PUT `/todos/{id}`
Update a to-do (e.g., change title or toggle completion).

**Method:** `PUT`  
**Request Body:**

```
{
  "title": "Read a novel",
  "completed": true
}
```

**Response:**

```
{
  "id": 3,
  "title": "Read a novel",
  "completed": true
}
```

---

### ❌ DELETE `/todos/{id}`
Delete a to-do.

**Method:** `DELETE`  
**Response:** HTTP 200 OK on success

---

## 🛠️ Tech Stack

- **Backend**: Go (Golang)
- **Database**: PostgreSQL (hosted on [Neon.tech](https://neon.tech))
- **Router**: Gorilla Mux
- **CORS**: RS CORS middleware

---

## 👤 Author

Pranjal Barnwal  
[GitHub](https://github.com/pranjalbarnwal) | [LinkedIn](https://linkedin.com/in/pranjalbarnwal)