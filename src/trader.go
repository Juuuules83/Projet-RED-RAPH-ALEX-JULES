package main

import "fmt"

func (c *Character) Trader(ItemName string, price int) {

	//variable à supprimer, juste là pour enlever l'erreur
	var chose int

	switch chose {
	case 0:
		clearScreen()
		printGameTitle()
		printBoxLine(Green + Bold + "Retour" + Reset)
		printBoxBottom()
		return

	case 1:
		fmt.Println("potion de vie [10]")
		c.AddInventory(PotionVie, 1)
		c.Money -= 10
	default:
		clearScreen()
		printGameTitle()
		printBoxLine(Red + Bold + "ERREUR" + Reset)
		printBoxLine("Option invalide. Veuillez réessayer.")
		printBoxBottom()
		waitForEnter()
	}
}
