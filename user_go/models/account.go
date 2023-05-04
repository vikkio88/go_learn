package models

import "github.com/oklog/ulid/v2"

const DefaultAccountName = "DEFAULT"

type Account struct {
	Id      string
	Name    string
	Balance *Money
}

func idGenerator() string {
	return ulid.Make().String()
}

func NewDefaultAccount(amount Money) Account {
	return Account{
		Id:      idGenerator(),
		Name:    DefaultAccountName,
		Balance: &amount,
	}
}
