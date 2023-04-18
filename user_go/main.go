package main

import (
	"os"
	"os/signal"
	"syscall"
	"user_store/app"
)

func main() {
	// name := console.GetStr("name")
	// surname := console.GetStr("surname")
	// balance := console.GetInt("balance")
	// u := models.NewUser(fmt.Sprintf("%s %s", name, surname), models.NewMoney(models.Dollar, balance))
	// fmt.Println(u.Str())

	// c := console.ChooseFrom("Groceries", []string{"Milk", "Onion", "Potatoes"})
	// fmt.Println("choice: ", c)
	// m := make(map[string]string)
	// m["a"] = "Milk"
	// m["b"] = "Onions"
	// m["q"] = "Quit"
	// c1 := console.ChooseFromMap("Groceries", m)
	// if val, ok := m[c1]; ok {
	// 	fmt.Println("You choose ", val)
	// }

	a := app.NewApp()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		a.Cleanup()
		os.Exit(0)
	}()

	a.Run()
}
