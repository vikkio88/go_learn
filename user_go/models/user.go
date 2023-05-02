package models

import (
	"fmt"
	"strings"
	"user_store/interfaces"

	"github.com/oklog/ulid/v2"
	"golang.org/x/exp/slices"
)

type Role uint8

const (
	Client = iota
	Admin
)

type User struct {
	Id       string
	Username string
	FullName string
	Accounts []Account
	password string
	role     Role
}

func NewUser(fullName string, balance Money) User {
	accounts := []Account{NewDefaultAccount(balance)}
	return User{
		Id:       ulid.Make().String(),
		Username: strings.ToLower(strings.ReplaceAll(strings.TrimSpace(fullName), " ", ".")),
		FullName: fullName,
		Accounts: accounts,
		password: "qwerty",
		role:     Client,
	}
}

func NewAdmin(username string) User {
	return User{
		Id:       ulid.Make().String(),
		Username: username,
		Accounts: nil,
		password: "s4f3p4ssw0rd!",
		role:     Admin,
	}
}

func (u *User) String() string {
	return fmt.Sprintf("%s %s", u.Id, u.Username)
}

func (u *User) ChangePassword(newPassword string) {
	u.password = newPassword
}

func (u *User) Check(username string, password string) bool {
	return username == u.Username && password == u.password
}

func (u *User) IsAdmin() bool {
	return u.role == Admin
}

func (u *User) GetAccountById(id string) Account {
	idx := slices.IndexFunc(u.Accounts, func(a Account) bool { return a.Id == id })

	return u.Accounts[idx]
}

func (u *User) GetDefaultAccount() *Account {
	idx := slices.IndexFunc(u.Accounts, func(a Account) bool { return a.Name == DefaultAccountName })
	return &u.Accounts[idx]
}

// func (u *User) Deposit(amount Money, accountId string) error {
func (u *User) Deposit(amount Money) error {
	account := u.GetDefaultAccount()
	return account.Balance.Add(amount)
}

// func (u *User) Withdraw(amount Money, accountId string) error {
func (u *User) Withdraw(amount Money) error {
	account := u.GetDefaultAccount()
	return account.Balance.Sub(amount)
}

func (u *User) DTO(cryptoHelper interfaces.CryptoLib) UserDTO {
	hiddenPassword, _ := cryptoHelper.Encrypt(u.password)
	based := cryptoHelper.B64Encode(hiddenPassword)
	return UserDTO{
		Id:       u.Id,
		Username: u.Username,
		FullName: u.FullName,
		Accounts: u.Accounts,
		Password: based,
		Role:     u.role,
	}
}

type UserDTO struct {
	Id       string    `json:"id"`
	Username string    `json:"username"`
	FullName string    `json:"fullName"`
	Accounts []Account `json:"accounts"`
	Password string    `json:"password"`
	Role     Role      `json:"role"`
}

func (u *UserDTO) User(cryptoHelper interfaces.CryptoLib) User {
	coded, _ := cryptoHelper.B64Decode(u.Password)
	password, _ := cryptoHelper.Decrypt(string(coded))
	//TODO: add error handling here
	return User{
		Id:       u.Id,
		Username: u.Username,
		FullName: u.FullName,
		Accounts: u.Accounts,
		password: password,
		role:     u.Role,
	}
}
