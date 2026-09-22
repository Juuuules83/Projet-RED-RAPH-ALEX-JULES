package trader

import (
	"fmt"

	"projet-red/player"
)

const PotionViePrice = 10

func BuyLifePotion(character *player.Character) string {
	if character.Money < PotionViePrice {
		return "Vous n'avez pas assez d'argent."
	}
	if character.TotalItems() >= player.StockageMax {
		return "Votre inventaire est plein."
	}
	if !character.AddInventory(player.PotionVie, 1) {
		return "Impossible d'ajouter la potion."
	}
	character.Money -= PotionViePrice
	return fmt.Sprintf("Vous avez acheté une potion de vie pour %d €.", PotionViePrice)
}
