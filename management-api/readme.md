# Logto Management API SDK for Golang

This is the official Logto Management API SDK for Golang, which is generated using openapi-generator and provides convenient access to the Logto Management API from Go applications.

This SDK is built on top of the [Logto Management API](https://docs.logto.io/integrate-logto/interact-with-management-api), which is a powerful tool for managing your Logto resources. By using this SDK, you can manage your Logto resources programmatically, such as users, roles, and applications.

## Installation

```bash
go get github.com/logto-io/go/management-api
```

## Prerequisites

- Go 1.22 or higher
- A Logto instance (local or cloud)
- An application created in Logto with Management API scope

## Usage

Here's a simple example of how to use the SDK:

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/logto-io/go/management-api/client"
    "github.com/logto-io/go/management-api/logto"
)

func main() {
    // Initialize the client
    client := client.NewClient(&client.Config{
        AppId:     "<your_app_id>",
        AppSecret: "<your_app_secret>",
        Endpoint:  "<your_logto_endpoint>", // e.g., "https://your-tenant-id.logto.app"
        // for cloud users, must set the ManagementApiResourceURL to the default endpoint https://{your_tenant_id}.logto.app/api
        ManagementApiResourceURL: "<your_logto_management_api_resource_url>", // e.g., "https://your-tenant-id.logto.app/api"
        Debug:     true,
    }, nil)

    // Example: Get user information
    userId := "<user_id>"
    req := client.UsersAPI.GetUser(client.Context, userId)
    result, resp, err := req.Execute()
    if err != nil {
        log.Fatalf("API call failed: %v", err)
    }

    // Print the response
    fmt.Printf("Response status code: %d\n", resp.StatusCode)
    jsonBody, _ := json.MarshalIndent(result, "", "  ")
    fmt.Printf("Response body: %s\n", string(jsonBody))
}
```

## Usage With CacheStorage

The SDK provides a `CacheStorage` interface that allows you to implement your own storage solution for caching access tokens. This is particularly useful when you need to share access tokens across multiple instances of your application, such as in a distributed system using Redis.

### CacheStorage Interface

To implement your own cache storage, you need to implement the following interface:

```go
type CacheStorage interface {
    Get(key string) (string, error)
    Setex(key string, value string, seconds int) error
}
```

### Example Implementation with Redis

Here's an example of how to implement the `CacheStorage` interface using Redis:

```go
package svc

import "github.com/pkg/errors"

// RedisCacheStorage implements cache storage using Redis
type RedisCacheStorage struct {
    redis interface {
        Get(key string) (string, error)
        Setex(key string, value string, seconds int) error
    }
    prefix string
}

// NewRedisCacheStorage creates a new Redis cache storage
func NewRedisCacheStorageWithPrefix(redis interface {
    Get(key string) (string, error)
    Setex(key string, value string, seconds int) error
}, prefix string) *RedisCacheStorage {
    return &RedisCacheStorage{
        redis:  redis,
        prefix: prefix,
    }
}

// Get implements CacheStorage interface
func (r *RedisCacheStorage) Get(key string) (string, error) {
    result, err := r.redis.Get(r.prefix + key)
    if err != nil {
        return "", errors.Wrap(err, "RedisCacheStorage.Get")
    }
    return result, nil
}

// Setex implements CacheStorage interface
func (r *RedisCacheStorage) Setex(key string, value string, seconds int) error {
    err := r.redis.Setex(r.prefix+key, value, seconds)
    if err != nil {
        return errors.Wrap(err, "RedisCacheStorage.Setex")
    }
    return nil
}
```

### Using Custom CacheStorage with the Client

To use your custom cache storage implementation with the Logto client:

```go
// Initialize Redis client (implementation depends on your Redis library)
redisClient := initRedisClient()

// Create cache storage instance
cacheStorage := svc.NewRedisCacheStorageWithPrefix(redisClient, "logto:")

// Initialize Logto client with cache storage
client := client.NewClient(&client.Config{
    AppId:     "<your_app_id>",
    AppSecret: "<your_app_secret>",
    Endpoint:  "<your_logto_endpoint>",
    Debug:     true,
}, cacheStorage)
```

This setup allows the Logto client to store and retrieve access tokens using your Redis implementation, making it possible to share tokens across multiple application instances.

## Features

- Full Management API coverage
- Type-safe API calls
- Automatic token management
- Debug mode for troubleshooting

## Documentation

For detailed API documentation, please refer to:
- [Logto Management API](https://docs.logto.io/integrate-logto/interact-with-management-api)
- [Logto OpenAPI](https://openapi.logto.io/)

## Contributing Guidelines


### Generate OpenAPI SDK

You can generate the SDK by running the following command:

```bash
brew install openapi-generator # for macOS user
sh gen_openapi.sh
```