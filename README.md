# 🛒 Go E-Commerce Backend

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.22.5-00ADD8?style=for-the-badge&logo=go)
![GraphQL](https://img.shields.io/badge/GraphQL-E10098?style=for-the-badge&logo=graphql)
![MongoDB](https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb)
![Gin](https://img.shields.io/badge/Gin-00ADD8?style=for-the-badge&logo=go)

**A modern, scalable e-commerce backend built with Go, Gin, GraphQL, and MongoDB**

</div>

## ✨ Features

- **🔐 Authentication & Authorization** - Secure API with JWT
- **📦 Product Management** - Complete CRUD for products and categories
- **🛍️ Order Processing** - End-to-end order fulfillment
- **👤 User Management** - User accounts, profiles, and preferences
- **🔍 Search & Filtering** - Advanced product search capabilities
- **📊 Analytics** - Basic sales and performance metrics
- **📱 GraphQL API** - Modern, flexible API powered by gqlgen
- **🧪 Apollo Studio Integration** - Enhanced API testing and exploration
- **🚀 Gin Framework** - High-performance HTTP web framework
- **⚡ Automatic Persisted Queries** - Optimized GraphQL requests with APQ caching

## 📂 Project Structure

This project follows a clean architecture approach with clear separation of concerns:

> **Important**: The `COPILOT_INSTRUCTIONS.md` file contains critical development guidelines and roadmap information. Never delete this file as it provides essential context for the project development.

- `/api` - GraphQL schema definitions
- `/apollo` - Apollo Studio configurations and example queries
- `/cmd` - Main application entry points
- `/docs` - Project documentation
- `/internal` - Internal packages (domain logic, repositories, services)
- `/pkg` - Reusable packages that could be imported by other projects
- `/scripts` - Utility scripts for development and operations

For a detailed explanation of each component, see the [Project Structure](PROJECT_STRUCTURE.md) document.

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

# Setup Gin framework
make gin-setup
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

### Testing with Apollo Studio

This project supports Apollo Studio for GraphQL exploration and testing:

```bash
# Generate Apollo Studio configuration
./scripts/apollo-config.sh
```

To use Apollo Studio:
1. Start your server with `make dev`
2. Open [Apollo Studio Explorer](https://studio.apollographql.com/sandbox/explorer)
3. Enter your GraphQL endpoint: `http://localhost:8080/graphql`
4. Start building and testing queries

See [docs/apollo-studio.md](docs/apollo-studio.md) for detailed instructions.

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
prototype01/
├── api/                  # GraphQL schema definitions
│   └── graphql/          # Schema definitions (*.graphql files)
├── apollo/               # Apollo Studio configuration
│   └── queries/          # Example GraphQL queries and mutations
├── cmd/                  # Application entry points
│   └── server/           # API server
├── docs/                 # Documentation
├── internal/             # Private application code
│   ├── api/              # GraphQL API implementation
│   ├── auth/             # Authentication and authorization
│   ├── config/           # Configuration handlers
│   ├── domain/           # Business logic and entities
│   ├── middleware/       # HTTP middleware components
│   ├── repository/       # Database interactions
│   └── service/          # Application services
├── logs/                 # Application logs
├── pkg/                  # Public libraries
│   ├── logger/           # Logging utilities
│   ├── utils/            # Common utilities
│   └── validator/        # Data validation utilities
└── scripts/              # Utility scripts
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
- [Project Structure](PROJECT_STRUCTURE.md) - Detailed explanation of the folder structure
- [Setup Summary](SETUP_SUMMARY.md) - Summary of completed setup tasks
- [GraphQL Guide](docs/graphql.md) - GraphQL implementation with gqlgen
- [GraphQL Examples](docs/graphql-examples.md) - Sample GraphQL operations and patterns
- [GraphQL Naming Conventions](docs/graphql-naming-conventions.md) - Standards for naming operations
- [Apollo Studio Guide](docs/apollo-studio.md) - Instructions for using Apollo Studio
- [Change Log](CHANGELOG.md) - Version history and updates

Sample GraphQL queries and mutations can be found in the `/apollo/queries/` directory.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
