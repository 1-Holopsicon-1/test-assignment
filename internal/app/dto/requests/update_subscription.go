package requests

import "time"

type UpdateSubscriptionRequest struct {
	ServiceName string     `json:"service_name" example:"Netflix Premium"`
	Price       int        `json:"price" example:"1299"`
	StartDate   time.Time  `json:"start_date" example:"2024-01-01T00:00:00Z"`
	EndDate     *time.Time `json:"end_date,omitempty" example:"2024-12-31T23:59:59Z"`
}
