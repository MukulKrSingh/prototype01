// Package resolvers contains GraphQL resolver implementations
package resolvers

// No imports needed in this file

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// This file contains the base resolver setup and dependency injection.
// The actual Resolver struct is defined in resolvers.go

// ResolverRoot provides root resolver instance for interface implementation
type ResolverRoot interface {
	Query() QueryResolver
	Mutation() MutationResolver
}

// Query returns a QueryResolver implementation
func (r *Resolver) Query() QueryResolver {
	return QueryResolver{r}
}

// Mutation returns a MutationResolver implementation
func (r *Resolver) Mutation() MutationResolver {
	return MutationResolver{r}
}
