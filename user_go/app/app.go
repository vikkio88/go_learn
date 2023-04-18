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
	fmt.Println("\n\nCTRL+C Intercepted\n\nrunning cleanup")
	fmt.Print("bye\n\n\n")
}

func (a *App) login() State {
	fmt.Print("\n\nLOGIN\n\n")
	username := console.GetStr("username")
	password := console.GetStr("password")
	fmt.Println(username, password)

	return Dashboard
}

func (a *App) dashboard() State {
	menu := []string{"Balance", "Withdraw", "Deposit", "Logout", "Quit"}
	c := console.ChooseFrom("Menu", menu)
	fmt.Println(c)

	return Quit
}

type State = uint8

const (
	Login State = iota
	Dashboard
	Quit
)
