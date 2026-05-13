# Golang API Backend

A production-style backend API built using Go with:

* REST APIs
* JWT Authentication
* Redis Session Management
* Kafka Event Streaming
* MySQL Database
* Docker Infrastructure
* Layered Architecture

---

# Tech Stack

| Technology | Purpose                   |
| ---------- | ------------------------- |
| Go         | Backend Language          |
| MySQL      | Database                  |
| Redis      | Session Storage           |
| Kafka      | Event Streaming           |
| Docker     | Infrastructure Management |
| JWT        | Authentication            |
| bcrypt     | Password Hashing          |
| kafka-go   | Kafka Client              |

---

# Features

* User Registration
* User Login
* JWT Authentication
* Redis Token Validation
* Logout API
* CRUD Operations
* Kafka Login Events
* Kafka Email Events
* Layered Architecture
* Centralized JSON Responses
* Logging System
* Input Validation

---

# Project Architecture

```text
Client
   ↓
Routes
   ↓
Handlers
   ↓
Services
   ↓
Repository
   ↓
MySQL
```

---

# Kafka Architecture

```text
Register User
   ↓
Publish EMAIL_EVENT
   ↓
Kafka Topic
   ↓
Email Consumer
```

```text
Login User
   ↓
Publish LOGIN_EVENT
   ↓
Kafka Topic
   ↓
Login Consumer
```

---

# Folder Structure

```text
golang-api/
│
├── config/
├── consumer/
│   ├── email/
│   └── login/
├── handlers/
├── logger/
├── middleware/
├── models/
├── repository/
├── routes/
├── services/
├── utils/
├── main.go
├── docker-compose.yml
└── README.md
```

---

# Layer Responsibilities

## Handlers

Responsible for:

* handling HTTP requests
* decoding JSON
* sending responses
* logging request errors

---

## Services

Responsible for:

* business logic
* validation
* Kafka event publishing
* Redis session management
* JWT handling

---

## Repository

Responsible for:

* database queries
* MySQL interaction

---

# Logging Levels

| Level | Purpose                |
| ----- | ---------------------- |
| INFO  | successful operations  |
| WARN  | validation/auth issues |
| ERROR | failures/exceptions    |
| DEBUG | internal debugging     |

---

# Docker Infrastructure

Services managed using Docker Compose:

* MySQL
* Redis
* Kafka
* Zookeeper

---

# Run Infrastructure

```bash
docker-compose up -d
```

---

# Verify Containers

```bash
docker ps
```

---

# Run Backend

```bash
go run main.go
```

---

# Run Kafka Consumers

## Email Consumer

```bash
go run consumer/email/main.go
```

## Login Consumer

```bash
go run consumer/login/main.go
```

---

# API Endpoints

## Authentication

| Method | Endpoint              |
| ------ | --------------------- |
| POST   | /api/v1/auth/register |
| POST   | /api/v1/auth/login    |
| POST   | /api/v1/auth/logout   |

---

## User APIs

| Method | Endpoint         |
| ------ | ---------------- |
| GET    | /api/v1/users    |
| POST   | /api/v1/users    |
| PUT    | /api/v1/users    |
| DELETE | /api/v1/users    |
| GET    | /api/v1/users/me |

---

# JWT Authentication

Protected routes use:

```text
Authorization: Bearer <token>
```

---

# Redis Usage

Redis stores:

* active JWT tokens
* session validation

If token exists in Redis:

```text
Valid Session
```

If token removed:

```text
Invalid Session
```

---

# Kafka Topics

| Topic      | Purpose      |
| ---------- | ------------ |
| user-login | login events |
| user-email | email events |

---

# Kafka Events

## Login Event

```json
{
  "user_id": 1,
  "email": "rahul@gmail.com",
  "event": "USER_LOGIN"
}
```

---

## Email Event

```json
{
  "email": "rahul@gmail.com",
  "name": "Rahul",
  "event": "SEND_WELCOME_EMAIL",
  "message": "Welcome to Golang API"
}
```

---

# Database Schema

```sql
CREATE TABLE users (

    id INT AUTO_INCREMENT PRIMARY KEY,

    name VARCHAR(100) NOT NULL,

    email VARCHAR(100) NOT NULL UNIQUE,

    password VARCHAR(255) NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    ON UPDATE CURRENT_TIMESTAMP
);
```

---

# Security Features

* Password hashing using bcrypt
* JWT authentication
* Redis-based token validation
* Protected middleware routes
* Input validation
* Unique email validation

---

# Important Backend Concepts Learned

* REST APIs
* Layered Architecture
* JWT Authentication
* Middleware
* Redis Sessions
* Kafka Producers
* Kafka Consumers
* Event-Driven Architecture
* Docker Containers
* Logging System
* Input Validation
* Repository Pattern
* Service Layer

---

# Future Improvements

* Real Email Integration
* Refresh Tokens
* Rate Limiting
* Swagger Documentation
* Unit Testing
* Role-Based Authentication
* Dockerized Go App
* Kubernetes Deployment
* CI/CD Pipeline
* Kafka Retry Mechanism

---

# Learning Outcome

This project demonstrates:

* scalable backend architecture
* asynchronous event processing
* production-style service layering
* containerized infrastructure
* secure authentication system
* industry backend practices
