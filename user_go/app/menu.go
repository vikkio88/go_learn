package app

import (
	"fmt"
	"user_store/console"
	"user_store/db"
	"user_store/models"
)

func balance(u *models.User) {
	fmt.Println("Current Balance")
	fmt.Println(fmt.Sprintf("%v", u.GetDefaultAccount().Balance))
	console.EtC()
}

func withdraw(u *models.User) {
	fmt.Println("Withdraw")
	account := u.GetDefaultAccount()
	val := console.GetF64("how much?")
	amount := models.NewMoneyFromF(account.Balance.Currency, val)
	account.Balance.Sub(amount)
	fmt.Println("Done!")
	console.EtC()
}

func deposit(u *models.User) {
	fmt.Println("Deposit")
	val := console.GetF64("how much?")
	account := u.GetDefaultAccount()
	amount := models.NewMoneyFromF(account.Balance.Currency, val)
	account.Balance.Add(amount)
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

func moveMoney(u *models.User, db *db.Db) {
	tries := 3
	fmt.Println("Move money")
	var u2 *models.User
	for {
		if tries < 1 {
			fmt.Println("No tries left, try again later")
			console.EtC()
			return
		}
		id := console.GetStr("Insert account Id")
		ux, err := db.GetUserById(id)
		if err != nil {
			tries--
			fmt.Println(fmt.Sprintf("No account found with that id... try again (%d/3 tries left)", tries))
			continue
		} else {
			u2 = ux
			break
		}

	}
	val := console.GetF64("how much money?")
	//TODO:  check if enough money in the account
	account := u.GetDefaultAccount()
	amount := models.NewMoneyFromF(account.Balance.Currency, val)
	res, err := db.MoveMoney(u.Id, u2.Id, amount)
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
