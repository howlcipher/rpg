package game

import "fmt"

func characterMenu() Character {

	fmt.Println("What type of character would you like to create?" + "\n1: Mage\n2: Fighter\n3: Return to Main Menu")
	var cMenuInput string
	var cname string

	fmt.Scanln(&cMenuInput)
	fmt.Println(cMenuInput)
	switch cMenuInput {
	case "1":
		fmt.Println("Creating Mage")
		fmt.Scanln(&cname)
		fmt.Println(cname)
		CreateMage(Mage{}, cname, RollforStats(20))
	case "2":
		fmt.Println("Creating Fighter")
		CreateFighter(Fighter{}, cname, RollforStats(20))
	case "3":
		fmt.Println("Returning to Main Menu")
		Menu()
	default:
		characterMenu()

	}

}

//func CreateCharacter(s struct{}) struct{} {
//	return s
//}
