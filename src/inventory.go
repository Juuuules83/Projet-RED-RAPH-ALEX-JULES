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
