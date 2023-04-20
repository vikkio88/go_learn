package models_test

import (
	"testing"
	"user_store/models"

	"github.com/stretchr/testify/assert"
)

func TestUserContruction(t *testing.T) {
	u := models.NewUser("Mario Marii", models.NewMoney(models.Euro, 0))
	assert.Equal(t, "mario.marii", u.Username)
	assert.False(t, u.IsAdmin())

	u2 := models.NewUser("Giacomo Cosino", models.NewMoney(models.Dollar, 0))
	assert.NotEqual(t, u.Id, u2.Id)

	assert.False(t, u2.IsAdmin())
}

func TestCheckingLogin(t *testing.T) {
	u := models.NewUser("Fabrizio Fumagalli", models.NewMoney(models.Euro, 0))

	assert.False(t, u.Check("fabro.fuma", "qwerty"))
	assert.True(t, u.Check("fabrizio.fumagalli", "qwerty"))
}

func TestPasswordUpdate(t *testing.T) {
	u := models.NewUser("Mario Rossi", models.Money{})
	assert.True(t, u.Check("mario.rossi", "qwerty"))

	u.ChangePassword("qwertyuiop")
	assert.True(t, u.Check("mario.rossi", "qwertyuiop"))
}

func TestBalanceOperations(t *testing.T) {
	u := models.NewUser("Mario Marii", models.NewMoney(models.Dollar, 0))

	assert.Equal(t, "0.00 $", u.Balance.Str())
	err := u.Deposit(models.NewMoney(models.Dollar, 300))
	assert.Nil(t, err)
	assert.Equal(t, "300.00 $", u.Balance.Str())
	err = u.Withdraw(models.NewMoney(models.Dollar, 150))
	assert.Nil(t, err)
	assert.Equal(t, "150.00 $", u.Balance.Str())

	err = u.Deposit(models.NewMoney(models.Euro, 300))
	assert.NotNil(t, err)
	assert.Equal(t, "150.00 $", u.Balance.Str())

	err = u.Withdraw(models.NewMoney(models.Euro, 300))
	assert.NotNil(t, err)
	assert.Equal(t, "150.00 $", u.Balance.Str())
}

func TestUserToJson(t *testing.T) {
	u := models.NewUser("Mario Bros", models.NewMoney(models.Dollar, 100))
	res, err := u.ToJson()
	assert.Nil(t, err)
	assert.Contains(t, res, "\"username\":\"mario.bros\",\"fullName\":\"Mario Bros\",\"balance\":{\"val\":10000,\"currency\":0},\"password\":\"qwerty\",\"role\":0}")
}
