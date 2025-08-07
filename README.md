# 📚 Book Store API – Golang + AWS

A cloud-ready **Book Store API** built with Golang using the **MVC architecture**.  
This project is designed to showcase **clean backend design**, **containerization with Docker**, and **AWS deployment with CI/CD pipelines**.

---

## 🧠 Project Overview

This API allows clients to **create, read, update, and delete books**.  
It uses **integer IDs** for unique identification, follows **REST principles**, and is designed to be **cloud-native**.

---

## 🚀 Tech Stack

- **Language:** Golang
- **Framework:** Chi Router
- **Architecture:** MVC
- **Database:** (In-memory for now, AWS RDS planned)
- **Cloud Provider:** AWS (Elastic Beanstalk, RDS, S3)
- **Containerization:** Docker
- **CI/CD:** GitHub Actions → AWS
- **Version Control:** Git + GitHub

---

## 🛠 Features

- ✅ GET `/books` → Get all books
- ✅ GET `/book/{id}` → Get a single book by ID
- ✅ POST `/book` → Add a new book (ID auto-generated)
- ✅ PUT `/book/{id}` → Update a book’s title
- ✅ DELETE `/book/{id}` → Remove a book

---

## 🏗 Milestones

## ✅ Stage 1 – Local Development

## 🔜 Stage 2 – Database Integration

## 🔜 Stage 3 – Containerization & Deployment

## 🔜 Stage 4 – Cloud-Native Upgrade
---

## 📸 AWS Deployment

All AWS proof will be stored in the `/images` folder:

- Elastic Beanstalk running instance screenshot
- RDS instance configuration screenshot
- S3 bucket with uploaded files screenshot
- GitHub Actions pipeline run screenshot

---

## 📊 Current Status
- ✅ Basic CRUD API with in-memory storage
- ✅ Chi router with proper HTTP methods
- ✅ JSON request/response handling
- ✅ Error handling for invalid requests
- ⏳ Database integration (Phase 2)

## 📂 Project Structure
```
book-store/
├── controllers/
│   └── bookController.go    # CRUD operations
├── models/
│   └── book.go             # Book data structure
├── routes/
│   └── routes.go           # API route definitions
├── main.go                 # Application entry point
├── go.mod                  # Go module dependencies
├── .gitignore             # Git ignore rules
└── README.md              # Project documentation
```


---

## 🔄 Running Locally

```bash
# Clone repo
git clone https://github.com/NeilNeel/go-book-store.git
cd go-book-store

# Install dependencies
go mod tidy

# Run server
go run main.go
```

Server will start on `http://localhost:3000`

## 📋 API Examples

### Get all books
```bash
curl http://localhost:3000/books
```

### Get a specific book
```bash
curl http://localhost:3000/book/1
```

### Create a new book
```bash
curl -X POST http://localhost:3000/book \
  -H "Content-Type: application/json" \
  -d '{"title":"The Go Programming Language"}'
```

### Update a book
```bash
curl -X PUT http://localhost:3000/book/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Book Title"}'
```

### Delete a book
```bash
curl -X DELETE http://localhost:3000/book/1
```

## 📝 API Response Format

### Book Object
```json
{
  "id": 1,
  "title": "Book Title"
}
```

### Error Response
```
HTTP 400 Bad Request
Error while creating a new Book!
```