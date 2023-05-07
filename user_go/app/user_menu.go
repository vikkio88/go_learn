package app

import (
	"fmt"
	"user_store/console"
	"user_store/db"
	"user_store/h"
	"user_store/models"
)

func balance(a *models.Account) {
	fmt.Println(h.F("\nAccount: %s", a.Name))
	fmt.Println(h.F("id: %s", a.Id))
	fmt.Println(h.F("currency: %s", a.Balance.Currency))
	fmt.Println("\n\nCurrent Balance")
	fmt.Println(h.F("%v\n", a.Balance))
	console.EtC()
}

func withdraw(a *models.Account) {
	fmt.Println("Withdraw")
	val := console.GetF64("how much?")
	amount := models.NewMoneyFromF(a.Balance.Currency, val)
	a.Balance.Sub(amount)
	fmt.Println("Done!")
	console.EtC()
}

func deposit(a *models.Account) {
	fmt.Println("Deposit")
	val := console.GetF64("how much?")
	amount := models.NewMoneyFromF(a.Balance.Currency, val)
	a.Balance.Add(amount)
	fmt.Println("Done!")
	console.EtC()
}

func changeAccount(ctx *Context, db *db.Db) {
	accNumber := len(ctx.user.Accounts)
	if accNumber < 2 {
		fmt.Println("You only have 1 account")
		console.EtC()
		return
	}
	accounts := make([]string, len(ctx.user.Accounts))
	for i, a := range ctx.user.Accounts {
		accounts[i] = a.Name
	}
	c := console.ChooseFrom("which account", accounts)

	ctx.account = &ctx.user.Accounts[c]
}

func changePassword(u *models.User) {
	fmt.Println("Change Password")
	newP := console.GetStr("new password")
	u.ChangePassword(newP)
	fmt.Println("Done!")
	console.EtC()
}

func moveMoney(u *models.User, account *models.Account, db *db.Db) {
	fmt.Println("Move money")
	payeeId := console.GetStr("Insert account Id")
	val := console.GetF64("how much money?")
	//TODO:  check if enough money in the account
	amount := models.NewMoneyFromF(account.Balance.Currency, val)
	res, err := db.MoveMoney(u.Id, payeeId, amount)
	if res {
		fmt.Println("Done!")
		console.EtC()
		return
	}

	if err != nil {
		fmt.Println(err.Error())
		console.EtC()
	}
}
