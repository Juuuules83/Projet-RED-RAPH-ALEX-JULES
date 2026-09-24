package ui

import (
	"fmt"
	"projet-red/player"
	"projet-red/utils"
	"projet-red/combat"
	"projet-red/marchand"
)

func MenuPrincipal(c *player.Character) {
	for {
		utils.ClearScreen()
		utils.PrintTitle()

		fmt.Println("1. Informations")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Ymatch")
		fmt.Println("4. Combat")
		fmt.Println("5. Quitter")
		fmt.Print("\nChoix : ")

		choice := utils.ReadInt()

		switch choice {
		case 1:
			AfficherPersonnage(c)
		case 2:
			player.Inventaire(c, false)
		case 3:
			marchand.Marchand(c)
		case 4:
			combat.Combat(c)
		case 5:
			fmt.Println("Ciao !")
			return
		default:
			fmt.Println("Choix invalide.")
			utils.Pause()
		}
	}
}

func AfficherPersonnage(c *player.Character) {
	utils.ClearScreen()
	fmt.Println("===== PERSONNAGE =====")
	fmt.Println("Nom :", c.Name)
	fmt.Println("Classe :", c.Classe)
	fmt.Printf("PV : %d/%d\n", c.Pv, c.PvMax)
	fmt.Println("Niveau :", c.Niveau)
	fmt.Println("Argent :", c.Money, "€")
	utils.Pause()
}