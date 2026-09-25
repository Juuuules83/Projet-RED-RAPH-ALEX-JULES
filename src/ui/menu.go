package ui

import (
	"fmt"
	"projet-red/combat"
	"projet-red/marchand"
	"projet-red/player"
	"projet-red/utils"
)

func MenuPrincipal(c *player.Character) {
	for {
		utils.ClearScreen()
		utils.PrintTitle()

		fmt.Println(utils.Bold + "[ 1 ] - Informations" + utils.Reset)
		fmt.Println(utils.Bold + "[ 2 ] - Inventaire" + utils.Reset)
		fmt.Println(utils.Bold + "[ 3 ] - La Fika" + utils.Reset)
		fmt.Println(utils.Bold + "[ 4 ] - Ymatch" + utils.Reset)
		fmt.Println(utils.Bold + "[ 5 ] - Combat" + utils.Reset)
		fmt.Println(utils.Red + utils.Bold + "[ 6 ] - Quitter" + utils.Reset)
		fmt.Print(utils.Yellow + "\nChoix : " + utils.Reset)

		choice := utils.ReadInt()

		switch choice {
		case 1:
			AfficherPersonnage(c)
		case 2:
			player.Inventaire(c, false)
		case 3:
			marchand.Marchand(c)
		case 4:
			marchand.Artisant(c)
		case 5:
			combat.Combat(c)
		case 6:
			fmt.Println(utils.Green + "Ciao !" + utils.Reset)
			return
		default:
			fmt.Println(utils.Red + "Choix invalide." + utils.Reset)
			utils.Pause()
		}
	}
}

func AfficherPersonnage(c *player.Character) {
	utils.ClearScreen()
	fmt.Println(utils.Magenta + "===== PERSONNAGE =====" + utils.Reset)
	fmt.Println(utils.Cyan+"Nom : "+utils.Reset, c.Name)
	fmt.Println(utils.Cyan+"Classe : "+utils.Reset, c.Classe)
	fmt.Printf(utils.Green+"PV : %d/%d\n"+utils.Reset, c.Pv, c.PvMax)
	fmt.Println(utils.Cyan+"Niveau : "+utils.Reset, c.Niveau)
	fmt.Println(utils.Yellow+"Crédits ECTS : "+utils.Reset, c.Money, "crédits ECTS")
	utils.Pause()
}
