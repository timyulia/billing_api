package service

import (
	"billing"
	"billing/pkg/repository"
)

type Billing interface {
	AddMoney(acc billing.Account) error
	Transfer(trans billing.TransferInfo) error
	Balance(id int) (int, error)
	GetAllAccs() ([]billing.Account, error)
}

type Service struct {
	Billing
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Billing: NewBillingService(r.Billing),
	}
}
