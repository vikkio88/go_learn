package app

import (
	"fmt"
	"user_store/console"
	"user_store/db"
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
	amount := models.NewMoneyFromF(u.Balance.GetCurrency(), val)
	u.Balance.Sub(amount)
	u2.Balance.Add(amount)
	fmt.Println("Done!")
	console.EtC()
}
