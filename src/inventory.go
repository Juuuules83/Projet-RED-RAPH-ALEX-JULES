package main

import "fmt"

// accessInventory affiche l'inventaire du personnage et gère ses choix.
func (c *Character) accessInventory() {
	for true {
		clearScreen()

		printBoxTitle("INVENTAIRE")

		for itemName, itemQuantity := range c.Inventaire {
			printBoxLine(fmt.Sprintf("%s•%s %-22s : %s%d%s",
				Cyan, Reset,
				itemName,
				Bold, itemQuantity, Reset,
			))
		}

		printBoxSeparator()

		printBoxLine(Yellow + "[1]" + Reset + "  Utiliser une potion")
		printBoxLine(Yellow + "[0]" + Reset + "  Retour au menu principal")

		printBoxSeparator()
		printBoxLine(Cyan + "Choisissez une option..." + Reset)
		printBoxBottom()

		fmt.Print("\n  > ")

		var chose int
		fmt.Scan(&chose)

		switch chose {
		case 0:
			return

		case 1:
			clearScreen()

			printBoxTitle("POTION DE VIE")
			c.takePot()
			printBoxBottom()

			waitForEnter()

		default:
			clearScreen()

			printBoxTitle("ERREUR")
			printBoxLine(Red + "Option invalide." + Reset)
			printBoxBottom()

			waitForEnter()
		}
	}
}

// AddInventory ajoute une quantité d'un objet à l'inventaire du personnage.
func (c *Character) AddInventory(itemName string, itemQuantity int) {
	if itemQuantity <= 0 {
		return
	}
	var totalItems int
	for _, quantity := range c.Inventaire {
		totalItems += quantity
	}
	if StockageMax > 0 && totalItems+itemQuantity > StockageMax {
		fmt.Println("Vous n'avez plus assez d'espace dans votre inventaire.")
		return
	}
	if c.Inventaire == nil {
		c.Inventaire = make(map[string]int)
	}
	c.Inventaire[itemName] += itemQuantity
	fmt.Printf("+%d %s\n", itemQuantity, itemName)
}
