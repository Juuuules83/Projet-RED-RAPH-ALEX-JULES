package player

import (
	"fmt"
	"projet-red/utils"
	"strings"
	"unicode"
)

type Character struct {
	Name      string
	Classe    string
	PvMax     int
	Pv        int
	Niveau    int
	Sorts     []string
	Inventory map[string]int
	Money     int
}

func (c *Character) InitCharacter(name string, classe string) {
	c.Name = utils.Bold + name + utils.Reset
	c.Classe = utils.Bold + utils.Cyan + classe + utils.Reset
	c.Money = 100
	c.Niveau = 1
	c.Sorts = []string{"Coup de Poing"}

	switch classe {
	case "AI & Data":
		c.PvMax = 100
	case "Info":
		c.PvMax = 80
	case "Cyber":
		c.PvMax = 120
	default:
		c.PvMax = 80
	}

	c.Pv = c.PvMax / 2

	c.Inventory = map[string]int{
		utils.PotionVie:    3,
		utils.PotionPoison: 1,
	}
}

func isLettersOnly(value string) bool {
	if value == "" {
		return false
	}

	for _, r := range value {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func formatName(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return ""
	}

	runes := []rune(strings.ToLower(name))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func CharacterCreation() Character {
	var character Character
	var name string

	for {
		fmt.Println(utils.Magenta + utils.Bold + "╔══════════ IDENTIFICATION ══════════╗" + utils.Reset)
		fmt.Println("Bienvenue dans l’aventure.")
		fmt.Print(utils.Cyan + "Entrez le nom de votre personnage : " + utils.Reset)
		name = utils.ReadLine()

		if isLettersOnly(name) {
			break
		}

		fmt.Println("Le nom doit contenir uniquement des lettres.")
	}

	name = formatName(name)
	Lore(name)

	var classe string

	for {
		fmt.Println("\n" + utils.Magenta + utils.Bold + "★ CHOIX DE LA CLASSE ★" + utils.Reset)
		fmt.Println(utils.Green + "1. AI & Data" + utils.Reset)
		fmt.Println(utils.Cyan + "2. Info" + utils.Reset)
		fmt.Println(utils.Red + "3. Cyber" + utils.Reset)
		fmt.Print(utils.Yellow + "Choix : " + utils.Reset)

		choice := utils.ReadInt()

		switch choice {
		case 1:
			classe = "AI & Data"
		case 2:
			classe = "Info"
		case 3:
			classe = "Cyber"
		default:
			fmt.Println("Choix invalide.")
			continue
		}
		break
	}
	character.InitCharacter(name, classe)
	return character
}
