package db

import (
	"fmt"
	"user_store/models"

	"golang.org/x/exp/slices"
)

type Db struct {
	users []models.User
}

func NewDb() *Db {
	users := make([]models.User, 3)
	users = append(users, models.NewUser("Mario Rossi", models.NewMoney(models.Euro, 350)))
	users = append(users, models.NewUser("Gianni Bianchi", models.NewMoney(models.Euro, 345_223)))
	return &Db{users: users}
}

func (d *Db) GetUserByLogin(username string, password string) (*models.User, error) {
	idx := slices.IndexFunc(d.users, func(u models.User) bool { return u.Check(username, password) })

	if idx == -1 {
		return nil, fmt.Errorf("No User")
	}

	return &d.users[idx], nil
}
