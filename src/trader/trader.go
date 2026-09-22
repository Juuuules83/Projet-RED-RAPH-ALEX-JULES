package trader

import (
	"fmt"
	"projet-red/player"
)

const PotionViePrice = 10

// BuyLifePotion achète une potion de vie au marchand.
// L'argent n'est retiré que si l'achat peut réellement être effectué.
func BuyLifePotion(character *player.Character) string {
	if character.Money < PotionViePrice {
		return "Vous n'avez pas assez d'argent."
	}

	if character.TotalItems() >= player.StockageMax {
		return "Votre inventaire est plein."
	}

	if !character.AddInventory(player.PotionVie, 1) {
		return "Impossible d'ajouter l'objet à l'inventaire."
	}

	character.Money -= PotionViePrice
	return fmt.Sprintf("Vous avez acheté une potion de vie pour %d €.", PotionViePrice)
}
