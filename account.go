package billing

type Account struct {
	Id     int `json:"id" db:"id"`
	Amount int `json:"amount" db:"amount"`
}

type TransferInfo struct {
	IdA    int `json:"idA"`
	IdB    int `json:"idB"`
	Amount int `json:"amount"`
}
