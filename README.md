# Boks-Boks-Boks API

**Backend REST API for the comprehensive storage box management application.**

The Boks-Boks-Boks API is a robust Go-based backend service that powers the storage management system. It provides secure authentication, data persistence, and comprehensive endpoints for managing storage boxes, items, and labels. Built with modern Go practices and designed for scalability and performance.

## What's Boks-Boks-Boks API?

The Boks-Boks-Boks API is a RESTful backend service that provides:

- **🔐 Secure Authentication** - JWT-based authentication with bcrypt password hashing
- **📦 Box Management** - Create, read, update, and delete storage boxes
- **📝 Item Management** - Manage items within boxes with quantities and relationships
- **🏷️ Label System** - Flexible labeling system with color-coding and descriptions
- **💾 Data Persistence** - PostgreSQL database with GORM ORM for reliable data storage
- **🚀 High Performance** - Built with Go and Gin framework for excellent performance
- **🐳 Container Ready** - Docker support for easy deployment and scaling
- **🔄 CORS Support** - Cross-origin resource sharing for frontend integration

Perfect as a backend service for inventory management systems, storage organization apps, or any application requiring structured item tracking.

## Requirements

- **Go**: Version 1.24.4 or higher
- **PostgreSQL**: Version 12 or higher
- **Docker**: Optional, for containerized deployment
- **Environment Variables**: See configuration section below

## Run Locally

### 1. Clone the repository

```bash
git clone https://github.com/boks-boks-boks/api.git
cd api
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Set up environment variables

Create a `.env` file in the root directory:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=boks_db
SECRET_TOKEN=your-secret-jwt-key
```

### 4. Set up PostgreSQL database

Make sure PostgreSQL is running and create the database:

```sql
CREATE DATABASE boks_db;
```

The application will automatically run the initialization SQL script on startup.

### 5. Run the development server

```bash
# Format code and run
./build.sh

# Or run directly
go run .
```

The API will be available at `http://localhost:8080`

### Production Build

To build for production:

```bash
go build -o main .
./main
```

### Using Docker

The project includes Docker support for easy deployment:

```bash
# Build and run with Docker Compose (includes PostgreSQL)
docker-compose up --build
```

The API will be available at `http://localhost:8080`.

## Deployment

### Docker Deployment

1. Build the Docker image:

```bash
docker build -t boks-api .
```

2. Run with environment variables:

```bash
docker run -p 8080:8080 \
  -e DB_HOST=your_db_host \
  -e DB_USER=your_db_user \
  -e DB_PASSWORD=your_db_password \
  -e DB_NAME=boks_db \
  -e SECRET_TOKEN=your-secret-key \
  boks-api
```

3. Or use Docker Compose for a complete setup:

```bash
docker-compose up -d
```

## Author

This project is entirely made by me (ASTOLFI Vincent). I suggest you to check on my github profile if you want to see the other project I've done for my studies or the ones I do in my free time.

## See also

Frontend of this project: [boks-boks-boks/boks-boks-boks](https://github.com/boks-boks-boks/boks-boks-boks)
