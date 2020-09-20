package card

import (
	// "fmt"
	"bank/pkg/bank/types"
	
)

var card = types.Card{
	ID:         1000,
	PAN:        "5058 xxxx xxxx 0001",
	Balance:    0,
	Currency:   "currency",
	Color:      "color",
	Name:       "name",
	Active:     true,
	MinBalance: 0,
}

func Withdraw(card types.Card, amount types.Money) types.Card {

	if card.Active == false || amount > card.Balance || amount < 0 || amount > 20_000_00 {
		return card
	}
	card.Balance = card.Balance - amount

	return card
}

func Deposit(card *types.Card, amount types.Money) {

	if !(*card).Active {
		return
	}

	if amount <= 0 {
		return
	}

	if amount > 50_000_00 {
		return
	}

	card.Balance = card.Balance + amount

}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {

	if !(*card).Active {
		return
	}

	if card.Balance <= 0 {
		return
	}

	bonus := int(card.MinBalance) * percent / 100 * daysInMonth / daysInYear

	if bonus > 5_000_00 {
		// bonus = 5_000_00
		return
	}

	card.Balance = card.Balance + types.Money(bonus)
}

func Total(cards []types.Card) types.Money {

	sum := types.Money(0)

	
	for _, card = range cards {
	if card.Balance >= 0 && card.Active {
			sum += card.Balance
		}
	}

	return sum
}
