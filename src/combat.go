package main

import "fmt"

//-------------------// TOUR DU GOBELIN //-------------------//
// goblinPattern : tour du gobelin.
// il inflige ses dégâts d'attaque.
// Tous les 3 tours, ses dégâts sont doublés.
func goblinPattern(goblin *Monster, player *Character, turn int) {
	damage := goblin.Attack

	if turn%3 == 0 {
		damage *= 2
	}

	player.Pv -= damage

	if player.Pv < 0 {
		player.Pv = 0
	}

	fmt.Printf("%s inflige %d dégâts à %s.\n", goblin.Name, damage, player.Name)
	fmt.Printf("%s : PV %d/%d\n", player.Name, player.Pv, player.PvMax)
}

//-------------------// FIN TOUR DU GOBELIN //-------------------//

//-------------------// TOUR DU PERSONNAGE //-------------------//
// characterTurn : gère un tour de joueur.
// Le joueur peut attaquer ou ouvrir son inventaire.	
func characterTurn(player *Character, monster *Monster) {
	for {
		clearScreen()

		printBoxTitle("COMBAT")
		printBoxLine(fmt.Sprintf("%s%s%s : %s%d/%d%s PV", Bold, player.Name, Reset, Red, player.Pv, player.PvMax, Reset))
		printBoxLine(fmt.Sprintf("%s%s%s : %s%d/%d%s PV", Bold, monster.Name, Reset, Red, monster.Pv, monster.PvMax, Reset))
		printBoxSeparator()

		printBoxLine(Yellow + "[1]" + Reset + "  Attaquer")
		printBoxLine(Yellow + "[2]" + Reset + "  Inventaire")

		printBoxSeparator()
		printBoxLine(Cyan + "Choisissez votre action..." + Reset)
		printBoxBottom()

		fmt.Print("\n  > ")

		var choice int
		if _, err := fmt.Scan(&choice); err != nil {
			clearScreen()
			printBoxTitle("ERREUR")
			printBoxLine(Red + "Veuillez entrer un nombre." + Reset)
			printBoxBottom()
			waitForEnter()
			continue
		}

		switch choice {
		case 1:
			const basicAttackDamage = 5

			monster.Pv -= basicAttackDamage
			if monster.Pv < 0 {
				monster.Pv = 0
			}

			fmt.Printf("%s utilise Attaque basique et inflige %d dégâts à %s.\n", player.Name, basicAttackDamage, monster.Name)
			fmt.Printf("%s : PV %d/%d\n", monster.Name, monster.Pv, monster.PvMax)
			return

		case 2:
			player.accessInventory()
			return

		default:
			clearScreen()
			printBoxTitle("ERREUR")
			printBoxLine(Red + "Option invalide. Veuillez réessayer." + Reset)
			printBoxBottom()
			waitForEnter()
		}
	}
}

//-------------------// FIN TOUR DU PERSONNAGE //-------------------//
