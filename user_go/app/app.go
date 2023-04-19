package app

import (
	"fmt"
	"user_store/console"
	"user_store/db"
	"user_store/models"
)

type Context struct {
	user *models.User
}

func (c *Context) login(user *models.User) {
	c.user = user
}

func (c *Context) logout() {
	c.user = nil
}

type App struct {
	db      *db.Db
	state   State
	context *Context
}

func NewApp(db *db.Db) App {
	return App{state: Login, context: &Context{user: nil}, db: db}
}

func (a *App) getUser() (*models.User, error) {
	if a.context.user != nil {
		return a.context.user, nil
	}

	return nil, fmt.Errorf("No User")
}

func (a *App) Run() {
mainLoop:
	for {
		switch a.state {
		case Login:
			{
				a.state = a.login()
			}

		case Dashboard:
			{
				a.state = a.dashboard()
			}

		case Quit:
			{
				break mainLoop
			}
		}
	}

	a.Cleanup()
}

func (a *App) Cleanup() {
	fmt.Println("running cleanup")
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
	u, err := a.getUser()
	if err != nil {
		return Login
	}
	console.Cls()
	menu := []string{"Balance", "Withdraw", "Deposit", "Change Password", "Logout", "Quit"}
	fmt.Println("Welcome ", u.FullName)
	c := console.ChooseFrom("Menu", menu)
	switch c {
	case 0:
		{
			fmt.Println("Current Balance")
			fmt.Println(u.Balance.Str())
			console.EtC()
		}
	case 1:
		{
			fmt.Println("Withdraw")
			val := console.GetInt("how much?")
			amount := models.NewMoney(u.Balance.Currency(), val)
			u.Balance.Sub(amount)
			fmt.Println("Done!")
			console.EtC()
		}
	case 2:
		{
			fmt.Println("Deposit")
			val := console.GetInt("how much?")
			amount := models.NewMoney(u.Balance.Currency(), val)
			u.Balance.Add(amount)
			fmt.Println("Done!")
			console.EtC()
		}
	case 3:
		{
			fmt.Println("Change Password")
			newP := console.GetStr("new password")
			u.ChangePassword(newP)
			fmt.Println("Done!")
			console.EtC()
		}
	case 4:
		{
			fmt.Println("Signing you out...")
			a.context.logout()
			console.EtC()
			return Login
		}
	case 5:
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
