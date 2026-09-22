package combat

// IsDead indique si le joueur n'a plus de PV.
func IsDead(characterPV int) bool {
	return characterPV <= 0
}
