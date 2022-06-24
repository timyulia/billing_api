package repository

import (
	"billing"
	"github.com/jmoiron/sqlx"
)

type Billing interface {
	AddMoney(acc billing.Account) error
	Transfer(trans billing.TransferInfo) error
	Balance(id int) (int, error)
	GetAllAccs() ([]billing.Account, error)
}

type Repository struct {
	Billing
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Billing: NewBillingPostgres(db)}
}
