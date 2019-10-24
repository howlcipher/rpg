package game

import (
	"encoding/json"
	"os"
)

func SaveCharacter(c string, d string) {
	file, err := os.Create(d + ".txt")
	if err != nil { // handle the error here return
	}
	file.WriteString(c)
}

func SaveStats(c Character, d string) {
	s, _ := json.MarshalIndent(c, "", "\t")
	file, err := os.Create(d + ".txt")
	if err != nil { // handle the error here return
	}
	file.Write(s)

}