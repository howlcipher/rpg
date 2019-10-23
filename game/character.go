package game

import (
	"math/rand"
	"strconv"
)

//character attributes
type Character struct {
	Name         string `json:"name"`
	Strength     int    `json:"strength"`
	Constitution int 	 `json:"constitution"`
	Dexterity    int    `json:"dexterity"`
	Charisma     int    `json:"charisma"`
	Wisdom       int    `json:"wisdom"`
	Intelligence int    `json:"intelligence"`
}

func Attributes(d20 int) *Character {
	return &Character{
		Name:         "Brody",
		Strength:     1 + rand.Intn(d20),
		Constitution: 1 + rand.Intn(d20),
		Dexterity:    1 + rand.Intn(d20),
		Charisma:     1 + rand.Intn(d20),
		Wisdom:       1 + rand.Intn(d20),
		Intelligence: 1 + rand.Intn(d20),
	}
}

func bubblesort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}

func stmod(s []int) []int {
	for i := 0; i < len(s); i++ {
		if s[i] >= 14 {
			s[i] += 4
		} else if s[i] >= 12 {
			s[i] += 2
		} else if s[i] >= 10 {
			s[i] += 1
		} else if s[i] < 9 {
			s[i] -= 2
		}
	}

	return []int{}
}


//func (c Character) StatsToString(x int)string{
//	return 	strconv.Itoa(x)
//	}

func (c Character) DisplayCharacter() string {

	st := strconv.Itoa(c.Strength)
	cn := strconv.Itoa(c.Constitution)
	dx := strconv.Itoa(c.Dexterity)
	ch := strconv.Itoa(c.Charisma)
	wi := strconv.Itoa(c.Wisdom)
	in := strconv.Itoa(c.Intelligence)

	x := "Your Character is: '" + c.Name + "'" +
		"\n\tStrength:\t\t" + st +
		"\n\tConstitution:\t" + cn +
		"\n\tAgility:\t\t" + dx +
		"\n\tCharisma:\t\t" + ch +
		"\n\tWisdom:\t\t\t" + wi +
		"\n\tIntelligence:\t" + in

	return x

}

type Display interface {
	DisplayCharacter() string
	//StatsToString() string
}
