# GraphQL with gqlgen

This project uses [gqlgen](https://github.com/99designs/gqlgen) for building GraphQL servers in Go.

## Why gqlgen?

- **Code-first approach:** Define your schema in GraphQL SDL, and it generates type-safe Go code
- **Performance:** One of the most performant GraphQL implementations for Go
- **Type safety:** Uses Go's type system for schema validation
- **Directives support:** Allows the use of custom directives like @auth and @hasRole
- **Flexibility:** Easily customizable through configuration

## Best Practices

When working with gqlgen in this project, follow these best practices:

1. **Schema-first development:**
   - Define your GraphQL schema in `.graphql` files
   - Run code generation to create Go types and resolver interfaces
   - Implement the resolver interfaces

2. **Naming conventions:**
   - Follow the [GraphQL naming conventions](./graphql-naming-conventions.md) for all operations
   - Use consistent patterns for queries and mutations
   - Be descriptive and specific in operation names

3. **Resolver implementation:**
   - Group related resolvers in their own files
   - Use dependency injection for services and repositories 
   - Keep business logic in domain services, not in resolvers

3. **Handling N+1 query problems:**
   - Use DataLoaders to batch database queries
   - Consider caching for frequently accessed data

4. **Authentication & Authorization:**
   - Use the `@auth` and `@hasRole` directives
   - Implement middleware for JWT verification
   - Pass user context through the request context

## Current Setup

In Step 1 (Project Setup):
- We've set up the basic GraphQL schema in `api/graphql/schema.graphql`
- Created a configuration file at `gqlgen.yml`
- Set up placeholder generated code (to be replaced by actual generation in Step 2)
- Added GraphQL handler and playground handler using gqlgen libraries

## Future Implementation (Step 2)

In the next step, we will:
1. Expand the GraphQL schema with proper e-commerce entities
2. Run `go run github.com/99designs/gqlgen generate` to generate production-ready code
3. Implement resolvers for all GraphQL queries and mutations
4. Add middleware for authentication and other cross-cutting concerns
5. Set up proper error handling

## Example Usage

To use the GraphQL API in development:
1. Start the server with `make run`
2. Visit the GraphQL Playground at http://localhost:8080/playground
3. Try a simple query:
   ```graphql
   {
     ping
   }
   ```
   
To generate GraphQL code:
```bash
make generate
```

## Resources

- [gqlgen Documentation](https://gqlgen.com/)
- [GraphQL Specification](https://spec.graphql.org/)
