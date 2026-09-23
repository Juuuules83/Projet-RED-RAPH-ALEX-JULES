package main

import ("fmt"
		"projet-red/")

func (c *Character) Trader(ItemName string, price int) {

	//variable à supprimer, juste là pour enlever l'erreur
	var chose int

	switch chose {
	case 0:
		ClearScreen()
		fmt.Println("Retour")

		return

	case 1:
		countfree := 0
		if countfree >= 0{
			countfree++
			fmt.Println("potion de vie [GRATUIT]")
			c.AddInventory(PotionVie, 1)
		}else if c.Money >= 10 && StockageMax > TotalItems {
			fmt.Println("potion de vie [10]")
			c.AddInventory(PotionPoison, 1)
			c.Money -= 10
		}else{
			fmt.Println("Pas assez d'argent ou d'espace dans l'inventaire")
		}

	case 2:
		if c.Money >= 25 && StockageMax > TotalItems {
			fmt.Println("potion de poison [25]")
			c.AddInventory(PotionPoison, 1)
			c.Money -= 25
		}else{
			fmt.Println("Pas assez d'argent ou d'espace dans l'inventaire")
		}
		
	default:
		ClearScreen()
		fmt.Println("Option invalide. Veuillez réessayer.")
	}
}