
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
    "fmt"
    "log"
    "github.com/nikhilryan/go-featuristic/config"
    "github.com/nikhilryan/go-featuristic/internal/models"
    "github.com/nikhilryan/go-featuristic/internal/services"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
    cfg, err := config.LoadConfig(".")
    if err != nil {
        log.Fatalf("could not load config: %v", err)
    }

    dsn := config.GetDSN(cfg)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    db.AutoMigrate(&models.FeatureFlag{})

    cacheService := services.NewCacheService(cfg.CacheHost + ":" + cfg.CachePort)
    featureFlagService := services.NewFeatureFlagService(db, cacheService)

    stringFlag := &models.FeatureFlag{
        Namespace: "test",
        Key:       "stringFeature",
        Value:     "example string",
        Type:      "string",
    }
    err = featureFlagService.CreateFlag(stringFlag)
    if err != nil {
        log.Fatalf("failed to create feature flag: %v", err)
    }

    value, err := featureFlagService.GetFlagValue("test", "stringFeature")
    if err != nil {
        log.Fatalf("failed to get feature flag value: %v", err)
    }
    fmt.Printf("Feature flag value: %v\n", value)
}
```

### Example for String Array

```go
package main

import (
    "fmt"
    "log"
    "encoding/json"
    "github.com/nikhilryan/go-featuristic/config"
    "github.com/nikhilryan/go-featuristic/internal/models"
    "github.com/nikhilryan/go-featuristic/internal/services"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
    cfg, err := config.LoadConfig(".")
    if err != nil {
        log.Fatalf("could not load config: %v", err)
    }

    dsn := config.GetDSN(cfg)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    db.AutoMigrate(&models.FeatureFlag{})

    cacheService := services.NewCacheService(cfg.CacheHost + ":" + cfg.CachePort)
    featureFlagService := services.NewFeatureFlagService(db, cacheService)

    stringArray := []string{"feature1", "feature2", "feature3"}
    stringArrayJSON, err := json.Marshal(stringArray)
    if err != nil {
        log.Fatalf("failed to marshal string array: %v", err)
    }
    stringArrayFlag := &models.FeatureFlag{
        Namespace: "test",
        Key:       "stringArrayFeature",
        Value:     string(stringArrayJSON),
        Type:      "stringArray",
    }
    err = featureFlagService.CreateFlag(stringArrayFlag)
    if err != nil {
        log.Fatalf("failed to create feature flag: %v", err)
    }

    value, err := featureFlagService.GetFlagValue("test", "stringArrayFeature")
    if err != nil {
        log.Fatalf("failed to get feature flag value: %v", err)
    }
    fmt.Printf("Feature flag value: %v\n", value)
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.
