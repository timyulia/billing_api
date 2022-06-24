package service

import (
	"billing"
	"billing/pkg/repository"
	"errors"
	"fmt"
)

type BillingService struct {
	repo repository.Billing
}

func NewBillingService(repo repository.Billing) *BillingService {
	return &BillingService{repo: repo}
}

func (s *BillingService) AddMoney(acc billing.Account) error {
	return s.repo.AddMoney(acc)
}

func (s *BillingService) Transfer(trans billing.TransferInfo) error {
	fmt.Println("hey")
	if trans.Amount < 0 {
		return errors.New("negative transfer amount")
	}
	fmt.Println(trans.Amount)
	return s.repo.Transfer(trans)
}

func (s *BillingService) Balance(id int) (int, error) {
	return s.repo.Balance(id)
}

func (s *BillingService) GetAllAccs() ([]billing.Account, error) {
	return s.repo.GetAllAccs()
}
