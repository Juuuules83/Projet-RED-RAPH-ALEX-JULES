package combat

import (
	"fmt"
	"projet-red/player"
)

const BasicAttackDamage = 5

// PlayerAttack applique l'attaque de base du joueur au monstre.
func PlayerAttack(character *player.Character, monster *Monster) int {
	if monster == nil || monster.Pv <= 0 {
		return 0
	}

	monster.Pv -= BasicAttackDamage
	if monster.Pv < 0 {
		monster.Pv = 0
	}

	return BasicAttackDamage
}

// MonsterAttack applique l'attaque du gobelin au joueur.
// Tous les 3 tours, les dégâts sont doublés.
func MonsterAttack(character *player.Character, monster *Monster, turn int) int {
	if character == nil || monster == nil || monster.Pv <= 0 || character.Pv <= 0 {
		return 0
	}

	damage := monster.Attack
	if turn%3 == 0 {
		damage *= 2
	}

	character.Pv -= damage
	if character.Pv < 0 {
		character.Pv = 0
	}

	return damage
}

// CombatStatus construit le message du tour.
func CombatStatus(character *player.Character, monster *Monster, playerDamage int, monsterDamage int) string {
	return fmt.Sprintf(
		"Vous infligez %d dégâts. %s inflige %d dégâts.",
		playerDamage,
		monster.Name,
		monsterDamage,
	)
}
