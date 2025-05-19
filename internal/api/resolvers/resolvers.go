// Package resolvers contains the GraphQL resolvers
package resolvers

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// Version information
const (
	versionNumber = "0.1.0"
	buildEnv      = "development" // Will be configurable in Step 2
)

// Resolver is the base GraphQL resolver
type Resolver struct {
	DB *mongo.Client
}

// Query resolves all GraphQL queries
type QueryResolver struct{ *Resolver }

// Mutation resolves all GraphQL mutations
type MutationResolver struct{ *Resolver }

// Version represents version information resolver
type VersionResolver struct{ *Resolver }

// Ping resolves the ping query
func (r *QueryResolver) Ping(ctx context.Context) string {
	return "GraphQL Server is running! Current time: " + time.Now().Format(time.RFC3339)
}

// Version resolves the version query
func (r *QueryResolver) Version(ctx context.Context) *VersionResolver {
	return &VersionResolver{r.Resolver}
}

// Number returns the version number
func (r *VersionResolver) Number() string {
	return versionNumber
}

// BuildDate returns the build date
func (r *VersionResolver) BuildDate() time.Time {
	// For now we'll use current time
	// In production this would be set during the build process
	return time.Now()
}

// Environment returns the environment name
func (r *VersionResolver) Environment() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = buildEnv
	}
	return env
}

// Noop is a placeholder mutation that does nothing
func (r *MutationResolver) Noop(ctx context.Context) *bool {
	return nil
}
