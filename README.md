# 📚 Author Book API

## 🚀 Project Setup

### 🛠 Install Dependencies

To set up the project and ensure all dependencies are installed, run the following commands:

```bash
go mod init
go mod tidy
```
🔧 Running the Application Locally
By default, the application runs on port 8080. To start the local API server, execute the following command from the cmd/author_book_api directory:
```bash
go run main.go
```
## 🌐 API Endpoints

All endpoints are accessible via `http://localhost:8080`.

### 👨‍💼 Author Endpoints
- **GET** `/authors` - Retrieve a list of all authors
- **GET** `/authors/{id}` - Retrieve a specific author by ID
- **POST** `/authors` - Create a new author
- **PUT** `/authors/{id}` - Update an existing author by ID
- **DELETE** `/authors/{id}` - Delete an author by ID

### 📖 Book Endpoints
- **GET** `/books` - Retrieve a list of all books
- **GET** `/books/{id}` - Retrieve a specific book by ID
- **POST** `/books` - Create a new book
- **PUT** `/books/{id}` - Update an existing book by ID
- **DELETE** `/books/{id}` - Delete a book by ID

## 🛠 Tools & Libraries Used
- 🔗 **GORM**: Object Relational Mapping (ORM) library for Go. For more information, visit the [GORM documentation](https://gorm.io/index.html).
- 🔥 **Gin**: HTTP web framework written in Go. Installation instructions can be found in the [Gin repository](https://github.com/gin-gonic/gin.git).
- ⚙️ **Auto Migration**: GORM's automatic schema migration tool. Refer to the [Auto Migration documentation](https://gorm.io/docs/migration.html) for further details.

