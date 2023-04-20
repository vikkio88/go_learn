package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"user_store/models"

	"golang.org/x/exp/slices"
)

const dbFilePath = "./db.json"

func generateUsers() []models.User {
	users := make([]models.User, 0)
	users = append(users, models.NewUser("Mario Rossi", models.NewMoney(models.Euro, 350)))
	users = append(users, models.NewUser("Gianni Bianchi", models.NewMoney(models.Euro, 345_223)))
	users = append(users, models.NewAdmin("admin1"))

	return users
}

type Db struct {
	users []models.User
}

func NewDb() *Db {
	db := Db{}

	db.Load()

	return &db
}

func (d *Db) GetUserByLogin(username string, password string) (*models.User, error) {
	idx := slices.IndexFunc(d.users, func(u models.User) bool { return u.Check(username, password) })

	if idx == -1 {
		return nil, fmt.Errorf("No User")
	}

	return &d.users[idx], nil
}

func (d *Db) Persist() {
	data, _ := json.Marshal(d.users)
	os.WriteFile(dbFilePath, data, 0644)
}

func (d *Db) Load() {
	if _, err := os.Stat(dbFilePath); errors.Is(err, os.ErrNotExist) {
		d.users = generateUsers()
		d.Persist()
		return
	}

	data, err := os.ReadFile(dbFilePath)
	if err != nil {
		d.users = generateUsers()
		d.Persist()
		return
	}

	users := make([]models.User, 2)
	json.Unmarshal(data, &users)
	d.users = users
}
