package main

import "fmt"

const (
	PotionVie = "potion de vie"
)

// takePot utilise une potion de vie sur le personnage.
func (c *Character) takePot() {
	potQuantity, potCheck := c.Inventaire[PotionVie]

	if !potCheck {
		fmt.Println("Vous n'avez pas de potion dans votre inventaire.")
		return
	}

	if potQuantity <= 0 {
		fmt.Println("Vous n'avez plus de potion dans votre inventaire.")
		return
	}

	c.Pv += 50
	if c.Pv > c.PvMax {
		c.Pv = c.PvMax
	}
	c.Inventaire[PotionVie]--

	fmt.Printf("-1 Potion, vous avez %d/%d PV\n", c.Pv, c.PvMax)
}
