package trader

import (
	"fmt"

	"projet-red/player"
)

const PotionViePrice = 10

func BuyLifePotion(c *player.Character) string {
	if c.Money < PotionViePrice {
		return "Vous n'avez pas assez d'argent."
	}
	if c.TotalItems() >= player.StockageMax {
		return "Votre inventaire est plein."
	}
/* 	if !c.AddInventory(player.PotionVie, 1) {
		return "Impossible d'ajouter la potion."
	} */
	c.Money -= PotionViePrice
	return fmt.Sprintf("Vous avez acheté une potion de vie pour %d €.", PotionViePrice)
}
