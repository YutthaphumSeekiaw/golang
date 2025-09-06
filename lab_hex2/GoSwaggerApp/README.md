# GoSwaggerApp

GoSwaggerApp is a simple Go application that provides a RESTful API for managing items. This project demonstrates how to set up an HTTP server, define routes, and handle requests using Go.

## Project Structure

```
GoSwaggerApp
├── cmd
│   └── main.go          # Entry point of the application
├── docs
│   └── swagger.yaml     # Swagger documentation for the API
├── internal
│   ├── handlers
│   │   └── items.go     # HTTP handler functions for managing items
│   └── models
│       └── item.go      # Data model for an item
├── go.mod               # Module definition for the Go project
├── go.sum               # Checksums for module dependencies
└── README.md            # Documentation for the project
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd GoSwaggerApp
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run cmd/main.go
   ```

4. **Access the API:**
   The API will be available at `http://localhost:8080/items`.

## Usage Examples

- **Get all items:**
  ```
  GET /items
  ```

- **Create a new item:**
  ```
  POST /items
  Content-Type: application/json

  {
      "name": "New Item"
  }
  ```

- **Get a specific item:**
  ```
  GET /items/{id}
  ```

- **Update an item:**
  ```
  PUT /items/{id}
  Content-Type: application/json

  {
      "name": "Updated Item"
  }
  ```

- **Delete an item:**
  ```
  DELETE /items/{id}
  ```

## API Documentation

For detailed API documentation, refer to the `docs/swagger.yaml` file. You can use Swagger UI to visualize and interact with the API.