package marchand

import (
	"fmt"
	"projet-red/player"
	"projet-red/utils"
)

func Artisant(c *player.Character) {
	for {
		utils.ClearScreen()

		fmt.Println(utils.Bold + utils.Magenta + "===== YMATCH =====" + utils.Reset)
		fmt.Printf(utils.Yellow+"Crédits ECTS : %d crédits ECTS\n"+utils.Reset, c.Money)
		fmt.Printf(utils.Cyan+"Inventaire : %d/%d\n\n"+utils.Reset, c.TotalItems(), utils.StockageMax)

		fmt.Println("1. " + utils.Bold + "[MENTOR] " + utils.Reset + utils.Cyan + "Vito" + utils.Reset + "\n\tObjets nécessaires : x1 Ticket Ytrack, x1 Fragment de code")
		fmt.Println("2. " + utils.Bold + "[MENTOR] " + utils.Reset + utils.Yellow + "Cyril" + utils.Reset + "\n\tObjets nécessaires : x2 Clés SSH, x1 Carte graphique")
		fmt.Println("3. " + utils.Bold + "[MENTOR] " + utils.Reset + utils.Green + "Lilian" + utils.Reset + "\n\tObjets nécessaires : x1 Clé SSH, x1 Fragment de code")
		fmt.Println(utils.Red + "0. Retour" + utils.Reset)

		fmt.Print(utils.Yellow + "\nVotre choix : " + utils.Reset)
		choice := utils.ReadInt()

		switch choice {
		case 0:
			utils.ClearScreen()
			fmt.Println(utils.Cyan + "Retour" + utils.Reset)
			return

		case 1:
			if c.Inventory["Mentor Vito"] > 0 {
				fmt.Println(utils.Yellow + "Vito est déjà recruté !" + utils.Reset)
			} else if c.Inventory[utils.RavenFeather] < 1 || c.Inventory[utils.BoarLeather] < 1 {
				fmt.Println(utils.Red + "Il te manque des objets pour recruter Vito." + utils.Reset)
			} else {
				c.RemoveInventory(utils.RavenFeather, 1)
				c.RemoveInventory(utils.BoarLeather, 1)
				c.AddInventory("Mentor Vito", 1)
				fmt.Println(utils.Green + "Vito a rejoint ton équipe !" + utils.Reset)
			}

		case 2:
			if c.Inventory["Mentor Cyril"] > 0 {
				fmt.Println(utils.Yellow + "Cyril est déjà recruté !" + utils.Reset)
			} else if c.Inventory[utils.WolfFurr] < 2 || c.Inventory[utils.TrollSkin] < 1 {
				fmt.Println(utils.Red + "Il te manque des objets pour recruter Cyril." + utils.Reset)
			} else {
				c.RemoveInventory(utils.WolfFurr, 2)
				c.RemoveInventory(utils.TrollSkin, 1)
				c.AddInventory("Mentor Cyril", 1)
				fmt.Println(utils.Green + "Cyril a rejoint ton équipe !" + utils.Reset)
			}

		case 3:
			if c.Inventory["Mentor Lilian"] > 0 {
				fmt.Println(utils.Yellow + "Lilian est déjà recruté !" + utils.Reset)
			} else if c.Inventory[utils.WolfFurr] < 1 || c.Inventory[utils.BoarLeather] < 1 {
				fmt.Println(utils.Red + "Il te manque des objets pour recruter Lilian." + utils.Reset)
			} else {
				c.RemoveInventory(utils.WolfFurr, 1)
				c.RemoveInventory(utils.BoarLeather, 1)
				c.AddInventory("Mentor Lilian", 1)
				fmt.Println(utils.Green + "Lilian a rejoint ton équipe !" + utils.Reset)
			}

		default:
			fmt.Println(utils.Red + "Choix invalide." + utils.Reset)
		}
		utils.Pause()
	}
}
