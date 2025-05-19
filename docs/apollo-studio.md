# Using Apollo Studio with Your Go GraphQL API

This guide explains how to use Apollo Studio to test and explore your GraphQL API.

## What is Apollo Studio?

Apollo Studio is a powerful web IDE for GraphQL development that offers:

- Interactive query building
- Schema exploration
- Testing and debugging tools
- Request tracing
- Query performance monitoring

## Setup Instructions

### 1. Start Your API Server

```bash
make dev
```

Your API will be available at:
- GraphQL Endpoint: http://localhost:8080/graphql
- GraphQL Playground: http://localhost:8080/playground

### 2. Connect to Apollo Studio

#### Option 1: Use Apollo Studio Explorer (Recommended)

1. Open [Apollo Studio Explorer](https://studio.apollographql.com/sandbox/explorer)
2. Enter your GraphQL endpoint URL: `http://localhost:8080/graphql`
3. Start exploring your schema and building queries

#### Option 2: Generate Local Apollo Config (For Apollo CLI Users)

1. Run the config generator script:
   ```bash
   ./scripts/apollo-config.sh
   ```
2. Use Apollo CLI with the generated config

## Features Available in Apollo Studio

- **Schema Explorer**: Browse your API's types, fields, and documentation
- **Operation Builder**: Construct and test GraphQL operations with autocompletion
- **Variables Panel**: Define and reuse operation variables
- **Headers Panel**: Set HTTP headers for authentication
- **Response Panel**: View formatted JSON responses and errors
- **Explorer History**: Save and organize your operations

## Authentication with Apollo Studio

When testing endpoints that require authentication:

1. Click the "Headers" tab in Apollo Studio
2. Add your authorization header:
   ```json
   {
     "Authorization": "Bearer YOUR_TOKEN_HERE"
   }
   ```

## Performance Insights

Apollo Studio provides performance metrics for your operations:

- Field-level timing information
- Network latency analysis
- Query complexity score

## Troubleshooting

- **CORS Errors**: If you see CORS errors, make sure the `CORSMiddleware` is properly configured for the `/graphql` endpoint
- **Authentication Issues**: Verify your authorization tokens are correctly formatted
- **Schema Not Loading**: Check that introspection is enabled in your API

## Further Resources

- [Apollo Studio Documentation](https://www.apollographql.com/docs/studio/)
- [GraphQL Best Practices](https://graphql.org/learn/best-practices/)
