package player

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	PotionVie    = "potion de vie"
	PotionPoison = "potion de poison"
)

// Character représente le personnage du joueur.
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

// InitCharacter initialise un personnage avec ses valeurs de départ.
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

// CharacterCreation crée le personnage au lancement du jeu.
func CharacterCreation() Character {
	var character Character

	var name string
	for {
		fmt.Print("Entrez le nom de votre personnage (lettres uniquement) : ")
		if _, err := fmt.Scan(&name); err != nil {
			fmt.Println("Erreur de saisie. Veuillez réessayer.")
			continue
		}

		if isLettersOnly(name) {
			break
		}

		fmt.Println("Le nom ne doit contenir que des lettres.")
	}

	name = formatName(name)

	var classe string
	for {
		fmt.Println("Choisissez votre classe :")
		fmt.Println("[1] AI & Data")
		fmt.Println("[2] Info")
		fmt.Println("[3] Cyber")
		fmt.Print("> ")

		var choice int
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("Veuillez entrer un nombre.")
			continue
		}

		switch choice {
		case 1:
			classe = "AI & Data"
		case 2:
			classe = "Info"
		case 3:
			classe = "Cyber"
		default:
			fmt.Println("Option invalide. Veuillez réessayer.")
			continue
		}
		break
	}

	character.InitCharacter(name, classe)
	return character
}

// ------- DEBUT INVENTAIRE -------//

// AddInventory ajoute des objets sans dépasser la capacité maximale.
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

// RemoveInventory retire des objets et supprime automatiquement une ligne à 0.
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

// TotalItems retourne le nombre total d'objets transportés.
func (c *Character) TotalItems() int {
	total := 0
	for _, quantity := range c.Inventory {
		if quantity > 0 {
			total += quantity
		}
	}
	return total
}

// UseLifePotion utilise une potion de vie et renvoie le résultat de l'action.
func (c *Character) UseLifePotion() string {
	quantity, exists := c.Inventory[PotionVie]
	if !exists || quantity <= 0 {
		return "Vous n'avez plus de potion de vie."
	}

	c.Pv += 50
	if c.Pv > c.PvMax {
		c.Pv = c.PvMax
	}

	c.RemoveInventory(PotionVie, 1)
	return fmt.Sprintf("Potion utilisée : %d/%d PV.", c.Pv, c.PvMax)
}

// ------- FIN INVENTAIRE -------//
