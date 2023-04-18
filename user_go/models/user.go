package models

import (
	"fmt"
	"strings"

	"github.com/oklog/ulid/v2"
)

type User struct {
	Id       string
	Username string
	FullName string
	Balance  Money
	password string
}

func NewUser(fullName string, balance Money) User {
	return User{
		Id:       ulid.Make().String(),
		Username: strings.ToLower(strings.ReplaceAll(strings.TrimSpace(fullName), " ", ".")),
		FullName: fullName,
		Balance:  balance,
		password: "qwerty",
	}

}

func (u *User) Str() string {
	return fmt.Sprintf("%s %s", u.Id, u.Username)
}

func (u *User) Check(username string, password string) bool {
	return username == u.Username && password == u.password
}

func (u *User) Deposit(amount Money) error {
	return u.Balance.Add(amount)
}

func (u *User) Withdraw(amount Money) error {
	return u.Balance.Sub(amount)
}
