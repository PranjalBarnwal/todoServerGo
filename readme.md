# 📘 Go ToDo API Documentation

This RESTful API allows you to manage a simple to-do list with full CRUD operations.

---

## 🔗 Frontend Live URL 
- **Live**: [https://keploygotodo.netlify.app/](https://keploygotodo.netlify.app/)
- **GitHub (Frontend)**: [https://github.com/PranjalBarnwal/todoClientGo](https://github.com/PranjalBarnwal/todoClientGo)

---

## 🔗 Base URL

- **Local**: `http://localhost:8080`
- **Production (example)**: `https://todoservergo.onrender.com`

---

## 📄 Endpoints

### 📍 GET `/todos`
Retrieve all to-do items.

**Method:** `GET`  
**Response:**
```json
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
```json
{
  "title": "Read a book",
  "completed": false
}
```

**Response:**
```json
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
```json
{
  "title": "Read a novel",
  "completed": true
}
```

**Response:**
```json
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

## 🔧 Local Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/PranjalBarnwal/todoServerGo.git
   cd todoServerGo
   ```

2. **Set environment variables**

   Create a `.env` file (or export manually):
   ```env
   PORT=8080
   DATABASE_URL=your_neon_postgres_connection_string
   ```

3. **Install dependencies**
   ```bash
   go mod tidy
   ```

4. **Run the server**
   ```bash
   go run main.go
   ```

   Server will run on `http://localhost:8080`.

5. **Optional**: Use [Postman](https://www.postman.com/) or [curl](https://curl.se/) to test API routes.

---

## 👤 Author

Pranjal Barnwal  
[GitHub](https://github.com/pranjalbarnwal) | [LinkedIn](https://linkedin.com/in/pranjalbarnwal)