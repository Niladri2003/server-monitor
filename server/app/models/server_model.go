package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Server struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	ServerID      string             `bson:"server_id"`      // Unique server identifier
	Name          string             `bson:"name"`           // Server name
	Description   string             `bson:"description"`    // Server description
	UserID        primitive.ObjectID `bson:"user"`           // Reference to the user who added the server
	APIKey        string             `bson:"api_key"`        // Unique API key for the server
	CloudProvider string             `bson:"cloud_provider"` // Server IP address
	CreatedAt     time.Time          `bson:"created_at"`     // Timestamp of when the server was added
	UpdatedAt     time.Time          `bson:"updated_at"`     // Timestamp of the last update
}
