package main

import "fmt"

// accessInventory affiche l'inventaire du personnage et gère ses choix.
func (c *Character) accessInventory() {
	for true {
		clearScreen()

		fmt.Println("=== INVENTAIRE ===")
		for itemName, itemQuantity := range c.Inventaire {
			fmt.Printf("\t - %s: %d\n", itemName, itemQuantity)
		}

		fmt.Println()
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
			waitForEnter()
		default:
			fmt.Println()
			fmt.Println("Option invalide. Veuillez réessayer.")
			waitForEnter()
		}
	}
}
