# Go Blog REST API

A backend REST API for a blog application built with Go, PostgreSQL, JWT authentication, Docker, and the Chi router.

**Repository:** https://github.com/kamrulcse404/go-blog-api  
**Author:** Md. Kamrul Hasan

---

## 🚀 Project Overview

This project is a RESTful Blog API that provides:

- User registration and login
- JWT-based authentication
- Password hashing with bcrypt
- Authenticated blog post creation, update, and deletion
- Public post listing and individual post retrieval
- Search, filtering, sorting, pagination
- PostgreSQL database integration
- Database migrations
- Dockerized PostgreSQL
- Request logging
- Panic recovery
- Ownership-based authorization for posts
- Basic request validation
- Database indexes for commonly used post queries

---

## 🛠️ Technologies Used

| Technology | Purpose |
|---|---|
| Go 1.25.0 | Backend programming language |
| Chi v5.3.2 | HTTP router |
| PostgreSQL 16 | Relational database |
| `database/sql` | Database access |
| lib/pq | PostgreSQL driver |
| JWT v5.3.1 | Authentication |
| bcrypt | Password hashing |
| godotenv | Loading environment variables |
| Docker / Docker Compose | PostgreSQL containerization |

---

## 📚 Topics Covered

This project covers several important backend development concepts in Go.

### Go Backend

- HTTP server with `net/http`
- REST API design
- Request handling
- JSON encoding and decoding
- Context handling
- HTTP status codes
- Error handling
- Middleware
- Structs and JSON tags
- Environment variables
- Package-based project structure

### REST API

- GET, POST, PUT, DELETE
- Route parameters
- Query parameters
- Request validation
- Pagination
- Searching
- Filtering
- Sorting
- Authentication and authorization
- Consistent HTTP error responses

### PostgreSQL

- Database and table design
- Primary keys
- Foreign keys
- Unique constraints
- Relationships between users and posts
- SQL queries
- JOIN
- INSERT / SELECT / UPDATE / DELETE
- `ILIKE` search
- `COUNT`
- Database indexes
- Parameterized queries
- PostgreSQL error handling

### Authentication & Security

- User registration
- Login
- JWT access tokens
- Bearer token authentication
- JWT expiration
- JWT algorithm validation
- Password hashing with bcrypt
- Password verification
- Protected routes
- User ownership authorization
- Password hash excluded from JSON responses

### Middleware

- Authentication middleware
- Request logger middleware
- Panic recovery middleware
- Context-based user ID

### Docker

- Docker Compose
- PostgreSQL container
- Environment-based database configuration
- Persistent PostgreSQL volume

---

## ✨ Features

### 1. User Registration

Users can create an account using:

- Name
- Email
- Password

Passwords are stored as bcrypt hashes instead of plain text.

**Endpoint**

```http
POST /users/register
```

Example request:

```json
{
  "name": "Kamrul Hasan",
  "email": "kamrul@example.com",
  "password": "password123"
}
```

---

### 2. User Login

Users can log in with their email and password.

```http
POST /users/login
```

Successful login returns:

- User information
- JWT access token

Example:

```json
{
  "user": {
    "id": 1,
    "name": "Kamrul Hasan",
    "email": "kamrul@example.com",
    "created_at": "2026-09-04T10:00:00Z",
    "updated_at": "2026-09-04T10:00:00Z"
  },
  "token": "JWT_TOKEN"
}
```

---

### 3. Get Current User

Returns the currently authenticated user.

```http
GET /users/me
```

Requires:

```http
Authorization: Bearer <token>
```

---

### 4. Create Post

Authenticated users can create blog posts.

```http
POST /posts
```

Example:

```json
{
  "title": "Learning Go",
  "content": "Go is a powerful language for backend development."
}
```

The authenticated user's ID is automatically assigned to the post.

---

### 5. Get All Posts

Public endpoint for retrieving posts.

```http
GET /posts
```

Supports:

- Pagination
- Search
- User filtering
- Sorting

Example:

```http
GET /posts?limit=10&offset=0
```

---

### 6. Search Posts

Searches both post title and content.

```http
GET /posts?search=golang
```

The implementation uses PostgreSQL `ILIKE`.

---

### 7. Filter Posts by User

Posts can be filtered by author/user ID.

