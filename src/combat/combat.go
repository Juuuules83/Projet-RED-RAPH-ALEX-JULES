package combat

import (
	"fmt"
	"projet-red/player"
	"projet-red/utils"
)

const BasicAttackDamage = 5
const FireBallDamage = 20

func PlayerAttack(c *player.Character, m *Monster) int {
	m.Pv -= BasicAttackDamage

	if m.Pv < 0 {
		m.Pv = 0
	}
	return BasicAttackDamage
}

func FireBall(c *player.Character, m *Monster) int {
	m.Pv -= FireBallDamage

	if m.Pv < 0 {
		m.Pv = 0
	}
	return FireBallDamage
}

func MonsterAttack(c *player.Character, m *Monster, t int) int {
	damage := m.Attack

	if t%3 == 0 {
		damage *= 2
	}

	c.Pv -= damage

	if c.Pv < 0 {
		c.Pv = 0
	}
	return damage
}

func Combat(c *player.Character) {
	monster := InitGoblin()
	turn := 1

	for {
		utils.ClearScreen()
		fmt.Println(utils.Bold + utils.Red + "===== COMBAT =====" + utils.Reset)
		fmt.Println()
		fmt.Printf(utils.Green+"%s : %d/%d PV\n"+utils.Reset, c.Name, c.Pv, c.PvMax)
		fmt.Printf(utils.Red+"%s : %d/%d PV\n"+utils.Reset, monster.Name, monster.Pv, monster.PvMax)
		fmt.Println()
		fmt.Println(utils.Bold + utils.Green + "[ 1 ] - Attaquer" + utils.Reset)
		fmt.Println(utils.Cyan + "[ 2 ] - Inventaire" + utils.Reset)
		fmt.Println(utils.Red + "[ 3 ] - Abandonner" + utils.Reset)
		fmt.Print(utils.Yellow + "\nChoix : " + utils.Reset)
		choice := utils.ReadInt()
		switch choice {

		case 1:
			utils.ClearScreen()
			fmt.Println(utils.Bold + utils.Magenta + "===== CHOISIR UNE ATTAQUE =====" + utils.Reset)
			fmt.Println()
			fmt.Println(utils.Green + "[ 1 ] - Coup de poing (5 dégâts)" + utils.Reset)
			if c.Inventory[utils.FireBall] > 0 {
				fmt.Println(utils.Magenta + utils.Bold + "[ 2 ] - Exploit de faille (20 dégâts)" + utils.Reset)
			} else {
				fmt.Println(utils.Red + utils.Bold + "[ 2 ] - Exploit de faille [VERROUILLÉ]" + utils.Reset)
			}

			fmt.Println(utils.Red + "[ 3 ] - Retour" + utils.Reset)
			fmt.Print(utils.Yellow + "\nChoix : " + utils.Reset)
			attackChoice := utils.ReadInt()
			var damage int
			switch attackChoice {
			case 1:
				damage = PlayerAttack(c, &monster)
				fmt.Printf(
					utils.Green+"\nVous infligez %d dégâts avec votre coup de poing.\n"+utils.Reset,
					damage,
				)

			case 2:
				if c.Inventory[utils.FireBall] == 0 {
					fmt.Println(utils.Yellow + "Achète l’Exploit de faille à Ymatch pour débloquer cette attaque !" + utils.Reset)
					utils.Pause()
					continue
				}

				damage = FireBall(c, &monster)

				fmt.Printf(
					utils.Green+"\nVous infligez %d dégâts avec l’Exploit de faille !\n"+utils.Reset,
					damage,
				)

			case 3:
				continue

			default:
				fmt.Println(utils.Red + "Choix invalide." + utils.Reset)
				utils.Pause()
				continue
			}

			if IsDead(monster.Pv) {
				fmt.Println(utils.Green + "Victoire ! Gopher est VAINCU !" + utils.Reset)
				utils.Pause()
				return
			}

			damage = MonsterAttack(c, &monster, turn)

			fmt.Printf(
				utils.Red+"%s vous inflige %d dégâts.\n"+utils.Reset,
				monster.Name,
				damage,
			)

			if IsDead(c.Pv) {
				c.Pv = 0
				fmt.Println(utils.Red + " GAME OVER" + utils.Reset)
				utils.Pause()
				return
			}

			turn++
			utils.Pause()
		case 2:
			player.Inventaire(c, true)
		case 3:
			return
		default:
			fmt.Println(utils.Red + "Choix invalide." + utils.Reset)
			utils.Pause()
		}
	}
}
