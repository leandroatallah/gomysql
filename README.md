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

## Getting Started

1. Make sure you have Go and MySQL installed
2. Clone the repository
3. Set up your MySQL database and update the connection string in `pkg/config/database.go`
4. Run the application:

```bash
go run cmd/main/main.go
```
