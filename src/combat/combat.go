package combat

import (
	"projet-red/player"
)

const BasicAttackDamage = 5

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
