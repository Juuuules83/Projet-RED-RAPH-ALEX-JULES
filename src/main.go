package main

import (
	"fmt"

	"projet-red/combat"
	"projet-red/player"
	"projet-red/trader"
	"projet-red/utils"
)

func main() {
	character := player.CharacterCreation()
	MenuPrincipal(&character)
}

func MenuPrincipal(c *player.Character) {
	for {
		utils.ClearScreen()
		utils.PrintTitle()

		fmt.Println("1. Informations du personnage")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Combat")
		fmt.Println("5. Quitter")
		fmt.Print("\nChoix : ")

		choice := utils.ReadInt()

		switch choice {
		case 1:
			AfficherPersonnage(c)
		case 2:
			Inventaire(c, false)
		case 3:
			Marchand(c)
		case 4:
			Combat(c)
		case 5:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
			utils.Pause()
		}
	}
}

func AfficherPersonnage(c *player.Character) {
	utils.ClearScreen()
	fmt.Println("===== PERSONNAGE =====")
	fmt.Println("Nom :", c.Name)
	fmt.Println("Classe :", c.Classe)
	fmt.Printf("PV : %d/%d\n", c.Pv, c.PvMax)
	fmt.Println("Niveau :", c.Niveau)
	fmt.Println("Argent :", c.Money, "€")
	utils.Pause()
}

func Inventaire(c *player.Character, depuisCombat bool) {
	for {
		utils.ClearScreen()
		fmt.Println("===== INVENTAIRE =====")
		fmt.Printf("Objets : %d/%d\n\n", c.TotalItems(), player.StockageMax)

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
		}
	}
}

func Marchand(c *player.Character) {
	for {
		utils.ClearScreen()
		fmt.Println("===== MARCHAND =====")
		fmt.Println("Argent :", c.Money, "€")
		fmt.Printf("Inventaire : %d/%d\n\n", c.TotalItems(), player.StockageMax)

		fmt.Println("1. Potion de vie - 10 €")
		fmt.Println("2. Retour")
		fmt.Print("\nChoix : ")

		choice := utils.ReadInt()

		switch choice {
		case 1:
			fmt.Println(trader.BuyLifePotion(c))
			utils.Pause()
		case 2:
			return
		default:
			fmt.Println("Choix invalide.")
			utils.Pause()
		}
	}
}

func Combat(c *player.Character) {
	monster := combat.InitGoblin()
	turn := 1

	for {
		utils.ClearScreen()
		fmt.Println("===== COMBAT =====")
		fmt.Println()
		fmt.Printf("%s : %d/%d PV\n", c.Name, c.Pv, c.PvMax)
		fmt.Printf("%s : %d/%d PV\n", monster.Name, monster.Pv, monster.PvMax)
		fmt.Println()

		fmt.Println("1. Attaquer")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Abandonner")
		fmt.Print("\nChoix : ")

		choice := utils.ReadInt()

		switch choice {
		case 1:
			damage := combat.PlayerAttack(c, &monster)
			fmt.Printf("\nVous infligez %d dégâts.\n", damage)

			if combat.IsDead(monster.Pv) {
				fmt.Println("Victoire ! Le gobelin est vaincu.")
				utils.Pause()
				return
			}

			damage = combat.MonsterAttack(c, &monster, turn)
			fmt.Printf("%s vous inflige %d dégâts.\n", monster.Name, damage)

			if combat.IsDead(c.Pv) {
				c.Pv = 0
				fmt.Println("\nVous êtes mort. Game Over.")
				utils.Pause()
				return
			}

			turn++
			utils.Pause()

		case 2:
			Inventaire(c, true)

		case 3:
			return

		default:
			fmt.Println("Choix invalide.")
			utils.Pause()
		}
	}
}
