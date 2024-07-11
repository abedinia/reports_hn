# Customer Reports API

This project is a Go-based API to collect and manage customer reports containing sensitive information. It uses Gin for the web framework, Cobra for command-line interface, Viper for configuration, Logrus for logging, and PostgreSQL for the database.

## Features

- User authentication with JWT tokens
- CRUD operations for reports
- Rate limiting to prevent abuse
- Migration and seeding for database setup

## Prerequisites

- Go 1.22+
- PostgreSQL (Dockerfile and docker-compose available on repo for testing)

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/abedinia/reports_hn_app.git
    cd reports_hn_app
    ```

2. Install dependencies:

    ```bash
    go mod download
    ```

3. modify file `config.yaml` for new env:


## Usage
#### check Makefile to run for migrate, seed, report, tests, and so on

### Running the Application

To start the application, run:

```bash
go run cmd/main.go report
```

```bash
go run cmd/main.go migrate
```

### Seed
```bash
go run cmd/main.go seed
```

### Login

```bash
curl --location 'localhost:8000/login' \
--header 'Content-Type: application/json' \
--data '{
    "username":"your_username",
    "password":"your_password"
}'
```

### Create Report

```bash
curl --location 'http://localhost:8000/reports' \
--header 'Authorization: Bearer your_jwt_token' \
--header 'Content-Type: application/json' \
--data '{
    "dataset_id": 1,
    "name": "name of product",
    "type": 1,
    "reason": "here is mine"
}'
```


### List Report

```bash
curl --location 'http://localhost:8000/reports' \
--header 'Authorization: Bearer your_jwt_token'
```