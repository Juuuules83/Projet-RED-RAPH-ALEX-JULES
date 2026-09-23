package player

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
		}
	}
}

func (c *Character) AddInventory (ItemName string, ItemQuantity int){
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
