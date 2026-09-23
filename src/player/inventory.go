package player

<<<<<<< HEAD
import "fmt"

// accessInventory affiche l'inventaire du personnage et gère ses choix.
func (c *Character) accessInventory() {
	for true {
		clearScreen()

		printBoxTitle("INVENTAIRE")

		for itemName, itemQuantity := range c.Inventaire {
			printBoxLine(fmt.Sprintf("%s•%s %-22s : %s%d%s",
				Cyan, Reset,
				itemName,
				Bold, itemQuantity, Reset,
			))
		}

		printBoxSeparator()

		printBoxLine(Yellow + "[1]" + Reset + "  Utiliser une potion")
		printBoxLine(Yellow + "[0]" + Reset + "  Retour au menu principal")

		printBoxSeparator()
		printBoxLine(Cyan + "Choisissez une option..." + Reset)
		printBoxBottom()

		fmt.Print("\n  > ")

		var chose int
		fmt.Scan(&chose)

		switch chose {
		case 0:
			return

		case 1:
			clearScreen()

			printBoxTitle("POTION DE VIE")
			c.takePot()
			printBoxBottom()

			waitForEnter()

		default:
			clearScreen()

			printBoxTitle("ERREUR")
			printBoxLine(Red + "Option invalide." + Reset)
			printBoxBottom()

			waitForEnter()
=======
import (
	"fmt"
	"projet-red/utils"
)

func Inventaire(c *Character, depuisCombat bool) {
	for {
		utils.ClearScreen()
		fmt.Println("===== INVENTAIRE =====")
		fmt.Printf("Objets : %d/%d\n\n", c.TotalItems(), StockageMax)

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
		choice := utils.ReadInt()

		switch choice {
		case 1:
			fmt.Println(c.UseLifePotion())
			utils.Pause()
		case 2:
			return
		default:
			fmt.Println("Choix invalide.")
			utils.Pause()
>>>>>>> d9ccf5d6011d9e3c51cab7a86f889b06aae79a3d
		}
	}
}

func (c *Character) AddInventory (ItemName string, ItemQuantity int){
<<<<<<< HEAD
	var TotalItems int
	for _, value := range(c.Inventory){
		TotalItems += value
	}
	if !(TotalItems + ItemQuantity)<= StockageMax{
		fmt.Println("MAIS TU ES MALADE GROS TU AS PLUS D'ESPACE LA, TU VEUX TE CASSER LE DOS !?")
		return
	}
	check := c.Inventory(ItemName)
	if (check){
		c.Inventory(ItemName) += ItemQuantity
	}else{
		c.Inventory(ItemName)= ItemQuantity
	}
	fmt.Println("+1", ItemName)	
}

func (c *Character) RemoveInventory (ItemName string, ItemQuantity int){
	ItemQuantity, InvCheck := c.Inventory[ItemName]
	if (InvCheck == false){
		fmt.Println(ItemName)
		return
	}else if (InvCheck && InvQuantity < ItemQuantity){
		fmt.Println("Quantité insuffisante")
		return
	}
	if (InvQuantity - ItemQuantity) == 0{
		delete(c.Inventory, ItemName)
		return
	}else{
		c.Inventory[ItemName]-= ItemQuantity
	}
	fmt.Println("-1", ItemName)	
}
=======

   for _, value := c.Inventory{
       TotalItems += value
   }
   If !(TotalItems + ItemQuantity)<= StockageMax{
       fmt.Println("MAIS TU ES MALADE GROS TU AS PLUS D'ESPACE LA, TU VEUX TE CASSER LE DOS ?")
       return
   }
   check := c.Inventory(ItemName)
   If (check){
       c.Inventory(ItemName) += ItemQuantity
   }else{
       c.Inventory(ItemName)= ItemQuantity
   }
   fmt.Println("+1" ItemName)  
}

func (c *Character) RemoveInventory (ItemName string, ItemQuantity int){
   ItemQuantity, InvCheck := c.Inventory[ItemName]
   if (InvCheck == false){
       fmt.Println(ItemName)
       return
   }else if (InvCheck && InvQuantity < ItemQuantity){
       fmt.Println("Quantité insuffisante")
       return
   }
   if (InvQuantity - ItemQuantity) == 0{
       delete(c.Inventory, ItemName)
       return
   }else{
       c.Inventory[ItemName]-= ItemQuantity
   }
   fmt.Println("-1", ItemName)
}
>>>>>>> d9ccf5d6011d9e3c51cab7a86f889b06aae79a3d
