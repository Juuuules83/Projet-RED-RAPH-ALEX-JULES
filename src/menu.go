package main

import "fmt"

// --------------- MENU PRINCIPAL ---------------

// displayInfo affiche les informations du personnage dans un cadre.
func displayInfo(c Character) {
	clearScreen()

	printGameTitle()
	printBoxLine(Bold + Cyan + "INFORMATIONS DU PERSONNAGE" + Reset)
	printBoxSeparator()

	printBoxLine(fmt.Sprintf("%sNom%s    : %s%s%s", Cyan, Reset, Bold, c.Name, Reset))
	printBoxLine(fmt.Sprintf("%sClasse%s : %s%s%s", Cyan, Reset, Bold, c.Classe, Reset))
	printBoxLine(fmt.Sprintf("%sPV%s     : %s%d/%d%s", Cyan, Reset, Red, c.Pv, c.PvMax, Reset))
	printBoxLine(fmt.Sprintf("%sArgent%s : %s%d%s", Cyan, Reset, Green, c.Money, Reset))

	printBoxBottom()
	waitForEnter()
}

// mainMenu affiche le menu principal.
func mainMenu(player *Character) {
	for true {
		clearScreen()

		printGameTitle()

		printBoxLine(Bold + White + "                 MENU PRINCIPAL" + Reset)
		printBoxSeparator()

		printBoxLine(Yellow + "[1]" + Reset + "  Informations du personnage")
		printBoxLine(Yellow + "[2]" + Reset + "  Accéder à l'inventaire")
		printBoxLine(Yellow + "[3]" + Reset + "  Combat de test")
		printBoxLine(Yellow + "[0]" + Reset + "  Quitter")

		printBoxSeparator()
		printBoxLine(Cyan + "Choisissez votre destination..." + Reset)
		printBoxBottom()

		fmt.Print("\n  " + Bold + Cyan + ">" + Reset + " ")

		var chose int
		fmt.Scan(&chose)

		switch chose {
		case 0:
			clearScreen()
			printGameTitle()
			printBoxLine(Green + Bold + "Merci d'avoir joué !" + Reset)
			printBoxBottom()
			return

		case 1:
			displayInfo(*player)

		case 2:
			player.accessInventory()

		case 3:
			testCombat(player)

		default:
			clearScreen()
			printGameTitle()
			printBoxLine(Red + Bold + "ERREUR" + Reset)
			printBoxLine("Option invalide. Veuillez réessayer.")
			printBoxBottom()
			waitForEnter()
		}
	}
}
