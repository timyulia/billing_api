package repository

import (
	"billing"
	"errors"
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
	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE id=$1", accountsTable)
	var count int
	err := r.db.Get(&count, query, acc.Id)
	if err != nil {
		return err
	}
	exists, err := r.checkExistence(acc.Id)
	if err != nil {
		return err
	}
	switch {
	case !exists && acc.Amount > 0:
		query = fmt.Sprintf("INSERT INTO %s  VALUES ($1, $2)", accountsTable)
		_, err := r.db.Exec(query, acc.Id, acc.Amount)
		return err
	case exists && acc.Amount > 0:
		query = fmt.Sprintf("UPDATE %s  SET amount=(amount+$1)  WHERE id=$2", accountsTable)
		_, err := r.db.Exec(query, acc.Amount, acc.Id)
		return err
	case exists && acc.Amount < 0:
		ok, err := r.checkBalance(acc.Id, (-1)*acc.Amount)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("not enough money")
		}
		query = fmt.Sprintf("UPDATE %s  SET amount=(amount+$1)  WHERE id=$2", accountsTable)
		_, err = r.db.Exec(query, acc.Amount, acc.Id)
		return err
	default: //count == 0 && acc.Amount < 0:
		return errors.New("there is no such account yet")
	}

}
func (r *BillingPostgres) Transfer(trans billing.TransferInfo) error {
	if ok, _ := r.checkExistence(trans.IdA); !ok {
		return errors.New("account A does not exist")
	}
	if ok, _ := r.checkExistence(trans.IdB); !ok {
		return errors.New("account B does not exist")
	}
	ok, err := r.checkBalance(trans.IdA, trans.Amount)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("not enough money")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	minusQuery := fmt.Sprintf("UPDATE %s SET amount=(amount-$1) WHERE id=$2", accountsTable)
	_, err = tx.Exec(minusQuery, trans.Amount, trans.IdA)
	if err != nil {
		tx.Rollback()
		return err
	}

	plusQuery := fmt.Sprintf("UPDATE %s SET amount=(amount+$1) WHERE id=$2", accountsTable)
	_, err = tx.Exec(plusQuery, trans.Amount, trans.IdB)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()

}
func (r *BillingPostgres) Balance(id int) (int, error) {
	query := fmt.Sprintf("SELECT amount from %s WHERE id=$1", accountsTable)
	var moneyAmount int
	err := r.db.Get(&moneyAmount, query, id)
	return moneyAmount, err
}

func (r *BillingPostgres) checkBalance(id, amount int) (bool, error) {
	query := fmt.Sprintf("SELECT amount FROM %s WHERE id=$1", accountsTable)
	var currAmount int
	err := r.db.Get(&currAmount, query, id)
	return currAmount > amount, err
}

func (r *BillingPostgres) checkExistence(id int) (bool, error) {
	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE id=$1", accountsTable)
	var count int
	err := r.db.Get(&count, query, id)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
