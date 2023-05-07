package models_test

import (
	"testing"
	"user_store/models"

	"github.com/stretchr/testify/assert"
)

func TestMoneyCreation(t *testing.T) {
	m := models.NewMoney(models.Dollar, 20)
	assert.Equal(t, "20.00 $", m.String())
	m = models.NewMoney(models.Euro, 20)
	assert.Equal(t, "20.00 €", m.String())
	m = models.NewMoney(models.Pound, 20)
	assert.Equal(t, "20.00 £", m.String())
}
func TestMoneyMath(t *testing.T) {
	m := models.NewMoney(models.Dollar, 300)
	n := models.NewMoney(models.Dollar, 150)

	err := m.Add(n)
	assert.Nil(t, err)
	assert.Equal(t, "450.00 $", m.String())

	err = m.Sub(n)
	assert.Nil(t, err)
	assert.Equal(t, "300.00 $", m.String())

	err = m.Sub(models.NewMoney(models.Euro, 140))
	assert.NotNil(t, err)
	assert.Errorf(t, err, "Currency")
	assert.Equal(t, "300.00 $", m.String())
}

func TestMoneyWithFractional(t *testing.T) {
	m := models.NewMoneyUF(models.Dollar, 20, 75)
	assert.Equal(t, "20.75 $", m.String())
	m = models.NewMoneyUF(models.Dollar, 20, 173)
	assert.Equal(t, "21.73 $", m.String())
}

func TestMoneyFromFloat(t *testing.T) {
	m := models.NewMoneyFromF(models.Dollar, 21.54)
	assert.Equal(t, "21.54 $", m.String())
	m = models.NewMoneyFromF(models.Dollar, 21.54)
	assert.Equal(t, "21.54 $", m.String())
}

func TestCurrencyFromString(t *testing.T) {
	assert.Equal(t, models.Pound, models.CurrencyFromString("£"))
	assert.Equal(t, models.Dollar, models.CurrencyFromString("$"))
	assert.Equal(t, models.Euro, models.CurrencyFromString("€"))

}
