package game


type Fighter struct {
	Character
}

//fighter with stat priority
func CreateFighter(m Fighter, s string, n []int) Fighter {
	return Fighter{Character{
		Name: s,
		Strength:     n[5],
		Constitution: n[4],
		Dexterity:    n[3],
		Charisma:     n[2],
		Wisdom:       n[1],
		Intelligence: n[0]}}
}

