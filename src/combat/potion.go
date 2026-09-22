package combat

import "projet-red/player"

// UseLifePotion utilise une potion de vie du joueur.
func UseLifePotion(character *player.Character) string {
	return character.UseLifePotion()
}
