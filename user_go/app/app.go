package app

import (
	"fmt"
	"user_store/console"
	"user_store/db"
	"user_store/h"
	"user_store/models"
)

type Context struct {
	user    *models.User
	account *models.Account
}

func (c *Context) login(user *models.User) {
	c.user = user
	if !user.IsAdmin() {
		c.account = user.GetDefaultAccount()
	}
}

func (c *Context) logout() {
	c.user = nil
	c.account = nil
}

type App struct {
	db      *db.Db
	state   State
	context *Context
}

func NewApp(db *db.Db) App {
	return App{state: Login, context: &Context{user: nil}, db: db}
}

func (a *App) getContext() (*Context, error) {
	if a.context.user != nil {
		return a.context, nil
	}

	return nil, fmt.Errorf("No User")
}

func (a *App) Run() {
mainLoop:
	for {
		switch a.state {
		case Login:
			a.state = a.login()
		case Dashboard:
			{
				a.state = a.dashboard()
			}
		case Quit:
			break mainLoop
		}
	}

	a.Cleanup()
}

func (a *App) Cleanup() {
	fmt.Println("running cleanup")
	a.db.Persist()
	fmt.Print("bye\n\n\n")
}

func (a *App) login() State {
	console.Cls()
	fmt.Print("\n\nLOGIN\n\n")
	username := console.GetStr("username")
	password := console.GetStr("password")
	u, err := a.db.GetUserByLogin(username, password)
	if err != nil {
		fmt.Print("\n\nInvalid Username/Password!\n")
		console.EtC()
		return Login
	}

	a.context.login(u)
	return Dashboard
}

func (a *App) dashboard() State {
	c, err := a.getContext()
	if err != nil {
		return Login
	}

	if c.user.IsAdmin() {
		return a.adminDashboard(c.user)
	}

	return a.userDashboard(c)
}

func (a *App) adminDashboard(u *models.User) State {
	console.Cls()
	menu := []string{"Reset User Password", "Logout", "Quit"}
	fmt.Println("ADMIN DASHBOARD ", u.Username)
	c := console.ChooseFrom("Menu", menu)
	switch c {
	case 0:
		resetUserPassword(a.db)
	case 1:
		{
			fmt.Println("\nLogging out...")
			a.context.logout()
			return Login
		}
	case 2:
		{
			fmt.Println("Quit")
			console.EtC()
			return Quit
		}
	}

	return Dashboard
}

func (a *App) userDashboard(ctx *Context) State {
	console.Cls()
	u := ctx.user
	acc := ctx.account

	menu := []string{"Balance", "Withdraw", "Deposit", "Move Money", "Change Account", "Change Password", "Logout", "Quit"}
	fmt.Println("Welcome ", u.FullName)
	fmt.Println(h.F("Account: %s (%v)", acc.Name, acc.Balance.Currency))
	c := console.ChooseFrom("Menu", menu)
	switch c {
	case 0:
		balance(acc)
	case 1:
		withdraw(acc)
	case 2:
		deposit(acc)
	case 3:
		moveMoney(u, acc, a.db)
	case 4:
		changeAccount(ctx, a.db)
	case 5:
		changePassword(u)
	case 6:
		{
			fmt.Println("\nLogging out...")
			a.context.logout()
			return Login
		}
	case 7:
		{
			fmt.Println("Quit")
			console.EtC()
			return Quit
		}
	}

	return Dashboard
}

type State = uint8

const (
	Login State = iota
	Dashboard
	Quit
)
