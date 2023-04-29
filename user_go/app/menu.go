package app

import (
	"fmt"
	"user_store/console"
	"user_store/models"
)

func balance(u *models.User) {
	fmt.Println("Current Balance")
	fmt.Println(u.Balance)
	console.EtC()
}

func withdraw(u *models.User) {
	fmt.Println("Withdraw")
	val := console.GetF64("how much?")
	amount := models.NewMoneyFromF(u.Balance.GetCurrency(), val)
	u.Balance.Sub(amount)
	fmt.Println("Done!")
	console.EtC()
}

func deposit(u *models.User) {
	fmt.Println("Deposit")
	val := console.GetF64("how much?")
	amount := models.NewMoneyFromF(u.Balance.GetCurrency(), val)
	u.Balance.Add(amount)
	fmt.Println("Done!")
	console.EtC()
}

func changePassword(u *models.User) {
	fmt.Println("Change Password")
	newP := console.GetStr("new password")
	u.ChangePassword(newP)
	fmt.Println("Done!")
	console.EtC()
}
