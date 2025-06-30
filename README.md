# Greenlight

An implementation of the Greenlight API from the "Let's Go Further" book.

## Overview

This project provides a backend API for managing a collection of movies. It is inspired by the "Greenlight" example from the "Let's Go Further" book.

## Prerequisites

- Go (1.18 or later)
- PostgreSQL
- [golang-migrate](https://github.com/golang-migrate/migrate) (for running migrations)

## Features

- RESTful API for CRUD operations on movies
- PostgreSQL database schema with migrations
- Example movie fields: title, year, runtime, genres

## Database

The database schema is managed via SQL migrations in the `migrations/` directory. The main table is `movies` with fields for ID, creation timestamp, title, year, runtime, genres, and version.

## Getting Started

1. Clone the repository.
2. Set up a PostgreSQL database and user. You can use the following credentials (or choose your own):
   - **User:** greenlight
   - **Password:** pa55word
   - **Database:** greenlight

   Example setup script (run as a superuser, e.g. `psql -U postgres`):

   ```sql
   CREATE USER greenlight WITH PASSWORD 'pa55word';
   CREATE DATABASE greenlight OWNER greenlight;
   GRANT ALL PRIVILEGES ON DATABASE greenlight TO greenlight;
   ```

3. Run the migrations in `migrations/` to create the required tables.  
   **Note:** Make sure you have [golang-migrate](https://github.com/golang-migrate/migrate) installed before running the migration command.

   ```
   migrate -database "postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable" -path migrations up
   ```

4. Set your database connection string in an environment variable:
   ```
   export GREENLIGHT_DB_DSN="postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable"
   ```

5. Start the API server:
   ```
   go run ./cmd/api
   ```

## License

MIT License.

