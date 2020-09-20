package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

//Withdraw_positive calculates balance with successful withdraw
func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 1000000
}

//Withdraw_noMoney checks the case with no or not enough money
func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 0, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 0
}

//Withdraw_inactive checks the case with disabled card
func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 100)
	fmt.Println(result.Balance)
	// Output: 2000000
}

//Withdraw_limit checks the case with amount being more than limit
func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 1_000_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000
}

//Deposit_positive calculates the successsful deposit
func ExampleDeposit_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 10_000_00)

	fmt.Println(card.Balance)
	// Output: 3000000
}

// Deposit_inactive checks if the card is not active
func ExampleDeposit_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Deposit(&card, 10_000_00)

	fmt.Println(card.Balance)
	// Output: 2000000
}

// Deposit_limit checks if the amount exceeds the limit
func ExampleDeposit_limit() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 60_000_00)

	fmt.Println(card.Balance)
	// Output: 2000000
}

// AddBonus_positive adds bonus to the amount
func ExampleAddBonus_positive() {
	card := types.Card{Balance: 20_000_00, MinBalance: 10_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: 2002465
}

// AddBonus_inactive checks wheather the card is inactive
func ExampleAddBonus_inactive() {
	card := types.Card{Balance: 20_000_00, MinBalance: 10_000_00, Active: false}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: 2000000
}

// ExampleAddBonus_negativeBalance checks if the balance is negative
func ExampleAddBonus_negativeBalance() {
	card := types.Card{Balance: -20_000_00, MinBalance: 10_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: -2000000
}

// AddBonus_aboveLimit adds bonus to the amount
func ExampleAddBonus_aboveLimit() {
	card := types.Card{Balance: 20_000_000_00, MinBalance: 30_000_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: 2000000000
}

// ExampleTotal_positive performs if all conditions are met
func ExampleTotal_positive() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
	}

	fmt.Println(Total(cards))
	// Output: 2000000
}

// ExampleTotal_negative checks the case with negative balance
func ExampleTotal_negative() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: -10_000_00,
			Active:  true,
		},
	}

	fmt.Println(Total(cards))
	// Output: 1000000
}

// ExampleTotal_inactive chacks case when one of the cards is inactive
func ExampleTotal_inactive() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  false,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
	}

	fmt.Println(Total(cards))
	// Output: 1000000
}
