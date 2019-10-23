package game

import (
	"fmt"
)

func Menu() string {
	//create menu
	//loop through till selection is made
	//after options run bring menu back
	//

	fmt.Println("Welcome to the rpg game\n" +
		"1: new character\n" +
		"2: load character")
	var menuInput string

	//force to string
	fmt.Scanln(&menuInput)
	fmt.Println(menuInput)
	switch menuInput {
	case "1":
		fmt.Println("Creating New Character")
	case "2":
		fmt.Println("LoadingCharacter")
	default:
		Menu()
	}
	return menuInput
}

func CreateCharacter(s struct{}) struct{} {
	return s
}
