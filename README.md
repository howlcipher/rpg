# rpg
This is table top like rpg in Go.  I am trying to create a game in which chacters can be created via text files and loaded into the game to battle.

Phase 1: Make dice struct
  -Dice can be sided however many is desired
  -Dice can roll to generate a random number 1 to the amount of sides
  
Phase 2:
  -Create character attributes struct
  -generate stats using created dice
  
Phase 3:
  -convert character struct to json
  -save character struct to a text file
  -convert json to strings and int
  -load character

Phase 4:
  -Create character classes
  -Classes get attribute modifiers based on type of character class (mage, fighter, etc)
  
Phase 5:
  -Construct user menu
  
Phase 6:
  -Have character simulate battle
  
Phase TBA:
  -More Features
  -game map
  -character inventory
  -spells
  -multi character fights
