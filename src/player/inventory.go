package player

import (
	"fmt"
	"projet-red/src/utils"
)

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

		switch choice {
		case 1:
			fmt.Println(c.UtiliserPotionVie())
			utils.Pause()
		case 2:
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

 func (c *Character) TotalItems() int {
	total := 0

	for _, quantity := range c.Inventory {
		total += quantity
	}

	return total
} 

func (c *Character) AddInventory(itemName string, itemQuantity int) bool {
	if itemQuantity <= 0 || itemName == "" {
		return false
	}

	if c.TotalItems()+itemQuantity > StockageMax {
		fmt.Println("MAIS TU ES MALADE GROS TU AS PLUS D'ESPACE LA, TU VEUX TE CASSER LE DOS ?")
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



