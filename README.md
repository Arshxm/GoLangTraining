# golangtraining

Training repository for Go programming exercises and projects from quera.org.

## Features

- HTTP API development with Gin and standard library
- Database operations with GORM and PostgreSQL
- WebSocket implementations
- Authentication and middleware patterns
- SQL query exercises

## Requirements

- Go 1.19 or later

## Installation

Each project is a standalone module. Navigate to a project directory and run:

```bash
cd CalcAPI
go mod download
go run main.go
```

## Usage

### Example: Calculator API

```go
package main

import "net/http"

func main() {
    server := NewServer("8000")
    server.Start()
}

// GET /add?numbers=1,2,3
// GET /sub?numbers=10,5,2
```

### Example: Bank System

```go
package main

func main() {
    account := NewSavingsAccount()
    Deposit(account, 1000)
    balance := CheckBalance(account)
    // balance == 1000
}
```

## Project Structure

```
.
├── AirplaneAgency/          # Agency management system
├── Armstrong/               # Armstrong number calculations
├── BankSystem/             # Banking operations with multiple account types
├── CalcAPI/                # HTTP API for arithmetic operations
├── chatroom_ws/            # WebSocket chatroom implementation
├── cryptoCurrency/         # Cryptocurrency operations
├── firstAPI/               # Basic HTTP API example
├── GameServer/             # Game server implementation
├── libraryAPI/             # Library management API
├── QueraLastTrainings/     # Recent quera.org training sessions
├── ReverseAuth/            # Authentication middleware with reversed password check
├── TLDR/                   # Database caching implementation
├── ToDo/                   # Todo application with repository pattern
├── ValidationTraining/     # Input validation exercises
└── [other training projects]
```

## Testing

Run tests for all projects:

```bash
go test ./...
```

Run tests for a specific project:

```bash
cd BankSystem
go test
```