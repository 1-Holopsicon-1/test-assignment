package repository

import (
	model "test-assignment/internal/app/models"
	"time"
)

type SubscriptionRepository interface {
	Create(sub *model.Subscription) error
	GetAll() ([]model.Subscription, error)
	GetByID(id uint) (*model.Subscription, error)
	Update(sub *model.Subscription) error
	Delete(id uint) error
	TotalCost(userID string, serviceName string, from, to time.Time) (int, error)
}
