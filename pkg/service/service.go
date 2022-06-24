package service

import (
	"billing"
	"billing/pkg/repository"
)

type Billing interface {
	AddMoney(acc billing.Account) error
	Transfer(idA, idB int) error
	Balance(id int) (int, error)
}

type Service struct {
	Billing
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Billing: r.Billing,
	}
}
