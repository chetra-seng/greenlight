# Greenlight

Greenlight is an implementation of the Greenlight API, inspired by the "Let's Go Further" book. It provides a RESTful API for managing a collection of movies.

## Features

- Add, update, delete, and retrieve movies
- Filter and search movies by various criteria
- Built-in data validation and error handling
- PostgreSQL database support with migrations

## Getting Started

### Prerequisites

- Go 1.18+
- PostgreSQL

### Setup

1. Clone the repository:
   ```sh
   git clone https://github.com/chetraseng/greenlight.git
   cd greenlight
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```


## Project Structure

- `cmd/api/` - Main application entrypoint and HTTP handlers
- `internal/data/` - Data models and database logic
- `internal/validator/` - Input validation helpers
- `migrations/` - Database migration scripts
