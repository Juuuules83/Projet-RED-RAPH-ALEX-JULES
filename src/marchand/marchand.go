package marchand

import (
	"fmt"
	"projet-red/player"
	"projet-red/utils"
)

func Marchand(c *player.Character) {
	for {
		utils.ClearScreen()

		fmt.Println(utils.Bold + utils.Magenta + "===== LA FIKA =====" + utils.Reset)
		fmt.Printf(utils.Yellow+"Crédits ECTS : %d crédits ECTS\n"+utils.Reset, c.Money)
		fmt.Printf(utils.Cyan+"Inventaire : %d/%d\n\n"+utils.Reset,
			c.TotalItems(), utils.StockageMax)

		fmt.Println(utils.Bold + utils.Green + "1. Café du dev : 10 crédits ECTS" + utils.Reset)
		fmt.Println(utils.Bold + utils.Magenta + "2. Soupe de doute : 25 crédits ECTS" + utils.Reset)
		fmt.Println(utils.Bold + utils.Red + "3. Exploit de faille : 75 crédits ECTS" + utils.Reset)
		fmt.Println(utils.Bold + utils.Bold + "4. Clé SSH : 4 crédits ECTS" + utils.Reset)
		fmt.Println(utils.Bold + utils.Bold + "5. Carte graphique : 7 crédits ECTS" + utils.Reset)
		fmt.Println(utils.Bold + utils.Bold + "6. Fragment de code : 3 crédits ECTS" + utils.Reset)
		fmt.Println(utils.Bold + utils.Bold + "7. Ticket Ytrack : 1 crédits ECTS" + utils.Reset)
		fmt.Println(utils.Bold + utils.Red + "0. Retour" + utils.Reset)

		fmt.Print(utils.Yellow + "\nVotre choix : " + utils.Reset)
		choice := utils.ReadInt()

		switch choice {
		case 0:
			utils.ClearScreen()
			fmt.Println(utils.Cyan + "Retour" + utils.Reset)
			return

		case 1:
			if utils.Countfree == 0 {
				if c.AddInventory(utils.PotionVie, 1) {
					utils.Countfree++
					fmt.Println(utils.Green + "Café du dev [GRATUIT]" + utils.Reset)
				}
			} else if c.Money >= 10 {
				if c.AddInventory(utils.PotionVie, 1) {
					fmt.Println(utils.Green + "Café du dev acheté [10 crédits ECTS]" + utils.Reset)
					c.Money -= 10
				}
			} else {
				fmt.Println(utils.Red + "Tu n'as pas assez de crédits ECTS !" + utils.Reset)
			}

		case 2:
			if c.Money >= 25 {
				if c.AddInventory(utils.PotionPoison, 1) {
					fmt.Println(utils.Green + "Soupe de doute achetée [25 crédits ECTS]" + utils.Reset)
					c.Money -= 25
				}
			} else {
				fmt.Println(utils.Red + "Tu n'as pas assez de crédits ECTS !" + utils.Reset)
			}

		case 3:
			if c.Inventory[utils.FireBall] == 0 {
				if c.Money >= 75 {
					if c.AddInventory(utils.FireBall, 1) {
						fmt.Println(utils.Green + "Exploit de faille acheté [75 crédits ECTS]" + utils.Reset)
						c.Money -= 75
					}
				} else {
					fmt.Println(utils.Red + "Tu n'as pas assez de crédits ECTS !" + utils.Reset)
				}
			} else {
				fmt.Println(utils.Yellow + "Sort déjà acheté !" + utils.Reset)
			}

		case 4:
			if c.Money >= 4 {
				if c.AddInventory(utils.WolfFurr, 1) {
					fmt.Println(utils.Green + "Clé SSH achetée [4 crédits ECTS]" + utils.Reset)
					c.Money -= 4
				}
			} else {
				fmt.Println(utils.Red + "Tu n'as pas assez de crédits ECTS !" + utils.Reset)
			}

		case 5:
			if c.Money >= 7 {
				if c.AddInventory(utils.TrollSkin, 1) {
					fmt.Println(utils.Green + "Carte graphique achetée [7 crédits ECTS]" + utils.Reset)
					c.Money -= 7
				}
			} else {
				fmt.Println(utils.Red + "Tu n'as pas assez de crédits ECTS !" + utils.Reset)
			}

		case 6:
			if c.Money >= 3 {
				if c.AddInventory(utils.BoarLeather, 1) {
					fmt.Println(utils.Green + "Fragment de code acheté [3 crédits ECTS]" + utils.Reset)
					c.Money -= 3
				}
			} else {
				fmt.Println(utils.Red + "Tu n'as pas assez de crédits ECTS !" + utils.Reset)
			}

		case 7:
			if c.Money >= 1 {
				if c.AddInventory(utils.RavenFeather, 1) {
					fmt.Println(utils.Green + "Ticket Ytrack acheté [1 crédits ECTS]" + utils.Reset)
					c.Money -= 1
				}
			} else {
				fmt.Println(utils.Red + "Tu n'as pas assez de crédits ECTS !" + utils.Reset)
			}

		default:
			fmt.Println(utils.Red + "Option invalide. Veuillez réessayer." + utils.Reset)
		}

		utils.Pause()
	}
}
