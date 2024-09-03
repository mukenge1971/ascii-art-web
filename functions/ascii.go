package functions

import (
	"os"
	"strings"
)

func AsciiArt(police string, texte string) string {
	input, err := os.ReadFile(police + ".txt")
	if err != nil {
		os.Exit(0)
	}
	result := ""
	if !IsAski(texte) {
		result = "Il y a des caractères non conformes"
		return result
	}
	table := strings.Split(string(input), "\n")
	newtable := strings.Split(texte, "\r\n")
	for j := 0; j <= len(newtable)-1; j++ {
		if newtable[j] == "" {
			result = result + "\n"
		} else {
			for i := 1; i <= 8; i++ {
				for _, w := range newtable[j] {
					result = result + (table[(int(w)-32)*9+i])
				}
				result = result + "\n"
			}
		}
	}
	return result
}
