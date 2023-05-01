package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"user_store/interfaces"
	"user_store/libs"
	"user_store/models"

	"github.com/joho/godotenv"
	"golang.org/x/exp/slices"
)

const dbFilePath = "DB_JSON_FILEPATH"
const dbCryptoKey = "DB_CRYPTO_KEY"

func generateUsers() []models.User {
	users := make([]models.User, 0)
	users = append(users, models.NewUser("Mario Rossi", models.NewMoney(models.Euro, 350)))
	users = append(users, models.NewUser("Gianni Bianchi", models.NewMoney(models.Euro, 345_223)))
	users = append(users, models.NewAdmin("admin1"))

	return users
}

type JsonDbConfig struct {
	DbFilePath  string
	DbCryptoKey string
}

func LoadConfigFromEnv() JsonDbConfig {
	err := godotenv.Load()
	if err != nil {
		panic("Could not load .env config")
	}
	return JsonDbConfig{
		DbFilePath:  os.Getenv(dbFilePath),
		DbCryptoKey: os.Getenv(dbCryptoKey),
	}
}

// TODO: Make this an interface so you cna replace it with another Db type
type Db struct {
	config JsonDbConfig
	users  []models.User
	crypto interfaces.CryptoLib
}

func NewDb(config JsonDbConfig) *Db {
	db := Db{
		config: config,
		crypto: libs.NewCrypto(config.DbCryptoKey),
	}

	db.Load()

	return &db
}

func (d *Db) GetUserById(id string) (*models.User, error) {
	idx := slices.IndexFunc(d.users, func(u models.User) bool { return u.Id == id })

	if idx == -1 {
		return nil, fmt.Errorf("No User")
	}

	return &d.users[idx], nil
}

func (d *Db) GetUserByLogin(username string, password string) (*models.User, error) {
	idx := slices.IndexFunc(d.users, func(u models.User) bool { return u.Check(username, password) })

	if idx == -1 {
		return nil, fmt.Errorf("No User")
	}

	return &d.users[idx], nil
}

func (d *Db) MoveMoney(payerId string, payeeId string, amount models.Money) (bool, error) {
	payer, err := d.GetUserById(payerId)
	payee, err2 := d.GetUserById(payeeId)
	if err != nil || err2 != nil {
		return false, NewErrorUserNotFound()
	}

	if payer.Balance.Cmp(amount) < 0 {
		return false, models.NewErrorInsufficientFunds()
	}

	if !payer.Balance.SameCurrency(*payee.Balance) {
		return false, models.NewErrorDifferentCurrency(payer.Balance.Currency, payee.Balance.Currency)
	}

	payer.Balance.Sub(amount)
	payee.Balance.Add(amount)

	return true, nil
}

func (d *Db) Persist() {
	users := len(d.users)
	dtos := make([]models.UserDTO, users)
	for i, u := range d.users {
		dtos[i] = u.DTO(d.crypto)
	}
	data, _ := json.Marshal(dtos)
	os.WriteFile(d.config.DbFilePath, data, 0644)
}

func (d *Db) Load() {
	filepath := d.config.DbFilePath

	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		d.users = generateUsers()
		d.Persist()
		return
	}

	data, err := os.ReadFile(filepath)
	if err != nil {
		d.users = generateUsers()
		d.Persist()
		return
	}

	var userDtos []models.UserDTO
	json.Unmarshal(data, &userDtos)

	users := make([]models.User, len(userDtos))
	for i, dto := range userDtos {
		users[i] = dto.User(d.crypto)
	}

	d.users = users
}

type ErrorUserNotFound struct{}

func NewErrorUserNotFound() ErrorUserNotFound {
	return ErrorUserNotFound{}
}
func (e ErrorUserNotFound) Error() string {
	return "User not found"
}
