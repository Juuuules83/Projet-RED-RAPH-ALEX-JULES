package main
<<<<<<< HEAD

import "fmt"

type Character struct {
	Name       string
	Classe     string
	PvMax      int
	Pv         int
	Inventaire map[string]int
}

const (
	test = "potionde vie"
)

func (c *Character) initCharacter(name string, classe string) {
	c.Name = name
	c.Classe = classe
	switch c.Classe {
	case "mentor":
		c.PvMax = 200
		c.Pv = c.PvMax / 2
	case "étudiant":
		c.PvMax = 50
		c.Pv = c.PvMax / 2
	}
	c.Inventaire = map[string]int{test: 3, "ordinateur": 1}
}

func (c Character) displayInfo() {
	fmt.Println("=== INFORMATION ===")
	fmt.Printf("\t Nom: %s\n", c.Name)
	fmt.Printf("\t Classe: %s\n", c.Classe)
	fmt.Printf("\t PV: %d/%d\n", c.Pv, c.PvMax)
	fmt.Printf("\t Pv max : %d\n", c.PvMax)
}

func (c Character) accessInventory() {

	
	fmt.Println("=== INVENTAIRE ===")
	for itemName, itemQuantity := range c.Inventaire {
		fmt.Printf("\t - %s: %d\n", itemName, itemQuantity)
	}
	fmt.Println("\t 1. Utiliser une potion")
	fmt.Println("\t 0. Retour au menu principal")

	fmt.Print("Choisissez une option: ")
	var chose int
	fmt.Scan(&chose)

	switch chose {
	case 0:
		return
	case 1:
		c.takePot()
	default:
		fmt.Println("Option invalide. Veuillez réessayer.")
	}
}

func (c *Character) takePot() {
	potQuantity, potCheck := c.Inventaire[test]
	if !potCheck {
		fmt.Println("Vous n'avez pas de potion dans votre inventaire.")
		return
	}

	if potQuantity <= 0 {
		fmt.Println("Vous n'avez plus de potion dans votre inventaire.")
		return
	}

	c.Pv += 50
	if c.Pv > c.PvMax {
		c.Pv = c.PvMax
	}
	c.Inventaire[test]--
	// quoi faire quand quantité = 0 ?

	fmt.Printf("-1 Potion, vous avez %d/%d\n", c.Pv, c.PvMax)
}

func main() {
	var player Character
	player.initCharacter("Cyril", "mentor")
	for true {
		fmt.Println("=== Main Menu ===")
		fmt.Println("\t 1. Afficher les informations du personnage")
		fmt.Println("\t 2. Accéder à l'inventaire")
		fmt.Println("\t 0. Quitter")

		fmt.Print("Choisissez une option: ")
		var chose int
		fmt.Scan(&chose)
		
		switch chose {
		case 0:
			fmt.Println("Au revoir!")
			return	

		case 1:
			player.displayInfo()
		case 2:
			player.accessInventory()
		default:
			fmt.Println("Option invalide. Veuillez réessayer.")
		}
	}
}

// pour le scan quand il attend un INT il faut mettre un string qui converti en INT pour eviter les erreurs de scan
// au lieu de mettre plusieurs fois certains éléments, mettre des constantes pour les réutiliser partout, exemple pour les potions
=======
>>>>>>> 0972df57170c58a740dbcf751e4fec7b1f64525b
