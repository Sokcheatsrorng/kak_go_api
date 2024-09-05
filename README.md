## Note

### Install go.mod file nd go.sum 

```bash 

   go mod init

   go mod tidy 

```
### Local Port 8080 
----------------------------------------

### Run local api with cmd/author_book_api/main.go

```bash

    go run main.go
 
```
### Model Author and Book 

### Endpoints Accessment with each models 

# Author
```
GET   : http://localhost/authors
GET   : http://localhost/authors/{id}
POST  : http://localhost/authors
PUT   : http://localhost/authors/{id}
DELETE: http://localhost/authors/{id}

```
# Book
```
GET   : http://localhost/books
GET   : http://localhost/books/{id}
POST  : http://localhost/books
PUT   : http://localhost/books/{id}
DELETE: http://localhost/books/{id}

```
### Using Reception with 
- gorm library  <a>https://gorm.io/index.html</a>
- gin package installation <a>https://github.com/gin-gonic/gin.git</a>
- gorm with auto migration <a>https://gorm.io/docs/migration.html</a>

### Thanks 


