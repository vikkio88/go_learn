package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"user_store/app"
	"user_store/db"
)

func setupCtrlC(a *app.App) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go handleCtrlC(c, a)
}

func handleCtrlC(c chan os.Signal, a *app.App) {
	<-c
	fmt.Print("\n\nCTRL+C Intercepted\n\n")
	a.Cleanup()
	os.Exit(0)
}

func main() {
	db := db.NewDb()
	a := app.NewApp(db)
	setupCtrlC(&a)

	a.Run()
}
