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

## ✅ 🧪 Tests & Code Coverage

This project includes:

- ✅ **Unit Tests** for handler logic
- ✅ **Integration Tests** for DB interactions
- ✅ **API Tests** for end-to-end endpoint coverage
- ✅ **Code coverage** tracking using `go test -cover`

### 🔁 Run All Tests with Coverage

```bash
go test "-coverpkg=go-todo-app/handlers,go-todo-app/db,go-todo-app/models" ./tests/... -coverprofile=coverage.out
```

> ✅ **Test coverage achieved**: ~75%+

![image](https://github.com/user-attachments/assets/31f0e6f5-d914-4e92-a098-3bdd6f73672c)

## 🚦 Keploy API Test Report

Keploy is integrated into the CI/CD pipeline to automatically test API endpoints using real user traffic and record/replay capabilities.

- **Keploy runs on every push and pull request via GitHub Actions.**
- **API test cases are auto-generated and replayed for regression detection.**
- **Test results are available in the Keploy Cloud dashboard.**

### 📊 View Keploy Test Results
![report testcases keploy](https://github.com/user-attachments/assets/7b571830-c2a3-43a5-a3fd-9e650c5cd662)

---
## CI/CD Configuration

GitHub Actions to run the Keploy test suite.
[View CI Workflow](.github/workflows/main.yml)
---

## 👤 Author

Pranjal Barnwal  
[GitHub](https://github.com/pranjalbarnwal) | [LinkedIn](https://linkedin.com/in/pranjalbarnwal)
