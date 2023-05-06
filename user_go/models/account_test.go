package models_test

import (
	"testing"
	"user_store/models"

	"github.com/stretchr/testify/assert"
)

func TestAccountGeneration(t *testing.T) {
	a := models.NewDefaultAccount(models.NewMoney(models.Dollar, 10), "someId")
	assert.Equal(t, models.DefaultAccountName, a.Name)
	assert.Equal(t, "10.00 $", a.Balance.String())
	a1 := models.NewDefaultAccount(models.NewMoney(models.Dollar, 100), "someId")
	assert.Equal(t, models.DefaultAccountName, a1.Name)
	assert.NotEqual(t, a.Id, a1.Id)
}
