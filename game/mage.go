package game

type Mage struct {
	Character
}

func CreateMage(m Mage, s string, n []int) Mage {
	return Mage{Character{Name: s,
		Strength:     n[0],
		Constitution: n[3],
		Dexterity:    n[1],
		Charisma:     n[2],
		Wisdom:       n[4],
		Intelligence: n[5]}}
}
