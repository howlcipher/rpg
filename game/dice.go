package game

import (
	"math/rand"
	"time"
)

//dice attributes
type Dice struct {
	Sides      int
	SideValues []int
}

//Function that creates new Dice
func NewDice(sides int) *Dice {
	return &Dice{Sides: sides,
		SideValues: SideNumbers(sides),
	}
}


//function that assigns dice sides []int{} values
func SideNumbers(S int) []int {
	SN := []int{}
	for i := 1; i <= S; i++ {
		SN = append(SN, i)
	}
	return SN
}

//roll method.  Takes type Dice and returns type int
func (S Dice) Roll() (int) {
	rand.Seed(time.Now().UTC().UnixNano())   //too make rand work
	roll := 1 + rand.Intn(len(S.SideValues)) //random number between 1 and highest SideValues from Dice struct
	return roll
}

//dice mechanics interface
type Actions interface {
	Roll() int
}



//create the dice slice of numbers
//type CreateDice interface {
//	NewDice() Dice
//}

//type DiceValues interface {
//	Sidenumbers() int
//}