```http
GET /posts?user_id=1
```

---

### 8. Sort Posts

Posts can be sorted by creation time.

Newest first:

```http
GET /posts?sort=newest
```

Oldest first:

```http
GET /posts?sort=oldest
```

---

### 9. Pagination

The API supports `limit` and `offset`.

```http
GET /posts?limit=10&offset=20
```

Maximum allowed limit:

```text
100
```

Example response structure:

```json
{
  "data": [],
  "pagination": {
    "limit": 10,
    "offset": 20,
    "total": 100
  }
}
```

---

### 10. Get Single Post

Public endpoint for retrieving a post by ID.

```http
GET /posts/{id}
```

Example:

```http
GET /posts/1
```

---

### 11. Update Post

Only the owner of a post can update it.

```http
PUT /posts/{id}
```

Requires authentication.

```http
Authorization: Bearer <token>
```

Example:

```json
{
  "title": "Updated Go Post",
  "content": "Updated blog post content."
}
```

---

### 12. Delete Post

Only the owner of a post can delete it.

```http
DELETE /posts/{id}
```

Requires authentication.

---

## 🔐 Authentication

Protected endpoints use JWT Bearer authentication.

Header:

```http
Authorization: Bearer <JWT_TOKEN>
```

The authentication middleware:

1. Reads the `Authorization` header
2. Validates the Bearer format
3. Validates the JWT
4. Checks the token expiration
5. Extracts the user ID
6. Stores the user ID in the request context
7. Passes the request to the next handler

JWT tokens expire after **24 hours**.

---

## 🗂️ API Endpoints

| Method | Endpoint | Auth | Description |
|---|---|---|---|
| POST | `/users/register` | No | Register a user |
| POST | `/users/login` | No | Login and receive JWT |
| GET | `/users/me` | Yes | Get current user |
| GET | `/posts` | No | Get posts |
| GET | `/posts/{id}` | No | Get single post |
| POST | `/posts` | Yes | Create post |
| PUT | `/posts/{id}` | Yes | Update own post |
| DELETE | `/posts/{id}` | Yes | Delete own post |

---

## 🔎 Query Parameters

### Pagination

```text
limit
offset
```

Example:

```http
GET /posts?limit=10&offset=0
```

### Search

```text
search
```

Example:

```http
GET /posts?search=go
```

### Filter

```text
user_id
```

Example:

```http
GET /posts?user_id=1
```

### Sorting

```text
sort=newest
sort=oldest
```

Example:

```http
GET /posts?sort=oldest
```

Parameters can also be combined:

```http
GET /posts?search=go&user_id=1&sort=newest&limit=10&offset=0
```

---

## 🧱 Database Structure

### Users

```text
users
├── id
├── name
├── email
├── password_hash
├── created_at
└── updated_at
```

### Posts

```text
posts
├── id
├── user_id
├── title
├── content
├── created_at
└── updated_at
```

Relationship:

```text
users 1 ──────────── N posts
```

A user can have multiple posts, while each post belongs to one user.

---

## ⚡ Database Indexes

The project includes indexes for post filtering and sorting.

### User + Created Time + ID

```sql
CREATE INDEX IF NOT EXISTS idx_posts_user_created_id
ON posts(user_id, created_at DESC, id DESC);
```

### Created Time + ID

```sql
CREATE INDEX IF NOT EXISTS idx_posts_created_id
ON posts(created_at DESC, id DESC);
```

These indexes support common post listing patterns involving user filtering and chronological sorting.

---

## 🗃️ Database Migrations

The project uses SQL migration files to create and modify the database schema.

Migration responsibilities include:

- Creating the `posts` table
- Creating the `users` table
- Adding `user_id` to posts
- Adding the foreign key relationship
- Creating database indexes

Example migration execution:

```bash
docker exec -i blog-postgres psql -U bloguser -d blogdb < migrations/001_create_posts.sql
```

Use the appropriate database username and database name from your environment configuration.

---

## 🐳 Docker Setup

PostgreSQL runs inside a Docker container using Docker Compose.

The PostgreSQL service uses:

```text
postgres:16
```

A named Docker volume is used so database data persists when the container is restarted.

Start PostgreSQL:

```bash
docker compose up -d
```

Check running containers:

```bash
docker ps
```

