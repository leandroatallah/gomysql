# GoMySQL

A simple REST API built with Go and MySQL.

## Features

- User CRUD operations
- MySQL database integration
- RESTful endpoints

## API Endpoints

- `GET /users` - Get all users
- `POST /users` - Create a new user
- `GET /users/{id}` - Get user by ID
- `DELETE /users/{id}` - Delete user
- `PUT /users/{id}` - Update user

## Getting Started

1. Make sure you have Go and MySQL installed
2. Clone the repository
3. Set up your MySQL database and update the connection string in `pkg/config/database.go`
4. Run the application:

```bash
go run cmd/main/main.go
```

## References

- HTTP Server [https://gowebexamples.com/http-server/](https://gowebexamples.com/http-server/)
- MySQL Database [https://gowebexamples.com/mysql-database/](https://gowebexamples.com/mysql-database/)
- Password Hashing [https://gowebexamples.com/password-hashing/](https://gowebexamples.com/password-hashing/)
- Go Wiki: Rate Limiting [https://go.dev/wiki/RateLimiting](https://go.dev/wiki/RateLimitingo)
- gorilla/mux [https://github.com/gorilla/mux](https://github.com/gorilla/mux)
- golang-migrate [https://github.com/golang-migrate/migrate/tree/master/database/mysql](https://github.com/golang-migrate/migrate/tree/master/database/mysql)
