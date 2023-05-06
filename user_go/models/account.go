package models

import "github.com/oklog/ulid/v2"

const DefaultAccountName = "DEFAULT"

type Account struct {
	Id      string
	OwnerId string
	Name    string
	Balance *Money
}

func idGenerator() string {
	return ulid.Make().String()
}

func NewDefaultAccount(amount Money, ownerId string) Account {
	return Account{
		Id:      idGenerator(),
		OwnerId: ownerId,
		Name:    DefaultAccountName,
		Balance: &amount,
	}
}
