package types

// Money представляет собой денеэную сумму в минималных единиц (центы, копейки, дирамы и т.д)
type Money int64

// Currency представляет код валюты
type Currency string

// Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

// PAN представляет номер карты
type PAN string

// Card представляет информацию о платежной карте
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Payment    Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}