Stop the services:

```bash
docker compose down
```

---

## ⚙️ Environment Variables

Create a `.env` file in the project root.

Example:

```env
POSTGRES_USER=bloguser
POSTGRES_PASSWORD=your_password
POSTGRES_DB=blogdb

DB_HOST=localhost
DB_PORT=5432
DB_USER=bloguser
DB_PASSWORD=your_password
DB_NAME=blogdb
DB_SSLMODE=disable

JWT_SECRET=your_super_secret_jwt_key
```

> Do not commit the `.env` file or real secrets to the repository.

---

## 💻 Installation & Setup

### 1. Clone the repository

```bash
git clone https://github.com/kamrulcse404/go-blog-api.git
cd go-blog-api
```

### 2. Check Go version

The project uses:

```text
Go 1.25.0
```

Check your installed version:

```bash
go version
```

### 3. Configure environment variables

Create:

```text
.env
```

and add the required database and JWT configuration.

### 4. Start PostgreSQL

```bash
docker compose up -d
```

Verify:

```bash
docker ps
```

### 5. Run database migrations

Execute the migration SQL files in the correct order.

For example:

```bash
docker exec -i blog-postgres psql -U bloguser -d blogdb < migrations/001_create_posts.sql
```

Use your configured PostgreSQL username/database name if they differ.

### 6. Download Go dependencies

```bash
go mod download
```

### 7. Run the API

```bash
go run .
```

The server starts on:

```text
http://localhost:8080
```

---

## 🧪 Example API Flow

A typical usage flow is:

### Step 1 — Register

```http
POST /users/register
```

### Step 2 — Login

```http
POST /users/login
```

Copy the JWT token from the response.

### Step 3 — Create a Post

Use:

```http
POST /posts
```

with:

```http
Authorization: Bearer <JWT_TOKEN>
```

### Step 4 — Get Posts

```http
GET /posts
```

### Step 5 — Update Your Post

```http
PUT /posts/1
```

with the JWT token.

### Step 6 — Delete Your Post

```http
DELETE /posts/1
```

with the JWT token.

---

## 🛡️ Validation

The API validates user and post input.

### User validation

- Name is required
- Name must be at least 3 characters
- Name maximum length: 100 characters
- Email is required
- Email format is validated
- Email maximum length: 255 characters
- Password is required
- Password minimum length: 8 characters
- Password maximum length: 72 bytes

### Post validation

- Title is required
- Title minimum length: 3 characters
- Title maximum length: 255 characters
- Content is required
- Content minimum length: 10 characters

---

## 🧩 Middleware

The project includes three middleware components.

### Authentication Middleware

Protects private routes and validates JWT tokens.

### Logger Middleware

Logs:

- HTTP method
- Request path
- Status code
- Request duration

Example:

```text
POST /posts 201 12.4ms
```

### Recovery Middleware

Recovers from unexpected panics and returns:

```http
500 Internal Server Error
```

instead of crashing the server.

---

## 📁 Project Structure

```text
go-blog-api/
│
├── handlers/
│   ├── post_handler.go
│   ├── post_validation.go
│   ├── user_handler.go
│   └── user_validation.go
│
├── middleware/
│   ├── auth.go
│   ├── logger.go
│   └── recovery.go
│
├── models/
│   ├── post.go
│   └── user.go
│
├── security/
│   ├── jwt.go
│   └── password.go
│
├── storage/
│   ├── db.go
│   ├── post_storage.go
│   └── user_storage.go
│
├── migrations/
│   ├── 001_create_posts.sql
│   ├── 002_create_users.sql
│   ├── 003_add_user_id_to_posts.sql
│   └── 004_posts_indexes.sql
│
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── .env
└── README.md
```

---

## 🎯 Project Goals

The main goal of this project was to build a practical REST API in Go while learning and applying:

- Go backend development
- REST API architecture
- PostgreSQL
- SQL
- Database relationships
- Authentication
- JWT
- Password hashing
- Middleware
- Input validation
- Pagination
- Search
- Filtering
- Sorting
- Database indexing
- Docker
- Environment configuration
- Error handling

---

## 👨‍💻 Author

**Md. Kamrul Hasan**

GitHub repository:

https://github.com/kamrulcse404/go-blog-api
