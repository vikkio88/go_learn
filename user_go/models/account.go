package models

const DefaultAccountName = "DEFAULT"

type Account struct {
	Id      string
	Name    string
	Balance *Money
}

func NewDefaultAccount(amount Money) Account {
	return Account{
		Id:      "",
		Name:    DefaultAccountName,
		Balance: &amount,
	}
}
