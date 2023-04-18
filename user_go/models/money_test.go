package models_test

import (
	"testing"
	"user_store/models"

	"github.com/stretchr/testify/assert"
)

func TestMoneyCreation(t *testing.T) {
	m := models.NewMoney(models.Dollar, 20)
	assert.Equal(t, "20.00 $", m.Str())
	m = models.NewMoney(models.Euro, 20)
	assert.Equal(t, "20.00 €", m.Str())
	m = models.NewMoney(models.Pound, 20)
	assert.Equal(t, "20.00 £", m.Str())
}
func TestMoneyMath(t *testing.T) {
	m := models.NewMoney(models.Dollar, 300)
	n := models.NewMoney(models.Dollar, 150)

	err := m.Add(n)
	assert.Nil(t, err)
	assert.Equal(t, "450.00 $", m.Str())

	err = m.Sub(n)
	assert.Nil(t, err)
	assert.Equal(t, "300.00 $", m.Str())

	err = m.Sub(models.NewMoney(models.Euro, 140))
	assert.NotNil(t, err)
	assert.Errorf(t, err, "Currency")
	assert.Equal(t, "300.00 $", m.Str())
}
