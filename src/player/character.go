package player

import (
	"fmt"
	"strings"
	"unicode"
	"projet-red/utils"
)

const (
	PotionVie    = "potion de vie"
	PotionPoison = "potion de poison"
	StockageMax  = 10
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
	c.Name = name
	c.Classe = classe
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
		PotionVie:    3,
		PotionPoison: 1,
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
		fmt.Print("Entrez le nom de votre personnage : ")
		name = utils.ReadLine()

		if isLettersOnly(name) {
			break
		}

		fmt.Println("Le nom doit contenir uniquement des lettres.")
	}

	name = formatName(name)

	var classe string

	for {
		fmt.Println("\nChoisissez votre classe :")
		fmt.Println("1. AI & Data")
		fmt.Println("2. Info")
		fmt.Println("3. Cyber")
		fmt.Print("Choix : ")

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

func (c *Character) AddInventory(itemName string, itemQuantity int) bool {
	if itemQuantity <= 0 || itemName == "" {
		return false
	}

	if c.TotalItems()+itemQuantity > StockageMax {
		return false
	}

	c.Inventory[itemName] += itemQuantity
	return true
}

func (c *Character) RemoveInventory(itemName string, itemQuantity int) bool {
	if itemQuantity <= 0 {
		return false
	}

	quantity, exists := c.Inventory[itemName]

	if !exists || quantity < itemQuantity {
		return false
	}

	quantity -= itemQuantity

	if quantity == 0 {
		delete(c.Inventory, itemName)
	} else {
		c.Inventory[itemName] = quantity
	}

	return true
}

func (c *Character) TotalItems() int {
	total := 0

	for _, quantity := range c.Inventory {
		total += quantity
	}

	return total
}

func (c *Character) UseLifePotion() string {
	quantity := c.Inventory[PotionVie]

	if quantity <= 0 {
		return "Vous n'avez plus de potion de vie."
	}

	c.Pv += 50

	if c.Pv > c.PvMax {
		c.Pv = c.PvMax
	}

	c.RemoveInventory(PotionVie, 1)

	return fmt.Sprintf("Potion utilisée. PV : %d/%d", c.Pv, c.PvMax)
}
