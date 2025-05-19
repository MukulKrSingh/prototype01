// Package models contains the domain models used throughout the application
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BaseModel provides common fields for all models
type BaseModel struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

// BeforeCreate is a hook that sets CreatedAt and UpdatedAt before creating a record
func (m *BaseModel) BeforeCreate() {
	now := time.Now()
	m.CreatedAt = now
	m.UpdatedAt = now

	// Generate ObjectID if not already set
	if m.ID.IsZero() {
		m.ID = primitive.NewObjectID()
	}
}

// BeforeUpdate is a hook that updates UpdatedAt before updating a record
func (m *BaseModel) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}
