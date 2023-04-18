package models

import "fmt"

type Money struct {
	val      int32
	currency Currency
}

const MULTIPLIER_100 int32 = 100
const MULTIPLIERF_100 float32 = 100

func NewMoney(currency Currency, unit int32) Money {
	return Money{
		val:      unit * MULTIPLIER_100,
		currency: currency,
	}
}

func (m *Money) Str() string {
	return fmt.Sprintf("%.2f %s", float32(m.val)/MULTIPLIERF_100, m.currency.Str())
}

func (m *Money) Add(n Money) error {
	if m.currency != n.currency {
		return fmt.Errorf("Currencies")
	}

	m.val += n.val
	return nil
}

func (m *Money) Sub(n Money) error {
	if m.currency != n.currency {
		return fmt.Errorf("Currencies")
	}

	m.val -= n.val
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
