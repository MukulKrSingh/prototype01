# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial project structure setup
- GraphQL schema definitions
- MongoDB integration through MCP
- Authentication system with JWT
- Product catalog management
- Order processing system
- User management functionality

## [0.1.0] - 2025-05-18

### Added
- Project initialization
- Basic directory structure
- README.md with setup instructions
- PROJECT.md with Go concepts explanation
- CHANGELOG.md for version tracking
- Initial Go module setup
- GraphQL API foundation with gqlgen
- MongoDB connection configuration
- Basic domain models (User, Product, Order)
- Docker configuration for development environment

### Security
- JWT-based authentication system
- Password hashing with bcrypt
- Role-based access control for API endpoints

---

## Version Guidelines

### Version Format

- **MAJOR** version for incompatible API changes
- **MINOR** version for new functionality in a backwards compatible manner
- **PATCH** version for backwards compatible bug fixes

### Commit Message Format

- **feat**: A new feature
- **fix**: A bug fix
- **docs**: Documentation only changes
- **style**: Changes that do not affect the meaning of the code
- **refactor**: A code change that neither fixes a bug nor adds a feature
- **perf**: A code change that improves performance
- **test**: Adding missing tests
- **chore**: Changes to the build process or auxiliary tools

---

## Release Process

1. Update CHANGELOG.md with changes
2. Update version number in relevant files
3. Create a git tag for the version
4. Push to GitHub
5. Create a GitHub release
6. Deploy to production environment
