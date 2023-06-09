package console

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"user_store/h"
)

func Cls() {
	fmt.Print("\033[H\033[2J")
}

// Press [Enter] to continue
func EtC() {
	GetStr("[ENTER] to continue...")
}

func GetStr(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt + " : ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func GetInt(prompt string) int32 {
	for {
		str := GetStr(prompt)
		res, err := strconv.ParseInt(str, 10, 0)

		if err != nil {
			fmt.Println("that is not a number, try again")
			continue
		}
		return int32(res)
	}
}

func GetF64(prompt string) float64 {
	for {
		str := GetStr(prompt)
		res, err := strconv.ParseFloat(str, 64)

		if err != nil {
			fmt.Println("that is not a number, try again")
			continue
		}
		return res
	}
}

// returns index of the array list passed as second parameter
func ChooseFrom(prompt string, list []string) uint {
	max := len(list)

	if max < 2 {
		fmt.Println("Can't choose from that list")
		return 0
	}
	fmt.Println(prompt)
	for i, item := range list {
		fmt.Println(h.F("%d . %v", i+1, item))
	}

	for {
		choice := GetInt(h.F("[1-%d]", max))
		if choice < 1 || choice > int32(max) {
			println("Choice not in the menu")
			continue
		}
		return uint(choice) - 1
	}
}

func ChooseFromMap(prompt string, list map[string]string) string {
	max := len(list)

	if max < 2 {
		fmt.Println("Can't choose from that list")
		return ""
	}
	fmt.Println(prompt)
	for key, item := range list {
		fmt.Println(h.F("%s . %s", key, item))
	}

	for {
		choice := GetStr("letter ")
		_, ok := list[choice]
		if !ok {
			println("Choice not in the menu")
			continue
		}
		return choice
	}
}

func Confirm() bool {
	str := GetStr("Are you sure? [y/n]")
	str = strings.ToLower(str)
	return str == "y" || str == "yes"
}
