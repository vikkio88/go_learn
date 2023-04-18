package app

import (
	"fmt"
	"user_store/console"
)

type App struct {
	State State
	last  string
}

func NewApp() App {
	return App{State: Login}
}

func (a *App) Run() {
	for {
		c := console.GetStr("Gimme")

		if c == "q" {
			break
		}

		a.last = c
		fmt.Println("You chose: ", a.last)
		fmt.Println()
	}

	a.Cleanup()
}

func (a *App) Cleanup() {
	fmt.Println("\n\nCTRL+C Intercepted\n\nrunning cleanup")
	fmt.Println("last choice was: ", a.last)
	fmt.Print("bye\n\n\n")
}

type State = uint8

const (
	Login State = iota
	Dashboard
)
