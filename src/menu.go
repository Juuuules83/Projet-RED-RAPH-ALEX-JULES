package main

import "fmt"

// displayInfo affiche les informations du personnage.
func displayInfo(c Character) {
	clearScreen()

	fmt.Println("=== INFORMATION ===")
	fmt.Printf("\t Nom: %s\n", c.Name)
	fmt.Printf("\t Classe: %s\n", c.Classe)
	fmt.Printf("\t PV: %d/%d\n", c.Pv, c.PvMax)
	fmt.Printf("\t Pv max : %d\n", c.PvMax)

	waitForEnter()
}

// mainMenu affiche le menu principal et gère les choix du joueur.
func mainMenu(player *Character) {
	for true {
		clearScreen()

		fmt.Println("=== Main Menu ===")
		fmt.Println("\t 1. Afficher les informations du personnage")
		fmt.Println("\t 2. Accéder à l'inventaire")
		fmt.Println("\t 0. Quitter")

		fmt.Print("Choisissez une option: ")

		var chose int
		fmt.Scan(&chose)

		switch chose {
		case 0:
			clearScreen()
			fmt.Println("Au revoir!")
			return
		case 1:
			displayInfo(*player)
		case 2:
			player.accessInventory()
		default:
			fmt.Println()
			fmt.Println("Option invalide. Veuillez réessayer.")
			waitForEnter()
		}
	}
}
