package service

import (
	"billing"
	"billing/pkg/repository"
)

type BillingService struct {
	repo repository.Billing
}

func NewTodoListService(repo repository.Billing) *BillingService {
	return &BillingService{repo: repo}
}

func (s *BillingService) AddMoney(acc billing.Account) error {
	return s.repo.AddMoney(acc)
}

func (s *BillingService) Transfer(idA, idB int) error {
	return s.repo.Transfer(idA, idB)
}

func (s *BillingService) Balance(id int) (int, error) {
	return s.repo.Balance(id)
}
