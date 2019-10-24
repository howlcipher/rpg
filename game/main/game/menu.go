package game

import (
	"fmt"
)

func Menu() {
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
		characterMenu()
	case "2":
		fmt.Println("LoadingCharacter")
		//load character
		//  Loads JSON data from text file
		loadCharacter := LoadCharacter("sBrody", Character{})
		//turns load data into slice of string or int
		lcs := TurnStructStr(loadCharacter)
		lci := TurnStructInt(loadCharacter)
		//creates a loaded character
		lc := Character{Name: lcs[0], Strength: lci[0], Constitution: lci[1], Dexterity: lci[2], Charisma: lci[3], Wisdom: lci[4], Intelligence: lci[5]}
		dBrody := Display.DisplayCharacter(lc)
		fmt.Println(dBrody)
	case "3":
		fmt.Println("Quitting, thanks for playing!")
	default:
		Menu()
	}

}
