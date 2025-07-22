package repository

import (
	model "test-assignment/internal/app/models"
	"time"

	"gorm.io/gorm"
)

type SubscriptionRepo struct {
	DB *gorm.DB
}

func (r *SubscriptionRepo) Create(sub *model.Subscription) error {
	return r.DB.Create(sub).Error
}

func (r *SubscriptionRepo) GetAll() ([]model.Subscription, error) {
	var subs []model.Subscription
	err := r.DB.Find(&subs).Error
	return subs, err
}

func (r *SubscriptionRepo) GetByID(id uint) (*model.Subscription, error) {
	var sub model.Subscription
	err := r.DB.First(&sub, id).Error
	return &sub, err
}

func (r *SubscriptionRepo) Update(sub *model.Subscription) error {
	return r.DB.Save(sub).Error
}

func (r *SubscriptionRepo) Delete(id uint) error {
	return r.DB.Delete(&model.Subscription{}, id).Error
}

func (r *SubscriptionRepo) TotalCost(userID string, serviceName string, from, to time.Time) (int, error) {
	var total int
	q := r.DB.Model(&model.Subscription{}).
		Where("start_date >= ? AND (end_date <= ? OR end_date IS NULL)", from, to)

	if userID != "" {
		q = q.Where("user_id = ?", userID)
	}
	if serviceName != "" {
		q = q.Where("service_name = ?", serviceName)
	}

	err := q.Select("SUM(price)").Scan(&total).Error
	return total, err
}
