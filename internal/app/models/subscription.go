package model

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          uint       `gorm:"primaryKey" json:"id" example:"1" swaggertype:"integer"`
	ServiceName string     `json:"service_name" example:"Netflix" binding:"required"`
	Price       int        `json:"price" example:"999" binding:"required,min=1" swaggertype:"integer"`
	UserID      uuid.UUID  `json:"user_id" example:"550e8400-e29b-41d4-a716-446655440000" binding:"required"`
	StartDate   time.Time  `json:"start_date" example:"2024-01-01T00:00:00Z" binding:"required"`
	EndDate     *time.Time `json:"end_date,omitempty" example:"2024-12-31T23:59:59Z"`
}
