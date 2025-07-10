# Book Rental API (Go + Gin)

A simple RESTful API to manage books and users for a rental system.

## Getting Started

### Setup
```bash
git clone https://github.com/your-username/book-rental.git
cd book-rental
go mod tidy
go run main.go
```

### Running Tests
```bash
cd book-rental\tests\
go test
```

### API Endpoints
- `POST /books` – Add a new book
- `GET /books` – List all books
- `POST /users` – Create a new user
- `POST /rent` – Rent a book
- `POST /return` – Return a book

[POSTMAN API Documentation](https://documenter.getpostman.com/view/29514478/2sB34fkfnr)

### Folder Structure
- `models/` – Struct definitions
- `handlers/` – Route logic
- `tests/` – Unit tests







