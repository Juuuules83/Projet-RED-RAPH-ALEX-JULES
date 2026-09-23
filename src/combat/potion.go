package combat

import "projet-red/src/player"

func UseLifePotion(c *player.Character) string {
	return c.UtiliserPotionVie()
}
