package marchant

import (
	"fmt"
	"projet-red/player"
	"projet-red/utils"
)

const PrixPotionVie = 10

func AcheterPotionVie(c *player.Character) string {
	if c.Money < PrixPotionVie {
		return "Vous n'avez pas assez d'argent."
	}
	if c.TotalItems() >= player.StockageMax {
		return "Votre inventaire est plein."
	}

	c.Money -= PrixPotionVie
	return fmt.Sprintf("Vous avez acheté une potion de vie pour %d €.", PrixPotionVie)
}

func Marchand(c *player.Character) {

	//variable à supprimer, juste là pour enlever l'erreur
	var chose int

	switch chose {
	case 0:
		utils.ClearScreen()
		fmt.Println("Retour")

		return

	case 1:
		countfree := 0
		if countfree >= 0 {
			countfree++
			fmt.Println("potion de vie [GRATUIT]")
			c.AddInventory(player.PotionVie, 1)
		} else if c.Money >= 10 && player.StockageMax > c.TotalItems() {
			fmt.Println("potion de vie [10]")
			c.AddInventory(player.PotionPoison, 1)
			c.Money -= 10
		} else {
			fmt.Println("Pas assez d'argent ou d'espace dans l'inventaire")
		}

	case 2:
		if c.Money >= 25 && player.StockageMax > c.TotalItems() {
			fmt.Println("potion de poison [25]")
			c.AddInventory(player.PotionPoison, 1)
			c.Money -= 25
		} else {
			fmt.Println("Pas assez d'argent ou d'espace dans l'inventaire")
		}

	default:
		utils.ClearScreen()
		fmt.Println("Option invalide. Veuillez réessayer.")
	}
}
