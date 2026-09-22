package combat

import "projet-red/player"

func UseLifePotion(character *player.Character) string {
	return character.UseLifePotion()
}
