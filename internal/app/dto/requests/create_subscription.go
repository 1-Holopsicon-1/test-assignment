package requests

import "time"

type CreateSubscriptionRequest struct {
	ServiceName string     `json:"service_name" example:"Netflix"`
	Price       int        `json:"price" example:"999"`
	UserID      string     `json:"user_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	StartDate   time.Time  `json:"start_date" example:"2024-01-01T00:00:00Z"`
	EndDate     *time.Time `json:"end_date,omitempty" example:"2024-12-31T23:59:59Z"`
}
