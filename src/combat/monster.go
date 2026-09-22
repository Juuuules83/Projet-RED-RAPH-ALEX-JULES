package combat

// Monster représente un adversaire du joueur.
type Monster struct {
	Name   string
	PvMax  int
	Pv     int
	Attack int
}

// InitGoblin initialise le gobelin d'entraînement.
func InitGoblin() Monster {
	return Monster{
		Name:   "Gobelin d'entraînement",
		PvMax:  40,
		Pv:     40,
		Attack: 5,
	}
}
