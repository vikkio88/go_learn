package app

import (
	"fmt"
	"user_store/console"
	"user_store/db"
	"user_store/h"
	"user_store/models"
)

func resetUserPassword(db *db.Db) {
	fmt.Println("Admin Panel > Reset User Password")
	var user *models.User
	for {
		id := console.GetStr("UserId")
		if id == "q" {
			return
		}
		u, err := db.GetUserById(id)

		if err != nil {
			fmt.Println("Not a valid user id, try again...\n[q] to quit")
		} else {
			user = u
			break
		}
	}

	fmt.Println("User: ", user)
	newPassword := console.GetStr("insert new password")
	user.ChangePassword(newPassword)
	fmt.Println("Done!")
	console.EtC()
}

func createNewUser(db *db.Db) {
	fmt.Println("User Creation")
	name := console.GetStr("Name:")
	surname := console.GetStr("Surname:")
	currencies := []string{models.Dollar.String(), models.Euro.String(), models.Pound.String()}
	idx := console.ChooseFrom("Default Account Currency", currencies)
	c := models.CurrencyFromString(currencies[idx])
	val := console.GetF64(h.F("Amount of money in the default account (%s)", c))
	amount := models.NewMoneyFromF(c, val)
	user := models.NewUser(h.F("%s %s", name, surname), amount)
	db.AddUser(user)

	fmt.Println(h.F("User created %s", &user))
	console.EtC()
}

func DeleteUser(db *db.Db) {
	return
	fmt.Println("User deleted!")
	console.EtC()
}
