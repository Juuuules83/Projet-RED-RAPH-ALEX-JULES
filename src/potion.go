package main

import (
	"fmt"
	"time"
)

const (
	PotionVie    = "potion de vie"
	PotionPoison = "potion de poison"
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


// poisonPot utilise une potion de poison sur l'adversaire (target), pas sur le joueur.
func (c *Character) poisonPot(target *Monster) {
	potQuantity, potCheck := c.Inventory[PotionPoison]

	if !potCheck || potQuantity <= 0 {
		fmt.Println(Red + "Vous n'avez pas de potion de poison dans votre inventaire." + Reset)
		return
	}

	c.Inventory[PotionPoison]--

	fmt.Printf("%s☠ Vous lancez une potion de poison sur %s...%s\n", Green, target.Name, Reset)

	for second := 1; second <= 3; second++ {
		time.Sleep(1 * time.Second)

		target.Pv -= 10
		if target.Pv < 0 {
			target.Pv = 0
		}

		fmt.Printf("%s☠ Poison (%d/3)%s : %s%s%s %s%d/%d PV%s\n",
			Green, second, Reset, Bold, target.Name, Reset, Red, target.Pv, target.PvMax, Reset)

		if target.Pv <= 0 {
			fmt.Printf("%s%s succombe au poison !%s\n", Red, target.Name, Reset)
			return
		}
	}
}

