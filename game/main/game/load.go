package game

import (
	"encoding/json"
	"io/ioutil"
)

func TurnStructInt(q string) []int {
	c := Character{}
	json.Unmarshal([]byte(q), &c)
	return []int{c.Strength, c.Constitution, c.Dexterity, c.Charisma, c.Wisdom, c.Intelligence}
}

func TurnStructStr(q string) []string {
	c := Character{}
	json.Unmarshal([]byte(q), &c)
	return []string{c.Name}
}

func LoadCharacter(d string, s Character) string {
	lc, err := ioutil.ReadFile(d + ".txt")
	if err != nil {
		return "error!"
	}
	c := string(lc)
	return c
}