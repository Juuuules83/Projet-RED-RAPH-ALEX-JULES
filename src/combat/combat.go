package combat

import (
	"projet-red/player"
)

const BasicAttackDamage = 5

func PlayerAttack(c *player.Character, m *Monster) int {
	if m == nil || m.Pv <= 0 {
		return 0
	}

	m.Pv -= BasicAttackDamage

	if m.Pv < 0 {
		m.Pv = 0
	}

	return BasicAttackDamage
}

func MonsterAttack(c *player.Character, m *Monster, t int) int {
	if c == nil || m == nil || m.Pv <= 0 || c.Pv <= 0 {
		return 0
	}

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
