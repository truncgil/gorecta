# GoRecta CMS

A modern and robust Content Management System backend built with Go, featuring a RESTful API architecture and comprehensive documentation.

## Features

### Core Features
- RESTful API architecture
- JWT-based authentication and authorization
- Role-based access control (RBAC)
- Swagger API documentation
- Docker containerization
- PostgreSQL database with GORM ORM
- Modular and clean code structure

### Content Management
- Blog post management with categories and tags
- Media file handling and storage
- User management system
- Rich text content support
- SEO-friendly slugs
- Featured images for posts

### Security Features
- Secure password hashing with bcrypt
- JWT token-based authentication
- Role-based middleware
- CORS configuration
- Environment-based configuration

### API Features
- Pagination and filtering
- Search functionality
- Sorting and ordering
- Error handling
- Response caching
- Rate limiting

## Tech Stack

### Backend
- Go 1.21+
- Gin Web Framework
- GORM ORM
- PostgreSQL 15
- JWT-Go
- Swagger/OpenAPI

### Development Tools
- Docker & Docker Compose
- Make (optional)
- Git
- Swagger UI

### Testing
- Go testing package
- Integration tests
- API tests
- Mock database support

## Project Structure

```
.
├── cmd/                    # Application entry points
│   └── api/
│       └── main.go        # Main application file
├── internal/              # Private application code
│   ├── api/              # API layer
│   │   ├── handlers/     # Request handlers
│   │   ├── middleware/   # Custom middleware
│   │   └── routes/       # Route definitions
│   ├── config/           # Configuration
│   ├── models/           # Database models
│   ├── repository/       # Data access layer
│   └── service/          # Business logic
├── pkg/                  # Public libraries
│   ├── auth/            # Authentication
│   ├── database/        # Database utilities
│   └── utils/           # Common utilities
├── docs/                # Documentation
├── scripts/             # Build and deployment scripts
├── .env.example         # Environment template
├── Dockerfile          # Docker build file
├── docker-compose.yml  # Docker compose config
├── go.mod             # Go modules file
└── README.md          # Project documentation
```

## Prerequisites

- Docker and Docker Compose
- Go 1.21 or higher (for local development)
- Git
- PostgreSQL (for local development)

## Quick Start

1. Clone the repository:
```bash
git clone https://github.com/truncgil/gorecta.git
cd gorecta
```

2. Copy and configure environment variables:
```bash
cp .env.example .env
# Edit .env file with your configurations
```

3. Start with Docker:
```bash
docker-compose up --build
```

The application will be available at:
- API: http://localhost:8080
- Swagger Documentation: http://localhost:8080/docs/index.html

## Environment Variables

Key environment variables that need to be configured:

```env
# Server
SERVER_PORT=8080
SERVER_HOST=0.0.0.0
GIN_MODE=debug|release

# Database
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=cms_db
DB_SSL_MODE=disable

# JWT
JWT_SECRET=your_secret_key
JWT_EXPIRATION=24h

# CORS
ALLOWED_ORIGINS=*
ALLOWED_METHODS=GET,POST,PUT,DELETE,OPTIONS
ALLOWED_HEADERS=Authorization,Content-Type
```

## API Documentation

### Authentication Endpoints
- POST /api/v1/auth/register - Register new user
- POST /api/v1/auth/login - User login

### User Management
- GET /api/v1/users - List users (Admin)
- GET /api/v1/users/:id - User details
- PUT /api/v1/users/:id - Update user
- DELETE /api/v1/users/:id - Delete user (Admin)

### Content Management
- GET /api/v1/posts - List blog posts
- POST /api/v1/posts - Create new post (Admin/Editor)
- GET /api/v1/posts/:id - Post details
- PUT /api/v1/posts/:id - Update post (Admin/Editor)
- DELETE /api/v1/posts/:id - Delete post (Admin)

### Category Management
- GET /api/v1/categories - List categories
- POST /api/v1/categories - Create category (Admin)
- GET /api/v1/categories/:id - Category details
- PUT /api/v1/categories/:id - Update category (Admin)
- DELETE /api/v1/categories/:id - Delete category (Admin)

### Tag Management
- GET /api/v1/tags - List tags
- POST /api/v1/tags - Create tag (Admin)
- GET /api/v1/tags/:id - Tag details
- PUT /api/v1/tags/:id - Update tag (Admin)
- DELETE /api/v1/tags/:id - Delete tag (Admin)

## Authentication

The API uses JWT (JSON Web Tokens) for authentication. Include the token in the Authorization header:
```
Authorization: Bearer <your_token>
```

## Role-Based Access Control

The system supports the following roles:
- Admin: Full access to all endpoints
- Editor: Can manage content but not users
- User: Can view content and manage their own profile

## Development

### Local Development Setup

1. Install dependencies:
```bash
go mod download
```

2. Run locally:
```bash
go run cmd/api/main.go
```

### Testing

Run tests:
```bash
go test ./...
```

### Building

Build the binary:
```bash
go build -o main cmd/api/main.go
```

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

### Coding Standards

- Follow Go best practices and idioms
- Use meaningful variable and function names
- Write tests for new features
- Update documentation as needed
- Follow semantic versioning

## Deployment

### Docker Deployment
```bash
docker-compose up --build
```

### Manual Deployment
1. Build the binary
2. Set up environment variables
3. Configure database
4. Run the application

## Monitoring and Maintenance

- Health check endpoint: GET /health
- Logging configuration in .env
- Database backup scripts in /scripts
- Monitoring endpoints for metrics

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Support

For support, please open an issue in the GitHub repository or contact the maintainers. 