package main

import (
	"fmt"
	"projet-red/player"
	"projet-red/ui"
	"projet-red/utils"
)

func main() {

	// MAIN MENU
	for {
		utils.ClearScreen()
		utils.PrintTitle()
		fmt.Println(utils.Green + utils.Bold + "[ 1 ] - Commencer" + utils.Reset)
		fmt.Println(utils.Red + utils.Bold + "[ 2 ] - Quitter" + utils.Reset)
		fmt.Print(utils.Yellow + "\nChoix : " + utils.Reset)

		switch utils.ReadInt() {
		case 1:
			utils.ClearScreen()
			character := player.CharacterCreation()
			ui.MenuPrincipal(&character)
			return
		case 2:
			utils.ClearScreen()
			fmt.Println(utils.Green + "Au revoir !" + utils.Reset)
			return
		default:
			fmt.Println(utils.Red + "Choix invalide." + utils.Reset)
			utils.Pause()
		}
	}
}
