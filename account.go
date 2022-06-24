package billing

type Account struct {
	Id     int `json:"id" db:"id"`
	Amount int `json:"amount" db:"amount"`
}
