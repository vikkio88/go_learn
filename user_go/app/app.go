package app

import (
	"fmt"
	"user_store/console"
)

type App struct {
	state State
}

func NewApp() App {
	return App{state: Login}
}

func (a *App) Run() {
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
				break
			}
		}

	}

	a.Cleanup()
}

func (a *App) Cleanup() {
	fmt.Println("\n\nCTRL+C Intercepted\n\nrunning cleanup")
	fmt.Print("bye\n\n\n")
}

func (a *App) login() State {
	println("\n\nLOGIN\n\n")
	username := console.GetStr("username")
	password := console.GetStr("password")
	println(username, password)

	return Dashboard
}

func (a *App) dashboard() State {
	menu := make(map[string]string, 4)
	menu["1"] = "Balance"
	menu["2"] = "Withdraw"
	menu["3"] = "Deposit"
	menu["4"] = "Logout"
	menu["5"] = "Logout"
	c := console.ChooseFromMap("Menu", menu)
	println(c)

	return Login
}

type State = uint8

const (
	Login State = iota
	Dashboard
	Quit
)
