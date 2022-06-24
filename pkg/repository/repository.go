package repository

import (
	"billing"
	"github.com/jmoiron/sqlx"
)

type Billing interface {
	AddMoney(acc billing.Account) error
	Transfer(idA, idB int) error
	Balance(id int) (int, error)
}

type Repository struct {
	Billing
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Billing: NewBillingPostgres(db)}
}
