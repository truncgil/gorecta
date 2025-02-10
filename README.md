# GoRecta CMS

A modern and robust Content Management System backend built with Go.

## Features

- RESTful API endpoints

- JWT Authentication
- Role-based access control (RBAC)
- Content management (Create, Read, Update, Delete operations)
- Media file handling
- User management
- Category and tag management
- Search functionality
- API documentation with Swagger

## Tech Stack

- Go 1.21+
- PostgreSQL
- GORM (ORM)
- Gin (Web Framework)
- JWT-Go
- Swagger
- Docker

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   ├── middleware/
│   │   └── routes/
│   ├── config/
│   ├── models/
│   ├── repository/
│   └── service/
├── pkg/
│   ├── auth/
│   ├── database/
│   └── utils/
├── docs/
├── .env.example
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```

## Prerequisites

- Docker and Docker Compose

## Quick Start

1. Clone the repository:
```bash
git clone <repository-url>
cd gorecta
```

2. Copy the example environment file:
```bash
cp .env.example .env
```

3. Edit the .env file and configure the necessary settings.

4. Start the application with Docker:
```bash
docker-compose up --build
```

This command will:
- Start PostgreSQL database
- Build and run the Go application
- Generate Swagger documentation
- Install all dependencies

## API Documentation

The API documentation is available at:
```
http://localhost:8080/docs/index.html
```

## API Endpoints

### Authentication
- POST /api/v1/auth/register - Register new user
- POST /api/v1/auth/login - User login

### Users
- GET /api/v1/users - List users (Admin)
- GET /api/v1/users/:id - User details
- PUT /api/v1/users/:id - Update user
- DELETE /api/v1/users/:id - Delete user (Admin)

### Posts
- GET /api/v1/posts - List blog posts
- POST /api/v1/posts - Create new post (Admin/Editor)
- GET /api/v1/posts/:id - Post details
- PUT /api/v1/posts/:id - Update post (Admin/Editor)
- DELETE /api/v1/posts/:id - Delete post (Admin)

### Categories
- GET /api/v1/categories - List categories
- POST /api/v1/categories - Create new category (Admin)
- GET /api/v1/categories/:id - Category details
- PUT /api/v1/categories/:id - Update category (Admin)
- DELETE /api/v1/categories/:id - Delete category (Admin)

## Authentication

The API uses JWT (JSON Web Tokens). Include the token in the Authorization header:
```
Authorization: Bearer <your-token>
```

## Development

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License. 