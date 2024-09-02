# API Golang Project

## Overview

This is a Go project that uses the Fiber framework and PostgreSQL for managing user data. It includes basic CRUD operations and is set up using Docker and Docker Compose. The project also integrates JWT authentication and password hashing for secure login functionality.

## Features

- Basic CRUD operations for users
- JWT authentication with 1-minute token expiration
- Password hashing and verification
- Docker and Docker Compose setup
- PostgreSQL database integration
- Pagination and search functionality

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
