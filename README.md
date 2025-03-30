## Overview

A secure file-sharing backend system built with Go, AWS S3, PostgreSQL, Redis, JWT authentication, and WebSocket notifications. This system provides a robust solution for file management with features like encryption, caching, and real-time updates.

## Project Structure

The project follows a modular package-based architecture:

```
.
├── pkg/
│   ├── api/                 # API handlers and routes
│   │   ├── file_handlers.go # File operation handlers
│   │   └── user_handlers.go # User authentication handlers
│   ├── auth/                # Authentication logic
│   │   ├── auth.go          # User authentication
│   │   └── middleware.go    # Authentication middleware
│   ├── cache/               # Caching implementation
│   │   └── cache.go         # Redis cache management
│   ├── db/                  # Database operations
│   │   └── db.go            # Database connection and queries
│   └── storage/             # File storage implementations
│       └── local_storage.go # Local file storage
├── main.go                  # Application entry point
└── README.md                # Project documentation
```

## Features

### Core Features

-  **User Authentication & Authorization**

   -  Secure registration and login
   -  JWT-based authentication
   -  Password hashing with bcrypt

-  **File Management**

   -  AES-256 encrypted file uploads to AWS S3
   -  File metadata caching with Redis
   -  File search by name, date, and type
   -  File sharing via pre-signed URLs
   -  Secure file deletion
   -  Periodic file expiration and cleanup

-  **Real-time Notifications**
   -  WebSocket-based notifications for file operations
   -  Real-time updates for file uploads and shares

### Technical Features

-  RESTful API endpoints
-  Middleware for authentication and rate limiting
-  Structured error handling
-  Comprehensive logging
-  Background worker for periodic tasks
-  Docker support for containerization

## Tech Stack

-  **Backend**: Go (1.18+)
-  **Database**: PostgreSQL (14+)
-  **Caching**: Redis (v6+)
-  **Storage**: AWS S3
-  **Authentication**: JWT
-  **Real-time**: WebSocket
-  **Security**: AES-256 encryption, bcrypt hashing

## Prerequisites

-  Go 1.18 or higher
-  PostgreSQL 14+
-  Redis v6+
-  AWS S3 Bucket
-  Docker (optional)

## Installation

### 1. Clone the Repository

```bash
git clone https://github.com/22BEY10041_Backend.git
cd 22BEY10041_Backend
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Set Up PostgreSQL

Create a new PostgreSQL database and user:

```sql
CREATE DATABASE filesharing;
CREATE USER filesharing_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE filesharing TO filesharing_user;
```

Create necessary tables:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE files (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    file_name TEXT NOT NULL,
    upload_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    size BIGINT NOT NULL,
    expiration_date TIMESTAMP,
    s3_key TEXT NOT NULL
);
```

### 4. Set Up Redis

Install and run Redis:

```bash
sudo apt-get install redis
redis-server
```

### 5. Set Up AWS S3

1. Create an S3 bucket on AWS
2. Configure AWS credentials:

   ```bash
   aws configure
   ```

### 6. Configure Environment Variables

Create a `.env` file:

```bash
# Server Configuration
PORT=8080

# Database
PG_HOST=localhost
PG_PORT=5432
PG_USER=filesharing_user
PG_PASSWORD=your_password
PG_DB=filesharing

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379

# AWS
AWS_REGION=your-region
AWS_S3_BUCKET=your-bucket-name

# Security
JWT_SECRET=your_jwt_secret_key
ENCRYPTION_KEY=your_encryption_key

# WebSocket
WS_PORT=8081
```

### 7. Run the Application

```bash
go run main.go
```

The server will start at `http://localhost:8080` and WebSocket server at `ws://localhost:8081`

## API Documentation

### Authentication

#### Register User

```http
POST /api/v1/auth/register
Content-Type: application/json

{
    "email": "user@example.com",
    "password": "password"
}
```

#### Login User

```http
POST /api/v1/auth/login
Content-Type: application/json

{
    "email": "user@example.com",
    "password": "password"
}
```

### File Operations

#### Upload File

```http
POST /api/v1/files/upload
Authorization: Bearer <JWT Token>
Content-Type: multipart/form-data

file: <your_file>
```

#### List Files

```http
GET /api/v1/files
Authorization: Bearer <JWT Token>
```

#### Search Files

```http
GET /api/v1/files/search?query=keyword&type=pdf
Authorization: Bearer <JWT Token>
```

#### Share File

```http
GET /api/v1/files/share/{file_id}
Authorization: Bearer <JWT Token>
```

#### Delete File

```http
DELETE /api/v1/files/{file_id}
Authorization: Bearer <JWT Token>
```

## Security

-  All API endpoints are protected with JWT authentication
-  Files are encrypted using AES-256 before being uploaded to S3
-  Passwords are hashed using bcrypt
-  Rate limiting is implemented to prevent abuse
-  Input validation is enforced for all requests