# 🛒 Go E-Commerce Backend

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.22.5-00ADD8?style=for-the-badge&logo=go)
![GraphQL](https://img.shields.io/badge/GraphQL-E10098?style=for-the-badge&logo=graphql)
![MongoDB](https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb)

**A modern, scalable e-commerce backend built with Go, GraphQL, and MongoDB**

</div>

## ✨ Features

- **🔐 Authentication & Authorization** - Secure API with JWT
- **📦 Product Management** - Complete CRUD for products and categories
- **🛍️ Order Processing** - End-to-end order fulfillment
- **👤 User Management** - User accounts, profiles, and preferences
- **🔍 Search & Filtering** - Advanced product search capabilities
- **📊 Analytics** - Basic sales and performance metrics
- **📱 GraphQL API** - Modern, flexible API powered by gqlgen

## 📋 Prerequisites

- [Go](https://go.dev/doc/install) 1.22 or higher
- [GVM](https://github.com/moovweb/gvm) for Go version management
- [MongoDB](https://www.mongodb.com/docs/manual/installation/)
- [Git](https://git-scm.com/downloads)

## 🚀 Quick Start

### Clone and Setup

```bash
# Clone the repository
git clone https://github.com/yourusername/go-ecommerce.git
cd go-ecommerce

# Use GVM to ensure correct Go version
gvm use go1.22.5

# Install dependencies
go mod download
```

### Configuration

Create a `.env` file in the project root:

```env
# Server Configuration
PORT=8080
ENV=development

# MongoDB Configuration
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=ecommerce

# JWT Configuration
JWT_SECRET=your-secret-key
JWT_EXPIRATION=24h
```

### Run the Server

```bash
# Standard run mode
make run

# Development mode with GraphQL playground
make dev

# Development mode with hot reload
make dev-live

# Using the development script
./scripts/dev-server.sh
```

The GraphQL API will be available at `http://localhost:8080/graphql`  
The GraphQL Playground will be available at `http://localhost:8080/playground`

### GraphQL Development with gqlgen

This project uses [gqlgen](https://github.com/99designs/gqlgen) for implementing GraphQL API:

```bash
# Generate GraphQL code from schema
make generate

# If you modify the schema in api/graphql/*.graphql
# you'll need to regenerate the code:
go run github.com/99designs/gqlgen generate
```

For more information about GraphQL development:
- See [docs/graphql.md](docs/graphql.md) for best practices
- See [docs/graphql-examples.md](docs/graphql-examples.md) for API usage examples

### Build for Production

```bash
make build
# Binary will be available in ./bin directory
```

## 📁 Project Structure

```
go-ecommerce/
├── cmd/                  # Application entry points
│   └── server/           # API server
├── internal/             # Private application code
│   ├── api/              # GraphQL API implementation
│   ├── config/           # Configuration handlers
│   ├── domain/           # Business logic and entities
│   └── repository/       # Database interactions
├── pkg/                  # Public libraries
│   ├── logger/           # Logging utilities
│   └── utils/            # Common utilities
├── api/                  # GraphQL schema definitions
├── scripts/              # Utility scripts
└── docs/                 # Documentation
```

## 📈 GraphQL API

Access the GraphQL playground at `http://localhost:8080/playground` (in development mode) to explore available queries and mutations.

### Example Queries

```graphql
# Get all products
query {
  products {
    id
    name
    description
    price
    category {
      name
    }
  }
}

# Create a new product
mutation {
  createProduct(input: {
    name: "Wireless Headphones"
    description: "High quality wireless headphones"
    price: 99.99
    categoryId: "category-id"
  }) {
    id
    name
  }
}
```

## 🔧 Development

### Generate GraphQL Code

```bash
# After updating schema in api/schema.graphql
make generate
```

### Run Tests

```bash
make test
```

## 📚 Documentation

For more detailed information:

- [Project Understanding Guide](PROJECT.md) - Learn about Go concepts used in this project
- [API Documentation](docs/API.md) - Detailed API documentation
- [Database Schema](docs/SCHEMA.md) - Database structure and relationships
- [Change Log](CHANGELOG.md) - Version history and updates

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
