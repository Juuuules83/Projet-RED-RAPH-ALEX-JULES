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
    utils.ClearScreen()

    fmt.Println("===== CHOISIR UNE ATTAQUE =====")
    fmt.Println()
fmt.Println("1. Attaque de base (5 dégâts)")

if c.Inventory[utils.FireBall] > 0 {
    fmt.Println("2. Fireball (20 dégâts)")
} else {
    fmt.Println("2. Fireball [VERROUILLÉ]")
}

fmt.Println("3. Retour")
    fmt.Print("\nChoix : ")

    attackChoice := utils.ReadInt()

    var damage int

    switch attackChoice {
    case 1:
        damage = PlayerAttack(c, &monster)
        fmt.Printf(
            "\nVous infligez %d dégâts avec votre attaque de base.\n",
            damage,
        )


case 2:
    if c.Inventory[utils.FireBall] == 0 {
        fmt.Println("ACHETE LA FIREBALL POUR ALLUMER LA GUEULE DE CE MISEREUX GOBLIN LA TU SAIS PAS LE FAIRE TROU D'UC !")
        utils.Pause()
        continue
    }

    damage = FireBall(c, &monster)

    fmt.Printf(
        "\nVous infligez %d dégâts avec Fireball !\n",
        damage,
    )

    case 3:
        continue

    default:
        fmt.Println("Choix invalide.")
        utils.Pause()
        continue
    }

    if IsDead(monster.Pv) {
        fmt.Println("Victoire ! Le gobelin est vaincu. WHOUHOUUUUUUUUUUUUUU !")
        utils.Pause()
        return
    }

    damage = MonsterAttack(c, &monster, turn)

    fmt.Printf(
        "%s vous inflige %d dégâts.\n",
        monster.Name,
        damage,
    )

    if IsDead(c.Pv) {
        c.Pv = 0
        fmt.Println("\nT'es NUL ! \n...Game Over...")
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
			fmt.Println("Choix invalide.")
			utils.Pause()
		}
	}
}