package player

import (
	"fmt"
	"projet-red/utils"
)

func Inventaire(c *Character, depuisCombat bool) {
	for {
		utils.ClearScreen()
		fmt.Println(utils.Bold + utils.Magenta + "===== INVENTAIRE =====" + utils.Reset)
		fmt.Printf(utils.Cyan+"Objets : %d/%d\n\n"+utils.Reset, c.TotalItems(), utils.StockageMax)
		if len(c.Inventory) == 0 {
			fmt.Println(utils.Yellow + "L'inventaire est vide." + utils.Reset)
		} else {
			fmt.Println(utils.Green + "Objets possédés :" + utils.Reset)
			for item, quantity := range c.Inventory {
				fmt.Printf(utils.Cyan+"- "+utils.Reset+"%s"+utils.Yellow+" : x%d\n"+utils.Reset, item, quantity)
			}
		}

		fmt.Println()
		fmt.Println(utils.Green + "1. Boire un Café" + utils.Reset)
		if depuisCombat {
			fmt.Println(utils.Red + "2. Retour au combat" + utils.Reset)
		} else {
			fmt.Println(utils.Red + "2. Retour au menu" + utils.Reset)
		}
		fmt.Print(utils.Yellow + "\nChoix : " + utils.Reset)

		choice := utils.ReadInt()
		switch choice {
		case 1:
			fmt.Println(c.UtiliserPotionVie())
			utils.Pause()
		case 2:
			return
		default:
			fmt.Println(utils.Red + "Choix invalide." + utils.Reset)
			utils.Pause()
		}
	}
}

func (c *Character) TotalItems() int {
	total := 0
	for _, quantity := range c.Inventory {
		total += quantity
	}
	return total
}

func (c *Character) AddInventory(itemName string, itemQuantity int) bool {
	if itemQuantity <= 0 || itemName == "" {
		return false
	}

	if c.TotalItems()+itemQuantity > utils.StockageMax {
		fmt.Println(utils.Red + "MAIS TU ES MALADE GROS TU AS PLUS D'ESPACE LA, TU VEUX TE CASSER LE DOS ?" + utils.Reset)
		return false
	}
	c.Inventory[itemName] += itemQuantity
	return true
}

func (c *Character) RemoveInventory(itemName string, itemQuantity int) bool {
	if itemQuantity <= 0 {
		return false
	}

	quantity, exists := c.Inventory[itemName]
	if !exists || quantity < itemQuantity {
		return false
	}

	quantity -= itemQuantity
	if quantity == 0 {
		delete(c.Inventory, itemName)
	} else {
		c.Inventory[itemName] = quantity
	}
	return true
}

func (c *Character) UtiliserPotionVie() string {
	quantity := c.Inventory[utils.PotionVie]

	if quantity <= 0 {
		return "T’as plus de Café du dev, mon reuf... tu vas manquer d’énergie !"
	}

	c.Pv += 50
	if c.Pv > c.PvMax {
		c.Pv = c.PvMax
	}

	c.RemoveInventory(utils.PotionVie, 1)
	return fmt.Sprintf(
		"miam le café ! PV : %d/%d",
		c.Pv,
		c.PvMax,
	)
}
