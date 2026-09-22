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

func MenuPrincipal(character *player.Character) {
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
			AfficherPersonnage(character)
		case 2:
			Inventaire(character, false)
		case 3:
			Marchand(character)
		case 4:
			Combat(character)
		case 5:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
			utils.Pause()
		}
	}
}

func AfficherPersonnage(character *player.Character) {
	utils.ClearScreen()
	fmt.Println("===== PERSONNAGE =====")
	fmt.Println("Nom :", character.Name)
	fmt.Println("Classe :", character.Classe)
	fmt.Printf("PV : %d/%d\n", character.Pv, character.PvMax)
	fmt.Println("Niveau :", character.Niveau)
	fmt.Println("Argent :", character.Money, "€")
	utils.Pause()
}

func Inventaire(character *player.Character, depuisCombat bool) {
	for {
		utils.ClearScreen()
		fmt.Println("===== INVENTAIRE =====")
		fmt.Printf("Objets : %d/%d\n\n", character.TotalItems(), player.StockageMax)

		if len(character.Inventory) == 0 {
			fmt.Println("L'inventaire est vide.")
		} else {
			fmt.Println("Objets possédés :")
			for item, quantity := range character.Inventory {
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
			fmt.Println(character.UseLifePotion())
			utils.Pause()
		case 2:
			return
		default:
			fmt.Println("Choix invalide.")
			utils.Pause()
		}
	}
}

func Marchand(character *player.Character) {
	for {
		utils.ClearScreen()
		fmt.Println("===== MARCHAND =====")
		fmt.Println("Argent :", character.Money, "€")
		fmt.Printf("Inventaire : %d/%d\n\n", character.TotalItems(), player.StockageMax)

		fmt.Println("1. Potion de vie - 10 €")
		fmt.Println("2. Retour")
		fmt.Print("\nChoix : ")

		choice := utils.ReadInt()

		switch choice {
		case 1:
			fmt.Println(trader.BuyLifePotion(character))
			utils.Pause()
		case 2:
			return
		default:
			fmt.Println("Choix invalide.")
			utils.Pause()
		}
	}
}

func Combat(character *player.Character) {
	monster := combat.InitGoblin()
	turn := 1

	for {
		utils.ClearScreen()
		fmt.Println("===== COMBAT =====")
		fmt.Println()
		fmt.Printf("%s : %d/%d PV\n", character.Name, character.Pv, character.PvMax)
		fmt.Printf("%s : %d/%d PV\n", monster.Name, monster.Pv, monster.PvMax)
		fmt.Println()

		fmt.Println("1. Attaquer")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Abandonner")
		fmt.Print("\nChoix : ")

		choice := utils.ReadInt()

		switch choice {
		case 1:
			damage := combat.PlayerAttack(character, &monster)
			fmt.Printf("\nVous infligez %d dégâts.\n", damage)

			if combat.IsDead(monster.Pv) {
				fmt.Println("Victoire ! Le gobelin est vaincu.")
				utils.Pause()
				return
			}

			damage = combat.MonsterAttack(character, &monster, turn)
			fmt.Printf("%s vous inflige %d dégâts.\n", monster.Name, damage)

			if combat.IsDead(character.Pv) {
				character.Pv = 0
				fmt.Println("\nVous êtes mort. Game Over.")
				utils.Pause()
				return
			}

			turn++
			utils.Pause()

		case 2:
			Inventaire(character, true)

		case 3:
			return

		default:
			fmt.Println("Choix invalide.")
			utils.Pause()
		}
	}
}
