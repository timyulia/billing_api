package repository

import (
	"billing"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type BillingPostgres struct {
	db *sqlx.DB
}

func NewBillingPostgres(db *sqlx.DB) *BillingPostgres {
	return &BillingPostgres{db: db}
}

func (r *BillingPostgres) AddMoney(acc billing.Account) error {
	query := fmt.Sprintf("INSERT INTO %s  VALUES ($1, $2)", accountsTable)
	_, err := r.db.Exec(query, acc.Id, acc.Amount)
	return err
}
func (r *BillingPostgres) Transfer(idA, idB int) error {
	return nil
}
func (r *BillingPostgres) Balance(id int) (int, error) {
	query := fmt.Sprintf("SELECT amount from %s WHERE id=$1", accountsTable)
	var moneyAmount int
	err := r.db.Get(&moneyAmount, query, id)
	return moneyAmount, err
}
