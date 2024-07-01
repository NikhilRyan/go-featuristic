
# Go-Featuristic

Go-Featuristic is a Golang library for managing feature flags and rollout mechanisms. It supports various data types and integrates with MySQL, PostgreSQL, Redis, and Memcached. The library allows you to store feature flags in a database and provides a caching layer for efficient retrieval.

## Features

- Manage feature flags with different data types (int, float, string, array)
- Group feature flags under namespaces
- Store feature flags in MySQL or PostgreSQL
- Caching layer using Redis or Memcached
- Gradual rollout of feature flags
- RESTful API for managing feature flags

## Installation

To install the Go-Featuristic library, add the module to your `go.mod` file:

```sh
go get github.com/nikhilryan/go-featuristic
```

## Usage

### Configuration

Create an `app.env` file in the root directory of your project:

```plaintext
DB_HOST=localhost
DB_PORT=5432
DB_USER=myuser
DB_PASSWORD=mypassword
DB_NAME=mydatabase
CACHE_HOST=localhost
CACHE_PORT=6379
SERVER_PORT=8080
```

### Example Project

#### Initialize a New Go Module

```sh
mkdir my-project
cd my-project
go mod init github.com/yourusername/my-project
```

#### Add Go-Featuristic as a Dependency

Edit the `go.mod` file to include the Go-Featuristic module:

```go
module github.com/yourusername/my-project

go 1.21

require (
    github.com/nikhilryan/go-featuristic latest
)
```

Run `go mod tidy` to download the dependencies:

```sh
go mod tidy
```

#### Import and Use the Library

Create a main.go file to use the Go-Featuristic library:

```go
package main

import (
    "log"
    "net/http"
    "github.com/nikhilryan/go-featuristic/config"
    "github.com/nikhilryan/go-featuristic/internal/services"
    "github.com/nikhilryan/go-featuristic/api/routers"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig(".")
    if err != nil {
        log.Fatalf("could not load config: %v", err)
    }

    // Connect to the database
    dsn := config.GetDSN(cfg)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    // Initialize services
    cacheService := services.NewCacheService(cfg.CacheHost + ":" + cfg.CachePort)
    featureFlagService := services.NewFeatureFlagService(db, cacheService)

    // Setup router
    r := routers.SetupRouter(featureFlagService)

    // Start the server
    log.Println("Server running on port", cfg.ServerPort)
    log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
}
```

#### Run the Project

Run the project using the `go run` command:

```sh
go run main.go
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.
