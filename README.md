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

- Go 1.21 or higher
- PostgreSQL
- Docker (optional)

## Getting Started

1. Clone the repository:
```bash
git clone <repository-url>
cd go-cms
```

2. Copy the example environment file:
```bash
cp .env.example .env
```

3. Update the environment variables in `.env` file with your configuration.

4. Install dependencies:
```bash
go mod download
```

5. Run the application:
```bash
go run cmd/api/main.go
```

Or using Docker:
```bash
docker-compose up --build
```

## API Documentation

The API documentation is available at `/swagger/index.html` when running the application.

## Database Schema

The application uses PostgreSQL with the following main tables:
- Users
- Posts
- Categories
- Tags
- Media
- Roles
- Permissions

## Authentication

The API uses JWT (JSON Web Tokens) for authentication. Include the token in the Authorization header:
```
Authorization: Bearer <your-token>
```

## API Endpoints

### Authentication
- POST /api/v1/auth/register
- POST /api/v1/auth/login
- POST /api/v1/auth/refresh

### Users
- GET /api/v1/users
- GET /api/v1/users/:id
- PUT /api/v1/users/:id
- DELETE /api/v1/users/:id

### Posts
- GET /api/v1/posts
- POST /api/v1/posts
- GET /api/v1/posts/:id
- PUT /api/v1/posts/:id
- DELETE /api/v1/posts/:id

### Categories
- GET /api/v1/categories
- POST /api/v1/categories
- GET /api/v1/categories/:id
- PUT /api/v1/categories/:id
- DELETE /api/v1/categories/:id

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

This project is licensed under the MIT License. 