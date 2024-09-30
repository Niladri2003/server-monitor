package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Alert struct {
	MetricType       string             `json:"metric_type" validate:"required,max=50"` // Required, max 50 characters
	Threshold        float64            `json:"threshold" validate:"required,gte=0"`    // Required, must be greater than or equal to 0
	User             primitive.ObjectID `json:"user" validate:"required"`               // Required field
	Interval         int64              `json:"interval" validate:"required,gt=0"`      // Required, must be greater than 0
	CreatedAt        time.Time          `json:"created_at"`                             // No validation for auto-generated fields
	Status           string             `json:"status" validate:"required"`             // Required, must be either "active" or "inactive"
	AlertMsg         bool               `json:"alert_msg" validate:"required"`          // Required field (boolean)
	NotificationType primitive.ObjectID `json:"notification" validate:"required"`       // Required field
	Server           primitive.ObjectID `json:"server" validate:"required"`             // Required field
}

type NotificationType struct {
	CommMedium string `json:"comm_medium" validate:"required,oneof=email sms webhook"` // Required, must be one of the specified mediums
	Verified   bool   `json:"verified" validate:"required"`                            // Required field (boolean)
	Value      string `json:"value" validate:"required,max=255"`                       // Required, max length 255 characters
}
