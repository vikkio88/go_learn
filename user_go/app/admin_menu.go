package app

import (
	"fmt"
	"user_store/console"
	"user_store/db"
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
