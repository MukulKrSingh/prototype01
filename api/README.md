# GraphQL API Directory Structure

This directory contains the GraphQL schema definitions for the e-commerce backend.

## Organization

- `api/graphql/schema.graphql`: The main GraphQL schema file used by gqlgen
- `api/schema/`: This directory is reserved for future schema organization when we split the schema into multiple files

## Schema Development Flow

1. Edit the schema files in `api/graphql/`
2. Run `make generate` to generate Go code from the schema
3. Implement the resolvers in `internal/api/resolvers/`
4. Test your implementation using the GraphQL playground

## Schema Style Guidelines

1. Use clear, descriptive names for types, fields, and operations
2. Document all types and fields with comments
3. Group related fields and types together
4. Use custom scalars for specialized data types
5. Use enums for fields with a fixed set of values
6. Properly define input types for mutations
7. Follow a consistent naming convention:
   - Types: PascalCase (e.g., `Product`)
   - Fields: camelCase (e.g., `productName`)
   - Enums: SCREAMING_SNAKE_CASE (e.g., `PRODUCT_TYPE`)
   - Directives: camelCase (e.g., `@hasRole`)

## Future Schema Organization (Step 2)

In Step 2, we'll organize the schema into multiple files:
- `schema.graphql`: Main schema file with imports
- `types/`: Directory for GraphQL type definitions
  - `product.graphql`: Product-related types
  - `user.graphql`: User-related types
  - `order.graphql`: Order-related types
- `queries.graphql`: Query operations
- `mutations.graphql`: Mutation operations
- `directives.graphql`: Custom directive definitions
