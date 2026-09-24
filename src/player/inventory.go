package player

import (
	"fmt"
	"projet-red/utils"
)

func Inventaire(c *Character, depuisCombat bool) {
	for {
		utils.ClearScreen()
		fmt.Println("===== INVENTAIRE =====")
		fmt.Printf("Objets : %d/%d\n\n", c.TotalItems(), utils.StockageMax)

		if len(c.Inventory) == 0 {
			fmt.Println("L'inventaire est vide.")
		} else {
			fmt.Println("Objets possédés :")
			for item, quantity := range c.Inventory {
				fmt.Printf("- %s : x%d\n", item, quantity)
			}
		}

		fmt.Println()
		fmt.Println("1. Utiliser une potion de vie")

		if depuisCombat {
			fmt.Println("2. Retour au combat")
		} else {
			fmt.Println("2. Retour au menu")
		}

		fmt.Print("\nChoix : ")

<<<<<<< HEAD
		switch utils.Choose {
=======
	choice := utils.ReadInt()
	switch choice {
>>>>>>> 5c9562dfffecd850e13546a4747f99b6b72973e0
		case 1:
			fmt.Println(c.UtiliserPotionVie())
			utils.Pause()
		case 2:
			return
		default:
			fmt.Println("Choix invalide.")
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
		fmt.Println("MAIS TU ES MALADE GROS TU AS PLUS D'ESPACE LA, TU VEUX TE CASSER LE DOS ?")
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
<<<<<<< HEAD
=======


func (c *Character) UtiliserPotionVie() string {
    quantity := c.Inventory[utils.PotionVie]

    if quantity <= 0 {
        return "T'as plus de potion mon reuf... tu vas crever"
    }

    c.Pv += 50

    if c.Pv > c.PvMax {
        c.Pv = c.PvMax
    }

    c.RemoveInventory(utils.PotionVie, 1)

    return fmt.Sprintf(
        "Potion utilisée. PV : %d/%d",
        c.Pv,
        c.PvMax,
    )
}
>>>>>>> 5c9562dfffecd850e13546a4747f99b6b72973e0
