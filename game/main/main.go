package main

import (
	"./game"
	"fmt"
)

func main() {

	//calls the menu
	game.Menu()
	//creates a d20 dice
	d20 := game.NewDice(20)
	//n := game.CreateMage()

	//creates a new characters attributes
	nc := game.Attributes(game.Actions.Roll(d20))
	game.SaveStats(*nc, "sBrody")

	//load character
	//  Loads JSON data from text file
	loadCharacter := game.LoadCharacter("sBrody", game.Character{})
	//turns load data into slice of string or int
	lcs := game.TurnStructStr(loadCharacter)
	lci := game.TurnStructInt(loadCharacter)
	//creates a loaded character
	lc := game.Character{Name: lcs[0], Strength: lci[0], Constitution: lci[1], Dexterity: lci[2], Charisma: lci[3], Wisdom: lci[4], Intelligence: lci[5]}

	//displays character attributes
	dBrody := game.Display.DisplayCharacter(lc)
	fmt.Println(dBrody)
	//Character.SaveCharacter(dBrody,"Brody")
	//Character.SaveStats(*Brody, "sBrody")

	//Rolls dice
	fmt.Printf("\nYou rolled a %v sided dice and rolled a: ", d20.Sides)
	fmt.Println(game.Actions.Roll(d20))
}
