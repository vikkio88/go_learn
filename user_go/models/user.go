package models

import (
	"fmt"
	"strings"

	"github.com/oklog/ulid/v2"
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
	Balance  Money
	password string
	role     Role
}

func NewUser(fullName string, balance Money) User {
	return User{
		Id:       ulid.Make().String(),
		Username: strings.ToLower(strings.ReplaceAll(strings.TrimSpace(fullName), " ", ".")),
		FullName: fullName,
		Balance:  balance,
		password: "qwerty",
		role:     Client,
	}

}

func (u *User) Str() string {
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

func (u *User) Deposit(amount Money) error {
	return u.Balance.Add(amount)
}

func (u *User) Withdraw(amount Money) error {
	return u.Balance.Sub(amount)
}
