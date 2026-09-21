package main

import "fmt"

const (
	PotionVie = "potion de vie"
)

// takePot utilise une potion de vie sur le personnage.
func (c *Character) takePot() {
	potQuantity, potCheck := c.Inventory[PotionVie]

	if !potCheck {
		fmt.Println(Red + "Vous n'avez pas de potion dans votre inventaire." + Reset)
		return
	}

	if potQuantity <= 0 {
		fmt.Println(Red + "Vous n'avez plus de potion de vie." + Reset)
		return
	}

	c.Pv += 50

	if c.Pv > c.PvMax {
		c.Pv = c.PvMax
	}

	c.Inventory[PotionVie]--

	fmt.Printf("%s✓ Potion utilisée !%s Vous avez maintenant %s%d/%d PV%s.\n",
		Green, Reset, Bold, c.Pv, c.PvMax, Reset)
}
