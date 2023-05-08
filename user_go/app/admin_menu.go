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

func deleteUser(db *db.Db) {
	fmt.Println("Delete User")
	search := console.GetStr("search by full name or username")
	users := db.GetUsers(search)
	if len(users) == 0 {
		fmt.Println("No user!")
		console.EtC()
		return
	}

	user := users[0]

	if len(users) > 1 {
		userList := make([]string, len(users))
		for i, u := range users {
			userList[i] = u.Username
		}
		idx := console.ChooseFrom("choose the user to delete", userList)
		user = users[idx]
	}

	fmt.Println(h.F("you choose to delete [%s]", user))
	confirm := console.Confirm()
	if confirm {
		db.DeleteUser(user.Id)
		fmt.Println("User Deleted!")
	}

	console.EtC()
}
