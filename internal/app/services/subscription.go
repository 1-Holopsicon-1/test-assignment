package services

import (
	model "test-assignment/internal/app/models"
	"test-assignment/internal/app/repository"
	"time"
)

type SubscriptionService struct {
	Repo repository.SubscriptionRepo
}

func (s *SubscriptionService) Create(sub *model.Subscription) error {
	return s.Repo.Create(sub)
}

func (s *SubscriptionService) GetAll() ([]model.Subscription, error) {
	return s.Repo.GetAll()
}

func (s *SubscriptionService) GetByID(id uint) (*model.Subscription, error) {
	return s.Repo.GetByID(id)
}

func (s *SubscriptionService) Update(sub *model.Subscription) error {
	return s.Repo.Update(sub)
}

func (s *SubscriptionService) Delete(id uint) error {
	return s.Repo.Delete(id)
}

func (s *SubscriptionService) TotalCost(userID, serviceName string, from, to time.Time) (int, error) {
	return s.Repo.TotalCost(userID, serviceName, from, to)
}
