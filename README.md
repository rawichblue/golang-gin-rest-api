# API Golang Project

## Overview

 This project is a Go-based API development utilizing the Gin framework, Bun ORM, and JWT for secure authentication. It also integrates PostgreSQL for database management and Cobra for command-line functionality. The project demonstrates CRUD operations and other core features for managing user data efficiently.

## Features

- Basic CRUD operations for users
- JWT Authentication Login
- Password hashing and verification
- Docker and Docker Compose setup
- PostgreSQL database integration
- Pagination and search 

## Project Structure

- `cmd/`: Contains the main application files and routes.
  - `http.go`: Defines HTTP server setup.
  - `root.go`: Handles root routes.
  - `product.go`: Manages product-related routes.
  - `api.go`: Contains API routes and handlers.

- `models/`: Contains Go models for the application.
  - `employee.go`: Defines the Employee model.
  - `customer.go`: Defines the Customer model.

- `services/`: Contains service logic and business rules.
  - `auth.svc.go`: Handles authentication-related services.
  - `employee.svc.go`: Manages employee-related services.

- `controllers/`: Contains controller logic for handling requests.
  - `auth.ctl.go`: Manages authentication-related requests.
  - `employee.ctl.go`: Handles employee-related requests.

- `migrations/`: Contains database migration files.

- `Dockerfile`: Docker configuration for building the Go application.

- `docker-compose.yml`: Docker Compose configuration for setting up services.

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/rawichblue/go-rest-ful-api-with-gin.git
