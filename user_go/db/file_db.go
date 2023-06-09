package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"user_store/h"
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
	config      JsonDbConfig
	users       []models.User
	accountsMap map[string]*models.Account
	crypto      interfaces.CryptoLib
}

func NewDb(config JsonDbConfig) *Db {
	db := Db{
		config:      config,
		crypto:      libs.NewCrypto(config.DbCryptoKey),
		accountsMap: map[string]*models.Account{},
	}

	db.Load()

	return &db
}

func (d *Db) AddUser(u models.User) {
	d.users = append(d.users, u)
	d.indexUserAccounts(&u)
	d.Persist()
}
func (d *Db) DeleteUser(id string) {
	idx := slices.IndexFunc(d.users, func(u models.User) bool { return u.Id == id })
	if idx < 0 {
		return
	}

	d.users[idx] = d.users[len(d.users)-1]
	d.users = d.users[:len(d.users)-1]
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

func (d *Db) GetUsers(search string) []*models.User {
	result := make([]*models.User, 0)
	search = strings.ToLower(search)
	for _, u := range d.users {
		//TODO: check why rossi returns also things that are not supposed to be there
		if strings.Contains(strings.ToLower(u.FullName), search) || strings.Contains(u.Username, search) {
			result = append(result, &u)
		}
	}

	return result
}

func (d *Db) GetAccountById(payeeAccountId string) (*models.Account, error) {
	if account, exists := d.accountsMap[payeeAccountId]; exists {
		return account, nil
	}

	return nil, errors.New(h.F("Account with id %s does not exist", payeeAccountId))
}

func (d *Db) MoveMoney(payerId string, payeeAccountId string, amount models.Money) (bool, error) {
	payer, err := d.GetUserById(payerId)
	if err != nil {
		return false, NewErrorUserNotFound()
	}
	payeeAccount, err2 := d.GetAccountById(payeeAccountId)
	if err2 != nil {
		return false, NewErrorAccountNotFound()
	}

	payerAccount := payer.GetDefaultAccount()

	if payerAccount.Balance.Cmp(amount) < 0 {
		return false, models.NewErrorInsufficientFunds()
	}

	if !payerAccount.Balance.SameCurrency(*payeeAccount.Balance) {
		return false, models.NewErrorDifferentCurrency(payeeAccount.Balance.Currency, payeeAccount.Balance.Currency)
	}

	payerAccount.Balance.Sub(amount)
	payeeAccount.Balance.Add(amount)

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

func (d *Db) indexUserAccounts(user *models.User) {
	for _, account := range user.Accounts {
		d.accountsMap[account.Id] = &account
	}
}
func (d *Db) generateUsers() {
	d.users = generateUsers()
	for _, user := range d.users {
		d.indexUserAccounts(&user)
	}

	d.Persist()
}

func (d *Db) Load() {
	filepath := d.config.DbFilePath

	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		d.generateUsers()
		return
	}

	data, err := os.ReadFile(filepath)
	if err != nil {
		d.generateUsers()
		return
	}

	var userDtos []models.UserDTO
	json.Unmarshal(data, &userDtos)

	users := make([]models.User, len(userDtos))
	for i, dto := range userDtos {
		users[i] = dto.User(d.crypto)
		d.indexUserAccounts(&users[i])
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

type ErrorAccountNotFound struct{}

func NewErrorAccountNotFound() ErrorUserNotFound {
	return ErrorUserNotFound{}
}
func (e ErrorAccountNotFound) Error() string {
	return "Account not found"
}
