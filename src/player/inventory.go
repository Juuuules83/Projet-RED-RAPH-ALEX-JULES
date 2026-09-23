package player

import (
	"fmt"
	"projet-red/utils"
)

func Inventaire(c *Character, depuisCombat bool) {
	for {
		utils.ClearScreen()
		fmt.Println("===== INVENTAIRE =====")
		fmt.Printf("Objets : %d/%d\n\n", c.TotalItems(), StockageMax)

		if len(c.Inventory) == 0 {
			fmt.Println("L'inventaire est vide.")
		} else {
			fmt.Println("Objets possédés :")
			for item, quantity := range c.Inventory {
				fmt.Printf("- %s : x%d\n", item, quantity)
			}
		}

		fmt.Println()
		fmt.Println("1. Utiliser une potion de vie")

		if depuisCombat {
			fmt.Println("2. Retour au combat")
		} else {
			fmt.Println("2. Retour au menu")
		}

		fmt.Print("\nChoix : ")
		choice := utils.ReadInt()

		switch choice {
		case 1:
			fmt.Println(c.UseLifePotion())
			utils.Pause()
		case 2:
			return
		default:
			fmt.Println("Choix invalide.")
			utils.Pause()
		}
	}
}
