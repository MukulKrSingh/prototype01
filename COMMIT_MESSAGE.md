refactor: migrate from standard HTTP to Gin framework

This commit migrates the API from using the standard Go HTTP package to using the Gin framework.
The migration provides several advantages including better middleware support, simpler routing, 
and improved error handling.

Major changes:

1. Add Gin framework as a dependency
2. Create Gin-specific middleware in internal/middleware/gin_middleware.go
3. Update context handling in auth package for Gin compatibility  
4. Modify GraphQL handler to work with Gin
5. Update main.go to use Gin router
6. Create documentation for the Gin migration
7. Update middleware to support both standard HTTP and Gin
8. Add new targets in Makefile for Gin setup
9. Update project documentation to reflect the migration

This migrates the application to Gin while maintaining backward compatibility
where possible. Refer to docs/gin-migration.md for more details on the migration.
