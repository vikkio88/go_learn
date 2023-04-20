package models

import (
	"encoding/json"
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
	Id       string `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Balance  *Money `json:"balance"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

func NewUser(fullName string, balance Money) User {
	return User{
		Id:       ulid.Make().String(),
		Username: strings.ToLower(strings.ReplaceAll(strings.TrimSpace(fullName), " ", ".")),
		FullName: fullName,
		Balance:  &balance,
		Password: "qwerty",
		Role:     Client,
	}
}

func NewAdmin(username string) User {
	return User{
		Id:       ulid.Make().String(),
		Username: username,
		Balance:  nil,
		Password: "s4f3p4ssw0rd!",
		Role:     Admin,
	}
}

func (u *User) Str() string {
	return fmt.Sprintf("%s %s", u.Id, u.Username)
}

func (u *User) ChangePassword(newPassword string) {
	u.Password = newPassword
}

func (u *User) Check(username string, password string) bool {
	return username == u.Username && password == u.Password
}

func (u *User) IsAdmin() bool {
	return u.Role == Admin
}

func (u *User) Deposit(amount Money) error {
	return u.Balance.Add(amount)
}

func (u *User) Withdraw(amount Money) error {
	return u.Balance.Sub(amount)
}

func (u *User) ToJson() (string, error) {
	obj, err := json.Marshal(u)
	return string(obj), err
}
