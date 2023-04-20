package models

import (
	"fmt"
	"math"
)

type Money struct {
	Val      int32    `json:"val"`
	Currency Currency `json:"currency"`
}

const MULTIPLIER_100 int32 = 100
const MULTIPLIERF_100 float32 = 100
const MULTIPLIERF64_100 float64 = 100

func NewMoney(currency Currency, unit int32) Money {
	return Money{
		Val:      unit * MULTIPLIER_100,
		Currency: currency,
	}
}

func NewMoneyUF(currency Currency, unit int32, fractional int32) Money {
	return Money{
		Val:      unit*MULTIPLIER_100 + fractional,
		Currency: currency,
	}
}

func NewMoneyFromF(currency Currency, amount float64) Money {
	val, _ := math.Modf(amount * 100)

	return Money{Val: int32(val), Currency: currency}
}

func (m *Money) GetCurrency() Currency {
	return m.Currency
}

func (m *Money) Str() string {
	return fmt.Sprintf("%.2f %s", float32(m.Val)/MULTIPLIERF_100, m.Currency.Str())
}

func (m *Money) Add(n Money) error {
	if m.Currency != n.Currency {
		return fmt.Errorf("Currencies")
	}

	m.Val += n.Val
	return nil
}

func (m *Money) Sub(n Money) error {
	if m.Currency != n.Currency {
		return fmt.Errorf("Currencies")
	}

	m.Val -= n.Val
	return nil
}

type Currency uint8

const (
	Dollar Currency = iota
	Euro
	Pound
)

func (c Currency) Str() string {
	switch c {
	case Dollar:
		return "$"
	case Euro:
		return "€"
	case Pound:
		return "£"
	}

	return ""
}
